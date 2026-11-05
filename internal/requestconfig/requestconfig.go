// File generated from our OpenAPI spec by Scalar. See README.md for details.

package requestconfig

import (
	"bytes"
	"context"
	cryptorand "crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"mime"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dedalus-labs/dedalus-go/internal"
	"github.com/dedalus-labs/dedalus-go/internal/apierror"
	"github.com/dedalus-labs/dedalus-go/internal/apiform"
	"github.com/dedalus-labs/dedalus-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-go/internal/apiquery"
	"github.com/dedalus-labs/dedalus-go/internal/param"
	"github.com/dedalus-labs/dedalus-go/internal/redact"
)

func getDefaultHeaders() map[string]string {
	return map[string]string{
		"User-Agent": fmt.Sprintf("Dedalus/Go %s", internal.PackageVersion),
	}
}

func getNormalizedOS() string {
	switch runtime.GOOS {
	case "ios":
		return "iOS"
	case "android":
		return "Android"
	case "darwin":
		return "MacOS"
	case "window":
		return "Windows"
	case "freebsd":
		return "FreeBSD"
	case "openbsd":
		return "OpenBSD"
	case "linux":
		return "Linux"
	default:
		return fmt.Sprintf("Other:%s", runtime.GOOS)
	}
}

func getNormalizedArchitecture() string {
	switch runtime.GOARCH {
	case "386":
		return "x32"
	case "amd64":
		return "x64"
	case "arm":
		return "arm"
	case "arm64":
		return "arm64"
	default:
		return fmt.Sprintf("other:%s", runtime.GOARCH)
	}
}

func getPlatformProperties() map[string]string {
	return map[string]string{
		"X-Scalar-Lang":            "go",
		"X-Scalar-Package-Version": internal.PackageVersion,
		"X-Scalar-OS":              getNormalizedOS(),
		"X-Scalar-Arch":            getNormalizedArchitecture(),
		"X-Scalar-Runtime":         "go",
		"X-Scalar-Runtime-Version": runtime.Version(),
	}
}

type RequestOption interface {
	Apply(*RequestConfig) error
}

type RequestOptionFunc func(*RequestConfig) error
type PreRequestOptionFunc func(*RequestConfig) error

func (s RequestOptionFunc) Apply(r *RequestConfig) error    { return s(r) }
func (s PreRequestOptionFunc) Apply(r *RequestConfig) error { return s(r) }

// firstDotSegment reports the first `.` or `..` segment in a request path, if any.
//
// These are RFC 3986 dot-segments. BaseURL.Parse below resolves the path as a reference
// against the base, which applies remove_dot_segments and walks up — so a path param
// carrying `..` retargets the request at a different endpoint no matter how the rest of
// it is escaped. url.PathEscape does not encode `.`, and PathEscapeReserved deliberately
// keeps `/`, so neither escaper stops this on its own.
//
// The query and fragment are excluded: only the path participates in dot-segment removal,
// and a `..` inside a query value is an ordinary character.
func firstDotSegment(path string) (string, bool) {
	if index := strings.IndexAny(path, "?#"); index >= 0 {
		path = path[:index]
	}
	for _, segment := range strings.Split(path, "/") {
		if segment == "." || segment == ".." {
			return segment, true
		}
	}
	return "", false
}

