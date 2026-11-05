// Package images provides upload and readiness workflows for custom images.
package images

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	dedalus "github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/option"
)

// Image is the registration and readiness information returned by the API.
type Image struct {
	ImageID string          `json:"image_id"`
	State   string          `json:"state"`
	Name    string          `json:"name"`
	Reason  string          `json:"reason"`
	Context json.RawMessage `json:"context"`
	Error   *Failure        `json:"error"`
}

// Failure describes why preparation failed.
type Failure struct {
	Code    string `json:"error_code"`
	Message string `json:"message"`
}

// PreparationError preserves the image ID on failure, timeout, or cancellation.
// Unwrap supports errors.Is with context.Canceled and context.DeadlineExceeded.
type PreparationError struct {
	Image *Image
	Err   error
}

func (e *PreparationError) Error() string { return fmt.Sprintf("image %s: %v", e.Image.ImageID, e.Err) }
func (e *PreparationError) Unwrap() error { return e.Err }

// PollOptions controls readiness polling. Zero durations select the defaults.
type PollOptions struct {
	Timeout      time.Duration // Default: one hour; does not cancel server preparation.
	PollInterval time.Duration // Default: two seconds.
	OnProgress   func(Image)   // Called synchronously for each observed image.
}

// ImportOptions controls a streamed docker-save archive import.
type ImportOptions struct {
	PollOptions
	NoWait         bool // Return after upload and registration; default is to wait.
	Name           string
	Platform       string // Default: linux/amd64; also accepts linux/arm64.
	IdempotencyKey string // Reuse for retries of the same archive and options.
	// Optional storage transport, separate from the authenticated API client.
	UploadTransport http.RoundTripper
}

func (o PollOptions) timing() (time.Duration, time.Duration, error) {
	timeout, interval := o.Timeout, o.PollInterval
	if timeout == 0 {
		timeout = time.Hour
	}
	if interval == 0 {
		interval = 2 * time.Second
	}
	if timeout < 0 || interval < 0 || timeout > 24*time.Hour || interval > 24*time.Hour {
		return 0, 0, errors.New("timeout and poll interval must be positive and at most 24 hours")
	}
	return timeout, interval, nil
}

func validateImage(image *Image, expected string) error {
	if image.ImageID == "" || image.State == "" || (expected != "" && image.ImageID != expected) {
		return errors.New("invalid or mismatched image API response")
	}
	return nil
}

func wait(ctx context.Context, client *dedalus.Client, image *Image, interval time.Duration, progress func(Image)) (*Image, error) {
	for {
		if err := ctx.Err(); err != nil {
			return nil, interrupted(image, err)
		}
		if progress != nil {
			progress(*image)
		}
		switch image.State {
		case "ready":
			return image, nil
		case "failed", "rejected", "revoked":
			message := image.Reason
			if message == "" {
				message = image.State
			}
			if image.Error != nil {
				if image.Error.Message != "" {
					message = image.Error.Message
				}
				if image.Error.Code != "" {
					message = image.Error.Code + ": " + message
				}
			}
			return nil, &PreparationError{Image: image, Err: errors.New(message)}
		}
		timer := time.NewTimer(interval)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, interrupted(image, ctx.Err())
		case <-timer.C:
		}
		var next Image
		if err := client.Get(ctx, "/v1/images/"+url.PathEscape(image.ImageID), nil, &next); err != nil {
			if ctx.Err() != nil {
				return nil, interrupted(image, ctx.Err())
			}
			return nil, &PreparationError{Image: image, Err: err}
		}
		if err := validateImage(&next, image.ImageID); err != nil {
			return nil, err
		}
		image = &next
	}
}

func interrupted(image *Image, err error) error {
	return &PreparationError{Image: image, Err: fmt.Errorf("waiting stopped; preparation continues, use this image ID to resume polling: %w", err)}
}

// Poll waits for an existing image to become ready. Unknown states keep polling.
func Poll(ctx context.Context, client *dedalus.Client, imageID string, options PollOptions) (*Image, error) {
	timeout, interval, err := options.timing()
	if err != nil {
		return nil, err
	}
	if imageID == "" || strings.ContainsAny(imageID, " \t\r\n") {
		return nil, errors.New("image ID must be nonempty without whitespace")
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return nil, interrupted(&Image{ImageID: imageID, State: "unknown"}, err)
	}
	var image Image
	if err := client.Get(ctx, "/v1/images/"+url.PathEscape(imageID), nil, &image); err != nil {
		if ctx.Err() != nil {
			return nil, interrupted(&Image{ImageID: imageID, State: "unknown"}, ctx.Err())
		}
		return nil, err
	}
	if err := validateImage(&image, imageID); err != nil {
		return nil, err
	}
	return wait(ctx, client, &image, interval, options.OnProgress)
}

type source struct {
	Kind     string `json:"kind"`
	UploadID string `json:"upload_id"`
	SHA256   string `json:"sha256"`
}

type reservation struct {
	Source source `json:"source"`
	Upload struct {
		URL       string              `json:"url"`
		Method    string              `json:"method"`
		Headers   map[string][]string `json:"headers"`
		ExpiresAt time.Time           `json:"expires_at"`
	} `json:"upload"`
}

