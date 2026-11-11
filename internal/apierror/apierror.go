// File generated from our OpenAPI spec by Scalar. See README.md for details.

package apierror

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/dedalus-labs/dedalus-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-go/internal/redact"
	"github.com/dedalus-labs/dedalus-go/packages/respjson"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	// The URL can carry a query credential or userinfo, and the body is whatever the server sent:
	// both are masked, and the body is cut and its control characters escaped, so printing the
	// error cannot leak a credential or forge log lines. RawJSON() keeps the body as received.
	return fmt.Sprintf("%s %q: %d %s %s", r.Request.Method, redact.URL(r.Request.URL), r.Response.StatusCode, http.StatusText(r.Response.StatusCode), redact.MessageBody(r.JSON.raw))
}

// DumpRequest returns the request that failed as it went out, with its credential headers
// replaced and the credential values in its URL and body masked. The request itself is left as it
// was sent.
func (r *Error) DumpRequest(body bool) []byte {
	// A fresh context: the request's own is usually cancelled by the time the error is read, and
	// the dump runs through a transport of its own that would refuse a cancelled request.
	request := r.Request.Clone(context.Background())
	request.Header = redact.Headers(r.Request.Header)
	// The dump goes through a transport, which would turn URL userinfo into an Authorization header.
	request.URL.User = nil
	if r.Request.GetBody != nil {
		request.Body, _ = r.Request.GetBody()
	}
	out, _ := httputil.DumpRequestOut(request, body)
	return redact.Text(out)
}

// DumpResponse returns the error response, with its credential headers (a `Set-Cookie`, say)
// replaced and any credential value it echoes masked.
func (r *Error) DumpResponse(body bool) []byte {
	response := *r.Response
	response.Header = redact.Headers(r.Response.Header)
	out, _ := httputil.DumpResponse(&response, body)
	if body {
		// DumpResponse drained the body into the dump; hand the copy's replacement back so the
		// error's response can be dumped again.
		r.Response.Body = response.Body
	}
	return redact.Text(out)
}