func NewRequestConfig(ctx context.Context, method string, u string, body interface{}, dst interface{}, opts ...RequestOption) (*RequestConfig, error) {
	// Reject before any request state is built: a dot-segment cannot be escaped into safety,
	// so the only correct handling is to refuse the call rather than send it somewhere else.
	if segment, found := firstDotSegment(u); found {
		return nil, fmt.Errorf("path %q contains a %q segment, which would resolve to a different path", u, segment)
	}

	var reader io.Reader

	contentType := "application/json"
	hasSerializationFunc := false

	if body, ok := body.(json.Marshaler); ok {
		content, err := body.MarshalJSON()
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(content)
		hasSerializationFunc = true
	}
	if body, ok := body.(apiform.Marshaler); ok {
		var (
			content []byte
			err     error
		)
		content, contentType, err = body.MarshalMultipart()
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(content)
		hasSerializationFunc = true
	}
	if body, ok := body.(apiquery.Queryer); ok {
		hasSerializationFunc = true
		params := body.URLQuery().Encode()
		if params != "" {
			parsed, err := url.Parse(u)
			if err != nil {
				return nil, err
			}
			if parsed.RawQuery != "" {
				parsed.RawQuery = parsed.RawQuery + "&" + params
				u = parsed.String()
			} else {
				u = u + "?" + params
			}
		}
	}
	if body, ok := body.([]byte); ok {
		reader = bytes.NewBuffer(body)
		hasSerializationFunc = true
	}
	if body, ok := body.(io.Reader); ok {
		reader = body
		hasSerializationFunc = true
	}

	// Fallback to json serialization if none of the serialization functions that we expect
	// to see is present.
	if body != nil && !hasSerializationFunc {
		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(true)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
		reader = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, u, nil)
	if err != nil {
		return nil, err
	}
	if reader != nil {
		req.Header.Set("Content-Type", contentType)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Scalar-Retry-Count", "0")
	req.Header.Set("X-Scalar-Timeout", "0")
	for k, v := range getDefaultHeaders() {
		req.Header.Add(k, v)
	}
	req.Header.Set("X-SDK-Version", "1.0.0")
	for k, v := range getPlatformProperties() {
		req.Header.Add(k, v)
	}
	if method != http.MethodGet {
		// Note this can be overridden with `option.WithHeader("Idempotency-Key", myIdempotencyKey)`
		req.Header.Set("Idempotency-Key", "scalar-go-"+newUUIDv4())
	}

	cfg := RequestConfig{
		MaxRetries: 2,
		Context:    ctx,
		Request:    req,
		HTTPClient: http.DefaultClient,
		Body:       reader,
	}
	cfg.ResponseBodyInto = dst
	err = cfg.Apply(opts...)
	if err != nil {
		return nil, err
	}

	// This must run after `cfg.Apply(...)` above in case the request timeout gets modified. We also only
	// apply our own logic for it if it's still "0" from above. If it's not, then it was deleted or modified
	// by the user and we should respect that.
	if req.Header.Get("X-Scalar-Timeout") == "0" {
		if cfg.RequestTimeout == time.Duration(0) {
			req.Header.Del("X-Scalar-Timeout")
		} else {
			req.Header.Set("X-Scalar-Timeout", strconv.Itoa(int(cfg.RequestTimeout.Seconds())))
		}
	}

	return &cfg, nil
}

func UseDefaultParam[T any](dst *param.Field[T], src *T) {
	if !dst.Present && src != nil {
		dst.Value = *src
		dst.Present = true
	}
}

// isUnionDestination reports whether dst points, through any number of pointers, at a union
// registered with apijson, which only apijson can decode into. Any other interface keeps
// encoding/json, which fills a non-nil one and returns an error, not a panic, for the rest.
func isUnionDestination(dst interface{}) bool {
	t := reflect.TypeOf(dst)
	for t != nil && t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Interface && apijson.IsUnion(t)
}

// This interface is primarily used to describe an [*http.Client], but also
// supports custom HTTP implementations.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// RequestConfig represents all the state related to one request.
//
// Editing the variables inside RequestConfig directly is unstable api. Prefer
// composing the RequestOption instead if possible.
type RequestConfig struct {
	MaxRetries     int
	RequestTimeout time.Duration
	Context        context.Context
	Request        *http.Request
	BaseURL        *url.URL
	// DefaultBaseURL will be used if BaseURL is not explicitly overridden using
	// WithBaseURL.
	DefaultBaseURL *url.URL
	CustomHTTPDoer HTTPDoer
	HTTPClient     *http.Client
	Middlewares    []middleware
	APIKey         string
	XAPIKey        string
	AsBaseURL      string
	DedalusOrgID   string
	BrowserToken   string
	// If ResponseBodyInto not nil, then we will attempt to deserialize into
	// ResponseBodyInto. If Destination is a []byte, then it will return the body as
	// is.
	ResponseBodyInto interface{}
	// ResponseInto copies the \*http.Response of the corresponding request into the
	// given address
	ResponseInto **http.Response
	Body         io.Reader
}

// middleware is exactly the same type as the Middleware type found in the [option] package,
// but it is redeclared here for circular dependency issues.
type middleware = func(*http.Request, middlewareNext) (*http.Response, error)

// middlewareNext is exactly the same type as the MiddlewareNext type found in the [option] package,
// but it is redeclared here for circular dependency issues.
type middlewareNext = func(*http.Request) (*http.Response, error)

func applyMiddleware(middleware middleware, next middlewareNext) middlewareNext {
	return func(req *http.Request) (res *http.Response, err error) {
		return middleware(req, next)
	}
}

func shouldRetry(req *http.Request, res *http.Response) bool {
	// If there is no way to recover the Body, then we shouldn't retry.
	if req.Body != nil && req.GetBody == nil {
		return false
	}

	// If there is no response, that indicates that there is a connection error
	// so we retry the request.
	if res == nil {
		return true
	}

	// If the header explicitly wants a retry behavior, respect that over the
	// http status code.
	if res.Header.Get("x-should-retry") == "true" {
		return true
	}
	if res.Header.Get("x-should-retry") == "false" {
		return false
	}

	return res.StatusCode == http.StatusRequestTimeout ||
		res.StatusCode == http.StatusConflict ||
		res.StatusCode == http.StatusTooManyRequests ||
		res.StatusCode >= http.StatusInternalServerError
}

func parseRetryAfterHeader(resp *http.Response) (time.Duration, bool) {
	if resp == nil {
		return 0, false
	}

	type retryData struct {
		header string
		units  time.Duration

		// custom is used when the regular algorithm failed and is optional.
		// the returned duration is used verbatim (units is not applied).
		custom func(string) (time.Duration, bool)
	}

	nop := func(string) (time.Duration, bool) { return 0, false }

	// the headers are listed in order of preference
	retries := []retryData{
		{
			header: "Retry-After-Ms",
			units:  time.Millisecond,
			custom: nop,
		},
		{
			header: "Retry-After",
			units:  time.Second,

			// retry-after values are expressed in either number of
			// seconds or an HTTP-date indicating when to try again
			custom: func(ra string) (time.Duration, bool) {
				t, err := time.Parse(time.RFC1123, ra)
				if err != nil {
					return 0, false
				}
				return time.Until(t), true
			},
		},
	}

	for _, retry := range retries {
		v := resp.Header.Get(retry.header)
		if v == "" {
			continue
		}
		if retryAfter, err := strconv.ParseFloat(v, 64); err == nil {
			// NaN, an infinity, or a wait past what a Duration holds has no meaningful
			// conversion (Go leaves it implementation-defined), so it is no answer at all.
			wait := retryAfter * float64(retry.units)
			if math.IsNaN(wait) || wait >= math.MaxInt64 || wait <= math.MinInt64 {
				continue
			}
			return time.Duration(wait), true
		}
		if d, ok := retry.custom(v); ok {
			return d, true
		}
	}

	return 0, false
}

// isBeforeContextDeadline reports whether the non-zero Time t is
// before ctx's deadline. If ctx does not have a deadline, it
// always reports true (the deadline is considered infinite).
func isBeforeContextDeadline(t time.Time, ctx context.Context) bool {
	d, ok := ctx.Deadline()
	if !ok {
		return true
	}
	return t.Before(d)
}

// bodyWithTimeout is an io.ReadCloser which can observe a context's cancel func
// to handle timeouts etc. It wraps an existing io.ReadCloser.
type bodyWithTimeout struct {
	stop func() // stops the time.Timer waiting to cancel the request
	rc   io.ReadCloser
}

func (b *bodyWithTimeout) Read(p []byte) (n int, err error) {
	n, err = b.rc.Read(p)
	if err == nil {
		return n, nil
	}
	if err == io.EOF {
		return n, err
	}
	return n, err
}

func (b *bodyWithTimeout) Close() error {
	err := b.rc.Close()
	b.stop()
	return err
}

// maxRetryAfter is the longest wait a `Retry-After`/`retry-after-ms` header can ask for and be
// honoured. The wait runs outside any per-attempt timeout, so a server asking for a day (or a date
// far ahead) would otherwise park the call for that long.
const maxRetryAfter = time.Minute

func retryDelay(res *http.Response, retryCount int) time.Duration {
	// If the backend tells us to wait a sane amount of time (positive, at most a minute), use that
	// value; anything else falls back to the exponential backoff below.
	if retryAfterDelay, ok := parseRetryAfterHeader(res); ok && retryAfterDelay > 0 && retryAfterDelay <= maxRetryAfter {
		return retryAfterDelay
	}

	maxDelay := 8 * time.Second
	delay := time.Duration(0.5 * float64(time.Second) * math.Pow(2, float64(retryCount)))
	if delay > maxDelay {
		delay = maxDelay
	}

	jitter := rand.Int63n(int64(delay / 4))
	delay -= time.Duration(jitter)
	return delay
}

// maxResponseBodySize is the most a response body read into memory may hold; a larger
// one fails the call instead of exhausting memory. A body handed to the caller unread (a
// download or a stream, decoded into an *http.Response) is not limited.
const maxResponseBodySize = 64 << 20

// maxErrorBodySize is how much of an error response's body is read: enough for any
// error document, and what is past it is dropped rather than read.
const maxErrorBodySize = 64 << 10

// defaultBodyReadTimeout bounds reading a body into memory when no per-attempt timeout
// is set. The default client bounds the wait for response headers the same way; without
// this, a server that sends headers and then stalls would hang the call forever. It
// starts once headers arrive, so a slow server-side operation is not cut short. A var
// rather than a const only so a test can shorten it.
var defaultBodyReadTimeout = 10 * time.Minute

// readBoundedBody reads a response body that is decoded in memory, refusing one past
// maxResponseBodySize.
func readBoundedBody(body io.Reader) ([]byte, error) {
	contents, err := io.ReadAll(io.LimitReader(body, maxResponseBodySize+1))
	if err != nil {
		return nil, err
	}
	if len(contents) > maxResponseBodySize {
		return nil, fmt.Errorf("the response body is larger than the %d MiB the SDK reads into memory", maxResponseBodySize>>20)
	}
	return contents, nil
}

// readErrorBody reads at most maxErrorBodySize of an error response's body.
func readErrorBody(body io.Reader) ([]byte, error) {
	return io.ReadAll(io.LimitReader(body, maxErrorBodySize))
}

// origin is the scheme, host and port of u, lower-cased and with the scheme's default
// port spelled out, so `https://API.example.com` and `https://api.example.com:443` are
// one origin and `http://host` and `http://host:8080` are two.
func origin(u *url.URL) string {
	scheme := strings.ToLower(u.Scheme)
	port := u.Port()
	if port == "" {
		switch scheme {
		case "http":
			port = "80"
		case "https":
			port = "443"
		}
	}
	return scheme + "://" + strings.ToLower(u.Hostname()) + ":" + port
}

// maxRedirects is how many redirects one request follows before giving up, the
// standard library's own limit.
const maxRedirects = 10

// withRedirectPolicy returns a shallow copy of client whose CheckRedirect runs
// checkRedirect on every hop. A CheckRedirect the caller set (a hop limit, a logger, or
// http.ErrUseLastResponse) still runs after it, on the request checkRedirect already
// stripped, and its answer stands: the caller's policy can follow fewer redirects than
// the SDK's, and is not handed the credentials the SDK's dropped.
func withRedirectPolicy(client *http.Client) *http.Client {
	if client == nil {
		client = http.DefaultClient
	}
	copied := *client
	caller := client.CheckRedirect
	if caller == nil {
		copied.CheckRedirect = checkRedirect
		return &copied
	}
	copied.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if err := checkRedirect(req, via); err != nil {
			return err
		}
		return caller(req, via)
	}
	return &copied
}