func (r reservation) capability(digest string, size int64) (*url.URL, http.Header, error) {
	invalid := errors.New("invalid image upload capability")
	endpoint, err := url.Parse(r.Upload.URL)
	if err != nil || endpoint.Scheme != "https" || endpoint.Hostname() == "" || endpoint.User != nil || endpoint.Fragment != "" || endpoint.Opaque != "" {
		return nil, nil, invalid
	}
	if r.Source.Kind != "docker_archive" || r.Source.UploadID == "" || r.Source.SHA256 != digest || r.Upload.Method != http.MethodPut || !r.Upload.ExpiresAt.After(time.Now()) {
		return nil, nil, invalid
	}
	headers := make(http.Header)
	for key, values := range r.Upload.Headers {
		switch strings.ToLower(key) {
		case "authorization", "proxy-authorization", "cookie":
			return nil, nil, invalid
		}
		for _, value := range values {
			headers.Add(key, value)
		}
	}
	checksum := headers.Get("X-Amz-Checksum-Sha256")
	if checksum == "" {
		checksum = endpoint.Query().Get("X-Amz-Checksum-Sha256")
	}
	bytes, err := hex.DecodeString(digest)
	if err != nil || checksum != base64.StdEncoding.EncodeToString(bytes) || len(headers.Values("X-Amz-Checksum-Sha256")) > 1 ||
		len(headers.Values("Content-Length")) != 1 || headers.Get("Content-Length") != strconv.FormatInt(size, 10) ||
		len(headers.Values("If-None-Match")) != 1 || headers.Get("If-None-Match") != "*" {
		return nil, nil, invalid
	}
	return endpoint, headers, nil
}

type contextReader struct {
	context.Context
	io.Reader
}

func (r contextReader) Read(p []byte) (int, error) {
	if err := r.Err(); err != nil {
		return 0, err
	}
	return r.Reader.Read(p)
}

// Import streams a docker-save archive, registers it, and waits by default.
// Timeout covers hashing, upload, registration, and polling. Cancelling ctx only
// stops the client; it never deletes an image or cancels accepted preparation.
func Import(ctx context.Context, client *dedalus.Client, path string, options ImportOptions) (*Image, error) {
	timeout, interval, err := options.PollOptions.timing()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	platform := options.Platform
	if platform == "" {
		platform = "linux/amd64"
	}
	if platform != "linux/amd64" && platform != "linux/arm64" {
		return nil, errors.New("platform must be linux/amd64 or linux/arm64")
	}
	key := options.IdempotencyKey
	if key == "" {
		var bytes [16]byte
		if _, err := rand.Read(bytes[:]); err != nil {
			return nil, err
		}
		key = hex.EncodeToString(bytes[:])
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	before, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !before.Mode().IsRegular() || before.Size() <= 0 || before.Size() > 5<<30 {
		return nil, errors.New("image archive must be a regular file between 1 byte and 5 GiB")
	}
	hash := sha256.New()
	count, err := io.Copy(hash, contextReader{ctx, file})
	if err != nil {
		return nil, err
	}
	after, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if count != before.Size() || after.Size() != before.Size() || !after.ModTime().Equal(before.ModTime()) {
		return nil, errors.New("image archive changed while hashing")
	}
	digest := hex.EncodeToString(hash.Sum(nil))
	var grant reservation
	requestOptions := []option.RequestOption{option.WithHeader("Idempotency-Key", key)}
	if err := client.Post(ctx, "/v1/images/uploads", map[string]interface{}{"sha256": digest, "size_bytes": count, "platform": platform}, &grant, requestOptions...); err != nil {
		return nil, err
	}
	endpoint, headers, err := grant.capability(digest, count)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, endpoint.String(), io.NewSectionReader(file, 0, count))
	if err != nil {
		return nil, errors.New("invalid image upload request")
	}
	request.Header, request.ContentLength = headers, count
	storage := &http.Client{Transport: options.UploadTransport, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	response, err := storage.Do(request)
	if err != nil {
		var urlError *url.Error
		if errors.As(err, &urlError) {
			err = urlError.Err
		} // Do not expose a signed URL in errors.
		return nil, fmt.Errorf("upload image archive: %w", err)
	}
	response.Body.Close()
	if (response.StatusCode < 200 || response.StatusCode >= 300) && response.StatusCode != http.StatusPreconditionFailed {
		return nil, fmt.Errorf("image archive upload failed: HTTP %d", response.StatusCode)
	}
	body := map[string]interface{}{"source": grant.Source, "platform": platform}
	if options.Name != "" {
		body["name"] = options.Name
	}
	var image Image
	if err := client.Post(ctx, "/v1/images", body, &image, requestOptions...); err != nil {
		return nil, err
	}
	if err := validateImage(&image, ""); err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, interrupted(&image, err)
	}
	if options.NoWait {
		// Check terminal failures even when readiness polling was not requested.
		if image.State != "failed" && image.State != "rejected" && image.State != "revoked" {
			if options.OnProgress != nil {
				options.OnProgress(image)
			}
			return &image, nil
		}
	}
	return wait(ctx, client, &image, interval, options.OnProgress)
}
