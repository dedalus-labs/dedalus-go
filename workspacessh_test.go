// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalus_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-go/option"
)

func TestWorkspaceSSHNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workspaces.SSH.New(
		context.TODO(),
		"workspace_id",
		dedalus.WorkspaceSSHNewParams{
			SSHSessionCreateParams: dedalus.SSHSessionCreateParams{
				PublicKey:    "public_key",
				WakeIfNeeded: dedalus.Bool(true),
			},
		},
	)
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkspaceSSHGet(t *testing.T) {
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
	_, err := client.Workspaces.SSH.Get(
		context.TODO(),
		"session_id",
		dedalus.WorkspaceSSHGetParams{
			WorkspaceID: "workspace_id",
		},
	)
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkspaceSSHListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workspaces.SSH.List(
		context.TODO(),
		"workspace_id",
		dedalus.WorkspaceSSHListParams{
			Cursor: dedalus.String("cursor"),
			Limit:  dedalus.Int(0),
		},
	)
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkspaceSSHDelete(t *testing.T) {
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
	_, err := client.Workspaces.SSH.Delete(
		context.TODO(),
		"session_id",
		dedalus.WorkspaceSSHDeleteParams{
			WorkspaceID: "workspace_id",
		},
	)
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