// followsSDKRedirectPolicy reports whether the request goes through checkRedirect, as
// opposed to an HTTPDoer of the caller's own.
func (cfg *RequestConfig) followsSDKRedirectPolicy() bool {
	return cfg.CustomHTTPDoer == nil
}

// checkRedirect is the redirect policy of every *http.Client the SDK sends through. It
// runs before each redirect is followed, after the standard client has copied the
// first request's headers onto req, the next request.
//
//   - It refuses a scheme other than http(s) and a step down from https to http, and
//     stops after maxRedirects.
//   - A 301 or 302 answering a method other than GET, HEAD or POST keeps its method (the
//     standard client would turn it into a GET that succeeds without doing what was
//     asked), and every hop that keeps such a method carries the first request's body,
//     a 307 after a 301 included; one whose body cannot be sent again is refused. 303,
//     and 301/302 answering a POST, stay the bodiless GET the client makes them, without
//     the first request's Content-* headers on that hop or any after it.
//   - Once any hop leaves the first request's origin, every credential header is dropped
//     from req (and the Referer, which would carry the first URL's query credentials),
//     for the rest of the chain. The standard client only drops Authorization and Cookie,
//     only for another host, and keeps them for another port or a downgrade; it never
//     knew about an API key in a header of its own. The redirect target's query is its
//     own, so no query credential follows across.
//   - While the chain stays on the first request's origin, the query credentials the
//     first request carried are kept on the target, as its headers are.
func checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= maxRedirects {
		return fmt.Errorf("stopped after %d redirects", maxRedirects)
	}
	scheme := strings.ToLower(req.URL.Scheme)
	if scheme != "http" && scheme != "https" {
		return fmt.Errorf("refusing to follow a redirect to a %q URL", req.URL.Scheme)
	}
	previous := via[len(via)-1]
	if strings.EqualFold(previous.URL.Scheme, "https") && scheme == "http" {
		return errors.New("refusing to follow a redirect from https to http")
	}

	first := via[0]
	if req.Response != nil && (req.Response.StatusCode == http.StatusMovedPermanently || req.Response.StatusCode == http.StatusFound) {
		switch previous.Method {
		case http.MethodGet, http.MethodHead, http.MethodPost:
		default:
			req.Method = previous.Method
		}
	}
	firstHadBody := first.Body != nil && first.Body != http.NoBody
	bodiless := req.Body == nil || req.Body == http.NoBody
	// A hop that keeps a method other than GET or HEAD needs the first request's body. The
	// standard client dropped it for the 301/302 kept above, and once one hop has dropped
	// it, it stays dropped on every hop after: a 307 after a 301 would send an empty PUT.
	if bodiless && firstHadBody && req.Method != http.MethodGet && req.Method != http.MethodHead {
		if first.GetBody == nil {
			// A body read from a stream has no GetBody, and a ContentLength of 0 there
			// means unknown, not empty: following would send the target an empty request.
			return errors.New("refusing to follow a redirect: the request body cannot be sent again")
		}
		body, err := first.GetBody()
		if err != nil {
			return err
		}
		req.Body, req.GetBody, req.ContentLength = body, first.GetBody, first.ContentLength
		bodiless = false
	}
	// A bodiless hop after a request that had a body carries none of its Content-* headers,
	// which the standard client copies from the first request onto every hop.
	if bodiless && firstHadBody {
		for name := range req.Header {
			if strings.HasPrefix(http.CanonicalHeaderKey(name), "Content-") {
				delete(req.Header, name)
			}
		}
	}

	home := origin(first.URL)
	crossed := origin(req.URL) != home
	for _, hop := range via[1:] {
		crossed = crossed || origin(hop.URL) != home
	}
	if crossed {
		for name := range req.Header {
			if isCredentialHeader(name) {
				delete(req.Header, name)
			}
		}
		return nil
	}
	// Same origin: the credential headers stay. The standard client already dropped
	// Authorization and Cookie when the host is spelled differently (`API.example.com:443`
	// for `api.example.com`), so they are taken back from the first request.
	for _, name := range []string{"Authorization", "Proxy-Authorization", "Cookie", "Cookie2"} {
		if values, ok := first.Header[name]; ok && len(req.Header.Values(name)) == 0 {
			req.Header[name] = values
		}
	}
	sent := first.URL.Query()
	target := req.URL.Query()
	kept := url.Values{}
	for _, name := range redact.SensitiveQueryParams {
		if values, ok := sent[name]; ok && !target.Has(name) {
			kept[name] = values
		}
	}
	// Appended, so the target's own query keeps its order.
	if len(kept) > 0 {
		if req.URL.RawQuery == "" {
			req.URL.RawQuery = kept.Encode()
		} else {
			req.URL.RawQuery += "&" + kept.Encode()
		}
	}
	return nil
}

