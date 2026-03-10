// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalus_test

import (
	"context"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-go/option"
)

func TestManualPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalus.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	page, err := client.Workspaces.List(context.TODO(), dedalus.WorkspaceListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, workspace := range page.Items {
		t.Logf("%+v\n", workspace.WorkspaceID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, workspace := range page.Items {
			t.Logf("%+v\n", workspace.WorkspaceID)
		}
	}
}
