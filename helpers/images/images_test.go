package images_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	dedalus "github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/helpers/images"
	"github.com/dedalus-labs/dedalus-go/option"
)

const imageID = "01994a10-1234-7000-8000-000000000001"

var archiveBytes = []byte("fixture docker-save archive")

type transport func(*http.Request) (*http.Response, error)

func (f transport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func response(request *http.Request, status int, body interface{}) *http.Response {
	data, _ := json.Marshal(body)
	return &http.Response{StatusCode: status, Header: http.Header{"Content-Type": {"application/json"}, "Retry-After-Ms": {"1"}}, Body: io.NopCloser(bytes.NewReader(data)), Request: request}
}

type fixture struct {
	client   *dedalus.Client
	storage  http.RoundTripper
	grant    map[string]interface{}
	requests []string
	keys     []string
	uploads  int
	states   []string
	status   int
	retry    bool
	wrongID  bool
}

func setup(t *testing.T, states ...string) *fixture {
	t.Helper()
	digest := sha256.Sum256(archiveBytes)
	f := &fixture{states: states, status: 200}
	f.grant = map[string]interface{}{
		"source": map[string]interface{}{"kind": "docker_archive", "upload_id": imageID, "sha256": hex.EncodeToString(digest[:])},
		"upload": map[string]interface{}{"url": "https://storage.example/image?signature=fixture", "method": "PUT", "expires_at": "2099-01-01T00:00:00Z",
			"headers": map[string][]string{"content-length": {strconv.Itoa(len(archiveBytes))}, "if-none-match": {"*"}, "x-amz-checksum-sha256": {base64.StdEncoding.EncodeToString(digest[:])}}},
	}
	client := dedalus.NewClient(option.WithAPIKey("fixture-secret"), option.WithBaseURL("https://api.example"), option.WithMaxRetries(1), option.WithHTTPClient(&http.Client{Transport: transport(func(r *http.Request) (*http.Response, error) {
		f.requests = append(f.requests, r.Method+" "+r.URL.Path)
		if r.Header.Get("Authorization") != "Bearer fixture-secret" {
			t.Error("API credential missing")
		}
		if r.Method == http.MethodPost {
			f.keys = append(f.keys, r.Header.Get("Idempotency-Key"))
		}
		if strings.HasSuffix(r.URL.Path, "/uploads") {
			if f.retry && len(f.keys) == 1 {
				return response(r, 503, map[string]string{}), nil
			}
			return response(r, 200, f.grant), nil
		}
		if r.URL.Path == "/v1/machines" {
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["image_id"] != imageID {
				t.Fatalf("machine image ID not preserved: %v, %v", body, err)
			}
			return response(r, 200, map[string]interface{}{"machine_id": "fixture", "lifecycle_status": map[string]interface{}{"autoresize": map[string]interface{}{
				"cpu": map[string]string{"blocked_reason": "MACHINE_QUOTA_EXCEEDED", "blocked_at": "2026-01-01T00:00:00Z"}, "memory": nil,
			}}}), nil
		}
		state, id := "submitted", imageID
		if r.Method == http.MethodGet {
			state = "preparing"
			if len(f.states) > 0 {
				state, f.states = f.states[0], f.states[1:]
			}
			if f.wrongID {
				id = "wrong"
			}
		}
		return response(r, 200, map[string]interface{}{"image_id": id, "state": state, "error": map[string]string{"error_code": "PREPARE_FAILED", "message": "archive rejected"}}), nil
	})}))
	f.client = &client
	f.storage = transport(func(r *http.Request) (*http.Response, error) {
		f.uploads++
		if r.Header.Get("Authorization") != "" || r.Header.Get("Cookie") != "" {
			t.Error("API credentials reached storage")
		}
		body, err := io.ReadAll(r.Body)
		if err != nil || !bytes.Equal(body, archiveBytes) || r.ContentLength != int64(len(archiveBytes)) {
			t.Error("archive bytes or size changed")
		}
		result := response(r, f.status, nil)
		result.Header.Set("Location", "https://redirect.example/")
		return result, nil
	})
	return f
}

func archive(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "image.tar")
	if err := os.WriteFile(path, archiveBytes, 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestImportDefaultWaitAndNoWait(t *testing.T) {
	for _, noWait := range []bool{false, true} {
		t.Run(map[bool]string{true: "no-wait", false: "wait"}[noWait], func(t *testing.T) {
			f := setup(t, "preparing", "ready")
			f.retry = true
			var progress []string
			result, err := images.Import(context.Background(), f.client, archive(t), images.ImportOptions{NoWait: noWait, UploadTransport: f.storage, PollOptions: images.PollOptions{PollInterval: time.Millisecond, OnProgress: func(image images.Image) { progress = append(progress, image.State) }}})
			if err != nil {
				t.Fatal(err)
			}
			want := []string{"submitted", "preparing", "ready"}
			if noWait {
				want = []string{"submitted"}
			}
			if !reflect.DeepEqual(progress, want) || result.State != want[len(want)-1] {
				t.Fatalf("progress = %v; result = %v", progress, result.State)
			}
			if f.uploads != 1 || len(f.keys) != 3 || f.keys[0] == "" || f.keys[0] != f.keys[1] || f.keys[1] != f.keys[2] {
				t.Fatal("upload count or retry identity changed")
			}
			if noWait && len(f.requests) != 3 {
				t.Fatal("no-wait requested status")
			}
		})
	}
}

func TestUnsafeGrantsNeverUploadOrRegister(t *testing.T) {
	for name, change := range map[string]func(map[string]interface{}){
		"http":    func(u map[string]interface{}) { u["url"] = "http://storage.example/x" },
		"bad URL": func(u map[string]interface{}) { u["url"] = "https://invalid%fixture" },
		"method":  func(u map[string]interface{}) { u["method"] = "POST" },
		"expired": func(u map[string]interface{}) { u["expires_at"] = "2000-01-01T00:00:00Z" },
		"auth": func(u map[string]interface{}) {
			u["headers"].(map[string][]string)["authorization"] = []string{"fixture"}
		},
		"size": func(u map[string]interface{}) { u["headers"].(map[string][]string)["content-length"] = []string{"1"} },
		"checksum": func(u map[string]interface{}) {
			u["headers"].(map[string][]string)["x-amz-checksum-sha256"] = []string{"wrong"}
		},
	} {
		t.Run(name, func(t *testing.T) {
			f := setup(t)
			change(f.grant["upload"].(map[string]interface{}))
			_, err := images.Import(context.Background(), f.client, archive(t), images.ImportOptions{UploadTransport: f.storage})
			if err == nil || f.uploads != 0 || len(f.requests) != 1 {
				t.Fatalf("unsafe grant accepted: %v", err)
			}
		})
	}
}

func TestTerminalFailuresAndUnknownStates(t *testing.T) {
	for _, state := range []string{"failed", "rejected", "revoked"} {
		f := setup(t, state)
		_, err := images.Poll(context.Background(), f.client, imageID, images.PollOptions{})
		var failure *images.PreparationError
		if !errors.As(err, &failure) || failure.Image.ImageID != imageID || !strings.Contains(err.Error(), "PREPARE_FAILED") {
			t.Fatalf("lost image failure: %v", err)
		}
	}
	f := setup(t, "future-state", "ready")
	if _, err := images.Poll(context.Background(), f.client, imageID, images.PollOptions{PollInterval: time.Millisecond}); err != nil {
		t.Fatal(err)
	}
	f.wrongID = true
	if _, err := images.Poll(context.Background(), f.client, imageID, images.PollOptions{}); err == nil {
		t.Fatal("accepted mismatched image ID")
	}
}

func TestDeadlineAndCancellationPreserveImage(t *testing.T) {
	f := setup(t)
	_, err := images.Import(context.Background(), f.client, archive(t), images.ImportOptions{UploadTransport: f.storage, PollOptions: images.PollOptions{Timeout: 100 * time.Millisecond, PollInterval: time.Millisecond}})
	var failure *images.PreparationError
	if !errors.Is(err, context.DeadlineExceeded) || !errors.As(err, &failure) || failure.Image.ImageID != imageID {
		t.Fatalf("deadline lost image ID: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	_, err = images.Poll(ctx, f.client, imageID, images.PollOptions{OnProgress: func(images.Image) { cancel() }})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation: %v", err)
	}
	for _, request := range f.requests {
		if strings.HasPrefix(request, "DELETE") {
			t.Fatal("cancel deleted image")
		}
	}
}

func TestInFlightCancellationAndUploadErrors(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	client := dedalus.NewClient(option.WithAPIKey("fixture"), option.WithHTTPClient(&http.Client{Transport: transport(func(r *http.Request) (*http.Response, error) {
		cancel()
		<-r.Context().Done()
		return nil, r.Context().Err()
	})}))
	if _, err := images.Poll(ctx, &client, imageID, images.PollOptions{}); !errors.Is(err, context.Canceled) {
		t.Fatalf("in-flight cancellation: %v", err)
	}
	for _, status := range []int{302, 500, 412} {
		f := setup(t)
		f.status = status
		result, err := images.Import(context.Background(), f.client, archive(t), images.ImportOptions{NoWait: true, UploadTransport: f.storage})
		if status == 412 {
			if err != nil || result.State != "submitted" {
				t.Fatalf("repeated PUT: %v", err)
			}
		} else if err == nil || len(f.requests) != 1 || f.uploads != 1 {
			t.Fatalf("upload status %d continued: %v", status, err)
		}
	}
	f := setup(t)
	_, err := images.Import(context.Background(), f.client, archive(t), images.ImportOptions{UploadTransport: transport(func(*http.Request) (*http.Response, error) { return nil, errors.New("transport failed") })})
	if err == nil || strings.Contains(err.Error(), "signature=") {
		t.Fatalf("upload error exposed URL: %v", err)
	}
}

func TestInvalidInputMakesNoRequests(t *testing.T) {
	f := setup(t)
	path := archive(t)
	if err := os.Truncate(path, 0); err != nil {
		t.Fatal(err)
	}
	if _, err := images.Import(context.Background(), f.client, path, images.ImportOptions{}); err == nil {
		t.Fatal("accepted empty archive")
	}
	if _, err := images.Poll(context.Background(), f.client, imageID, images.PollOptions{Timeout: -1}); err == nil {
		t.Fatal("accepted negative timeout")
	}
	if len(f.requests) != 0 {
		t.Fatal("invalid input made API request")
	}
}

func TestReadyImageToGeneratedMachine(t *testing.T) {
	f := setup(t, "ready")
	ready, err := images.Import(context.Background(), f.client, archive(t), images.ImportOptions{UploadTransport: f.storage, PollOptions: images.PollOptions{PollInterval: time.Millisecond}})
	if err != nil {
		t.Fatal(err)
	}
	machine, err := f.client.Machines.New(context.Background(), dedalus.MachineNewParams{CreateParams: dedalus.CreateParams{ImageID: dedalus.String(ready.ImageID)}})
	if err != nil {
		t.Fatal(err)
	}
	if machine.MachineID != "fixture" {
		t.Fatalf("lost refusal: %+v", machine)
	}
	if len(f.requests) != 4 || f.requests[3] != "POST /v1/machines" {
		t.Fatalf("unexpected workflow: %v", f.requests)
	}
}

func TestGeneratedPatchResizeAndFlatErrors(t *testing.T) {
	requests := []string{}
	envelope := map[string]interface{}{"error_code": "IMAGE_SOURCE_INVALID", "message": "archive rejected", "retryable": false, "details": map[string]string{"stage": "validation"}, "retry_after_ms": 500}
	client := dedalus.NewClient(option.WithAPIKey("fixture"), option.WithBaseURL("https://api.example"), option.WithMaxRetries(0), option.WithHTTPClient(&http.Client{Transport: transport(func(r *http.Request) (*http.Response, error) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		if strings.Contains(r.URL.Path, "/images/") {
			return response(r, 422, envelope), nil
		}
		return response(r, 200, map[string]interface{}{"status": map[string]interface{}{"autoresize": map[string]interface{}{"cpu": map[string]string{"blocked_reason": "MACHINE_QUOTA_EXCEEDED"}, "memory": nil}}}), nil
	})}))
	ctx := context.Background()
	if _, err := client.Machines.Autoresizing.Update(ctx, dedalus.MachineAutoresizingUpdateParams{MachineID: imageID}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Organization.Autoresizing.Update(ctx, dedalus.OrganizationAutoresizingUpdateParams{}); err != nil {
		t.Fatal(err)
	}
	resized, err := client.Machines.Resize(ctx, dedalus.MachineResizeParams{MachineID: imageID})
	if err != nil {
		t.Fatal(err)
	}
	if resized.Status.Autoresize.CPU.BlockedReason != "MACHINE_QUOTA_EXCEEDED" {
		t.Fatalf("lost refusal: %+v", resized.Status)
	}
	_, err = client.Images.Get(ctx, dedalus.ImageGetParams{ImageID: imageID})
	var apiErr *dedalus.Error
	if !errors.As(err, &apiErr) || !strings.Contains(err.Error(), "archive rejected") {
		t.Fatalf("missing flat error: %v", err)
	}
	var decoded map[string]interface{}
	if err := json.Unmarshal([]byte(apiErr.RawJSON()), &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["error_code"] != envelope["error_code"] || decoded["retryable"] != false || decoded["retry_after_ms"] != float64(500) {
		t.Fatalf("lost fields: %v", decoded)
	}
	if requests[0] != "PATCH /v1/machines/"+imageID+"/autoresizing" || requests[1] != "PATCH /v1/organization/autoresizing" || requests[2] != "POST /v1/machines/"+imageID+"/resize" {
		t.Fatalf("wrong methods: %v", requests)
	}
}

func TestPreparationFailureRetainsRetryMetadata(t *testing.T) {
	var failure images.Failure
	if err := json.Unmarshal([]byte(`{"error_code":"IMAGE_SOURCE_UNAVAILABLE","message":"registry offline","retryable":true,"retry_after_ms":500,"details":{"stage":"download"}}`), &failure); err != nil {
		t.Fatal(err)
	}
	if !failure.Retryable || failure.RetryAfterMs == nil || *failure.RetryAfterMs != 500 || failure.Details["stage"] != "download" {
		t.Fatalf("lost failure metadata: %+v", failure)
	}
}

func TestUnsupportedPlatformNeverReserves(t *testing.T) {
	f := setup(t)
	_, err := images.Import(context.Background(), f.client, archive(t), images.ImportOptions{Platform: "linux/arm64", UploadTransport: f.storage})
	if err == nil || !strings.Contains(err.Error(), "linux/amd64") || len(f.requests) != 0 || f.uploads != 0 {
		t.Fatalf("unsupported platform caused IO: %v", err)
	}
}