// isCredentialHeader reports whether a header may carry a credential, and so must not
// follow a redirect to another origin.
func isCredentialHeader(name string) bool {
	switch http.CanonicalHeaderKey(name) {
	case "Authorization", "Proxy-Authorization", "Cookie", "Cookie2", "Referer":
		return true
	}
	return redact.IsSensitiveHeader(name)
}

func (cfg *RequestConfig) Execute() (err error) {
	if cfg.BaseURL == nil {
		if cfg.DefaultBaseURL != nil {
			cfg.BaseURL = cfg.DefaultBaseURL
		} else {
			return fmt.Errorf("requestconfig: base url is not set")
		}
	}

	cfg.Request.URL, err = cfg.BaseURL.Parse(strings.TrimLeft(cfg.Request.URL.String(), "/"))
	if err != nil {
		return err
	}
	// A path that resolves to an absolute URL (one passed to the raw `Client.Get`, or a path
	// parameter that starts with `scheme:`) would send the client's credentials to whatever
	// host it names. Only the base URL's origin is ever addressed; WithBaseURL changes it.
	if origin(cfg.Request.URL) != origin(cfg.BaseURL) {
		return fmt.Errorf("requestconfig: refusing to send a request to %s, which is not on the base URL's origin; set option.WithBaseURL to address another host", redact.URL(cfg.Request.URL))
	}

	if cfg.Body != nil && cfg.Request.Body == nil {
		switch body := cfg.Body.(type) {
		case *bytes.Buffer:
			b := body.Bytes()
			cfg.Request.ContentLength = int64(body.Len())
			cfg.Request.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(b)), nil }
			cfg.Request.Body, _ = cfg.Request.GetBody()
		case *bytes.Reader:
			cfg.Request.ContentLength = int64(body.Len())
			cfg.Request.GetBody = func() (io.ReadCloser, error) {
				_, err := body.Seek(0, 0)
				return io.NopCloser(body), err
			}
			cfg.Request.Body, _ = cfg.Request.GetBody()
		default:
			if rc, ok := body.(io.ReadCloser); ok {
				cfg.Request.Body = rc
			} else {
				cfg.Request.Body = io.NopCloser(body)
			}
		}
	}

	// Every *http.Client follows redirects through checkRedirect, a caller's own included,
	// and one with a CheckRedirect of its own runs that after it. An HTTPDoer that is not an
	// *http.Client follows redirects however it was built to (see option.WithHTTPClient).
	handler := withRedirectPolicy(cfg.HTTPClient).Do
	if cfg.CustomHTTPDoer != nil {
		handler = cfg.CustomHTTPDoer.Do
	}
	for i := len(cfg.Middlewares) - 1; i >= 0; i -= 1 {
		handler = applyMiddleware(cfg.Middlewares[i], handler)
	}

	// Don't send the current retry count in the headers if the caller modified the header defaults.
	shouldSendRetryCount := cfg.Request.Header.Get("X-Scalar-Retry-Count") == "0"

	// Whether the success body is read into memory here, rather than handed to the caller
	// unread (a download or a stream, decoded into an *http.Response).
	_, intoCustomResponseBody := cfg.ResponseBodyInto.(**http.Response)
	buffered := cfg.ResponseBodyInto != nil && !intoCustomResponseBody

	var res *http.Response
	// cancel releases the current attempt's context: one with the per-attempt timeout, or,
	// for a body read into memory, one that lets that read be bounded once the headers
	// arrive. An earlier attempt's is released before the next is sent, and the last one
	// when Execute returns, unless it is handed to the response body below. Any other call
	// gets no context of its own: one handed to a body the caller never closes would stay
	// registered on the caller's context.
	var cancel context.CancelFunc
	for retryCount := 0; retryCount <= cfg.MaxRetries; retryCount += 1 {
		ctx := cfg.Request.Context()
		if cfg.RequestTimeout != time.Duration(0) && isBeforeContextDeadline(time.Now().Add(cfg.RequestTimeout), ctx) {
			ctx, cancel = context.WithTimeout(ctx, cfg.RequestTimeout)
		} else if buffered {
			ctx, cancel = context.WithCancel(ctx)
		}
		if cancel != nil {
			defer func() {
				// The cancel function is nil if it was handed off to be handled in a different scope.
				if cancel != nil {
					cancel()
				}
			}()
		}

		req := cfg.Request.Clone(ctx)
		if shouldSendRetryCount {
			req.Header.Set("X-Scalar-Retry-Count", strconv.Itoa(retryCount))
		}

		res, err = handler(req)
		if ctx != nil && ctx.Err() != nil {
			return ctx.Err()
		}
		if !shouldRetry(cfg.Request, res) || retryCount >= cfg.MaxRetries {
			break
		}

		// Prepare next request and wait for the retry delay
		if cfg.Request.GetBody != nil {
			cfg.Request.Body, err = cfg.Request.GetBody()
			if err != nil {
				return err
			}
		}

		// Can't actually refresh the body, so we don't attempt to retry here
		if cfg.Request.GetBody == nil && cfg.Request.Body != nil {
			break
		}

		// Close the response body before retrying to prevent connection leaks
		if res != nil && res.Body != nil {
			_ = res.Body.Close()
		}
		// This attempt is over; release its context before waiting for the next.
		if cancel != nil {
			cancel()
			cancel = nil
		}

		// The wait is bounded by the caller's context, not the finished attempt's.
		select {
		case <-cfg.Request.Context().Done():
			return cfg.Request.Context().Err()
		case <-time.After(retryDelay(res, retryCount)):
		}
	}

	// Save *http.Response if it is requested to, even if there was an error making the request. This is
	// useful in cases where you might want to debug by inspecting the response. Note that if err != nil,
	// the response should be generally be empty, but there are edge cases.
	if cfg.ResponseInto != nil {
		*cfg.ResponseInto = res
	}
	if responseBodyInto, ok := cfg.ResponseBodyInto.(**http.Response); ok {
		*responseBodyInto = res
	}

	// If there was a connection error in the final request or any other transport error,
	// return that early without trying to coerce into an APIError. The standard client names
	// the full request URL in it, so its credentials are masked first.
	if err != nil {
		return redact.URLError(err)
	}

	// The standard client hands back a 307 or 308 it could not follow because the request
	// body was read from a stream and cannot be sent again. That is not the server's answer
	// to the request, so it is not returned as one.
	if (res.StatusCode == http.StatusTemporaryRedirect || res.StatusCode == http.StatusPermanentRedirect) &&
		res.Header.Get("Location") != "" && cfg.Request.Body != nil && cfg.Request.GetBody == nil && cfg.followsSDKRedirectPolicy() {
		_ = res.Body.Close()
		return fmt.Errorf("requestconfig: the server answered with a %d redirect, which cannot be followed because the request body was read from a stream and cannot be sent again", res.StatusCode)
	}

	readsBody := res.StatusCode >= 400 || buffered
	// A body read into memory is bounded in time as well as size. With a per-attempt
	// timeout, that timeout already covers it; without one, the read gets
	// defaultBodyReadTimeout from the moment the headers arrived. The read is cut short by
	// cancelling the attempt's context, or, for an error body of a call that has none (a
	// void call, or one returning the *http.Response), by closing the body under it.
	var bodyTimedOut atomic.Bool
	if readsBody && cfg.RequestTimeout == time.Duration(0) {
		stop := cancel
		if stop == nil {
			body := res.Body
			stop = func() { _ = body.Close() }
		}
		timer := time.AfterFunc(defaultBodyReadTimeout, func() {
			bodyTimedOut.Store(true)
			stop()
		})
		defer timer.Stop()
	}
	// bodyReadError explains a read the body deadline cut short.
	bodyReadError := func(err error) error {
		if bodyTimedOut.Load() {
			return fmt.Errorf("the response body was not read within %s: %w", defaultBodyReadTimeout, err)
		}
		return err
	}

	if res.StatusCode >= 400 {
		contents, err := readErrorBody(res.Body)
		_ = res.Body.Close()
		if err != nil {
			return bodyReadError(err)
		}

		// If there is an APIError, re-populate the response body so that debugging
		// utilities can conveniently dump the response without issue.
		res.Body = io.NopCloser(bytes.NewBuffer(contents))

		// Load the contents into the error format if it is provided.
		aerr := apierror.Error{Request: cfg.Request, Response: res, StatusCode: res.StatusCode}
		err = aerr.UnmarshalJSON(contents)
		if err != nil {
			return err
		}
		return &aerr
	}

	if !readsBody {
		// A body no one can reach (a void call that asked for no *http.Response) is closed
		// here, so its connection is released and the deferred cancel, when a per-attempt
		// timeout made one, frees its context.
		if cfg.ResponseInto == nil && !intoCustomResponseBody {
			_ = res.Body.Close()
			return nil
		}
		// We aren't reading the response body in this scope, but whoever is will need the
		// cancel func from the context to observe request timeouts.
		// Put the cancel function in the response body so it can be handled elsewhere.
		if cancel != nil {
			res.Body = &bodyWithTimeout{rc: res.Body, stop: cancel}
			cancel = nil
		}
		return nil
	}

	contents, err := readBoundedBody(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return fmt.Errorf("error reading response body: %w", bodyReadError(err))
	}
	if len(contents) == 0 {
		return nil
	}

	// If we are not json, return plaintext
	contentType := res.Header.Get("content-type")
	mediaType, _, _ := mime.ParseMediaType(contentType)
	isJSON := strings.Contains(mediaType, "application/json") || strings.HasSuffix(mediaType, "+json")
	if !isJSON {
		switch dst := cfg.ResponseBodyInto.(type) {
		case *string:
			*dst = string(contents)
		case **string:
			tmp := string(contents)
			*dst = &tmp
		case *[]byte:
			*dst = contents
		case **[]byte:
			tmp := contents
			*dst = &tmp
		default:
			return fmt.Errorf("expected destination type of 'string' or '[]byte' for responses with content-type '%s' that is not 'application/json'", contentType)
		}
		return nil
	}

	switch dst := cfg.ResponseBodyInto.(type) {
	// If the response happens to be a byte array, deserialize the body as-is.
	case *[]byte:
		*dst = contents
	case **[]byte:
		tmp := contents
		*dst = &tmp
	default:
		if isUnionDestination(cfg.ResponseBodyInto) && json.Valid(contents) && !bytes.Equal(bytes.TrimSpace(contents), []byte("null")) {
			if err = apijson.Unmarshal(contents, cfg.ResponseBodyInto); err != nil {
				return fmt.Errorf("error parsing response json: %w", err)
			}
			return nil
		}
		err = json.NewDecoder(bytes.NewReader(contents)).Decode(cfg.ResponseBodyInto)
		if err != nil {
			return fmt.Errorf("error parsing response json: %w", err)
		}
	}

	return nil
}

