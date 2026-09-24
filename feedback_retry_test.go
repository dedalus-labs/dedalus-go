package dedalus_test

import (
	"context"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"

	sdk "github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/option"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestInvariantAutomaticRetriesPreserveValidFeedbackIdentity(t *testing.T) {
	keys := []string{}
	pattern := regexp.MustCompile(`^[0-9a-f]{12}7[0-9a-f]{3}[89ab][0-9a-f]{15}$`)
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		key := request.Header.Get("Idempotency-Key")
		if !pattern.MatchString(key) {
			t.Fatalf("invalid automatic key: %q", key)
		}
		keys = append(keys, key)
		status := 201
		if len(keys) == 1 {
			status = 503
		}
		body := `{"id":"fb_` + key + `","source":"cli","reported_request_id":null,"debug":{"included":false}}`
		return &http.Response{
			StatusCode: status,
			Header:     http.Header{"Content-Type": {"application/json"}, "Retry-After-Ms": {"1"}},
			Body:       io.NopCloser(strings.NewReader(body)),
			Request:    request,
		}, nil
	})
	client := sdk.NewClient(
		option.WithBaseURL("https://feedback.invalid"), option.WithAPIKey("fixture"),
		option.WithMaxRetries(1), option.WithHTTPClient(&http.Client{Transport: transport}),
	)
	body := map[string]string{"message": "fixture", "source": "cli"}
	for i := 0; i < 2; i++ {
		if err := client.Post(context.Background(), "/v1/feedback", body, nil); err != nil {
			t.Fatal(err)
		}
	}
	if len(keys) != 3 {
		t.Fatalf("requests: got %d, want 3", len(keys))
	}
	if keys[0] != keys[1] {
		t.Fatal("retry changed key")
	}
	if keys[1] == keys[2] {
		t.Fatal("new submission reused key")
	}
}