// PathEscapeReserved percent-encodes a path parameter while leaving `/` intact,
// implementing RFC 6570 reserved expansion (`{+var}`) for parameters the OpenAPI
// document flags with `allowReserved` — typically a file path or object key.
//
// The slash is part of the route shape for these values: `docs/example.txt` must
// stay nested under `.../files/docs/example.txt` rather than collapsing into a
// single `docs%2Fexample.txt` segment that addresses a different backend path.
// Every `/`-separated segment still goes through url.PathEscape, so all other
// URL-significant characters are encoded by exactly the same rules as a plain
// path param and a reserved value cannot smuggle in `?` or `#`.
func PathEscapeReserved(value string) string {
	segments := strings.Split(value, "/")
	for i, segment := range segments {
		segments[i] = url.PathEscape(segment)
	}
	return strings.Join(segments, "/")
}

// EncodeCookieValue percent-encodes a cookie parameter's value the way JavaScript's
// encodeURIComponent does: letters, digits and `-_.!~*'()` stay as they are, every other
// byte becomes `%XX`. A value is the caller's, and a raw `;`, `,`, space, quote, backslash
// or control character in it would end the cookie and start one of the caller's choosing.
func EncodeCookieValue(value string) string {
	const hexDigits = "0123456789ABCDEF"
	var encoded strings.Builder
	for i := 0; i < len(value); i++ {
		c := value[i]
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || strings.IndexByte("-_.!~*'()", c) >= 0 {
			encoded.WriteByte(c)
			continue
		}
		encoded.WriteByte('%')
		encoded.WriteByte(hexDigits[c>>4])
		encoded.WriteByte(hexDigits[c&0x0f])
	}
	return encoded.String()
}

// WithCookieParams returns a RequestOption that adds an operation's cookie parameters,
// each an already-encoded `name=value`, to the request's Cookie header. The cookies already
// there (a credential cookie the client set) are kept, except one with the same name as a
// parameter, which the parameter replaces.
func WithCookieParams(parts ...string) RequestOption {
	return RequestOptionFunc(func(r *RequestConfig) error {
		if len(parts) == 0 {
			return nil
		}
		names := make(map[string]bool, len(parts))
		for _, part := range parts {
			name, _, _ := strings.Cut(part, "=")
			names[name] = true
		}
		var cookies []string
		for _, header := range r.Request.Header.Values("Cookie") {
			for _, cookie := range strings.Split(header, ";") {
				cookie = strings.TrimSpace(cookie)
				name, _, _ := strings.Cut(cookie, "=")
				if cookie != "" && !names[name] {
					cookies = append(cookies, cookie)
				}
			}
		}
		r.Request.Header.Set("Cookie", strings.Join(append(cookies, parts...), "; "))
		return nil
	})
}

func ExecuteNewRequest(ctx context.Context, method string, u string, body interface{}, dst interface{}, opts ...RequestOption) error {
	cfg, err := NewRequestConfig(ctx, method, u, body, dst, opts...)
	if err != nil {
		return err
	}
	return cfg.Execute()
}

func (cfg *RequestConfig) Clone(ctx context.Context) *RequestConfig {
	if cfg == nil {
		return nil
	}
	req := cfg.Request.Clone(ctx)
	var err error
	if req.Body != nil {
		req.Body, err = req.GetBody()
	}
	if err != nil {
		return nil
	}
	new := &RequestConfig{
		MaxRetries:     cfg.MaxRetries,
		RequestTimeout: cfg.RequestTimeout,
		Context:        ctx,
		Request:        req,
		BaseURL:        cfg.BaseURL,
		DefaultBaseURL: cfg.DefaultBaseURL,
		// The caller's own HTTPDoer (a proxy, mTLS or request-signing client) sends every
		// request of the walk, not just the first page.
		CustomHTTPDoer: cfg.CustomHTTPDoer,
		HTTPClient:     cfg.HTTPClient,
		Middlewares:    cfg.Middlewares,
		APIKey:         cfg.APIKey,
		XAPIKey:        cfg.XAPIKey,
		AsBaseURL:      cfg.AsBaseURL,
		DedalusOrgID:   cfg.DedalusOrgID,
		BrowserToken:   cfg.BrowserToken,
	}
	// Request.Clone copies headers, so a non-GET clone would otherwise reuse the
	// original idempotency key and the server would dedupe the next request as a
	// replay (e.g. body-cursor pagination over POST). GET requests never carry a
	// key (see NewRequestConfig), so there is nothing to refresh.
	if req.Method != http.MethodGet {
		new.Request.Header.Set("Idempotency-Key", "scalar-go-"+newUUIDv4())
	}

	return new
}

// newUUIDv4 returns a random (version 4) UUID in its canonical lowercase
// 8-4-4-4-12 hex form, the value each idempotency key carries.
func newUUIDv4() string {
	var b [16]byte
	// crypto/rand, not math/rand: a predictable key could collide with, or be
	// guessed from, another request and be deduplicated as its replay.
	if _, err := cryptorand.Read(b[:]); err != nil {
		// The system randomness source failing is unrecoverable; panic rather than
		// send a guessable key.
		panic(err)
	}
	// RFC 9562: the high nibble of byte 6 is the version (4) and the top two bits
	// of byte 8 are the variant (10).
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func (cfg *RequestConfig) Apply(opts ...RequestOption) error {
	for _, opt := range opts {
		err := opt.Apply(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

// PreRequestOptions is used to collect all the options which need to be known before
// a call to [RequestConfig.ExecuteNewRequest], such as path parameters
// or global defaults.
// PreRequestOptions will return a [RequestConfig] with the options applied.
//
// Only request option functions of type [PreRequestOptionFunc] are applied.
func PreRequestOptions(opts ...RequestOption) (RequestConfig, error) {
	cfg := RequestConfig{}
	for _, opt := range opts {
		if opt, ok := opt.(PreRequestOptionFunc); ok {
			err := opt.Apply(&cfg)
			if err != nil {
				return cfg, err
			}
		}
	}
	return cfg, nil
}

// WithDefaultBaseURL returns a RequestOption that sets the client's default Base URL.
// This is always overridden by setting a base URL with WithBaseURL.
// WithBaseURL should be used instead of WithDefaultBaseURL except in internal code.
func WithDefaultBaseURL(baseURL string) RequestOption {
	u, err := url.Parse(baseURL)
	return RequestOptionFunc(func(r *RequestConfig) error {
		if err != nil {
			return err
		}
		r.DefaultBaseURL = u
		return nil
	})
}
