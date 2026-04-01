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

func TestMachinePreviewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Machines.Previews.New(context.TODO(), dedalus.MachinePreviewNewParams{
		MachineID: "machine_id",
		PreviewCreateParams: dedalus.PreviewCreateParams{
			Port:       0,
			Protocol:   dedalus.PreviewCreateParamsProtocolHTTP,
			Visibility: dedalus.PreviewCreateParamsVisibilityPublic,
		},
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMachinePreviewGet(t *testing.T) {
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
	_, err := client.Machines.Previews.Get(context.TODO(), dedalus.MachinePreviewGetParams{
		MachineID: "machine_id",
		PreviewID: "preview_id",
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMachinePreviewListWithOptionalParams(t *testing.T) {
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
	_, err := client.Machines.Previews.List(context.TODO(), dedalus.MachinePreviewListParams{
		MachineID: "machine_id",
		Cursor:    dedalus.String("cursor"),
		Limit:     dedalus.Int(0),
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMachinePreviewDelete(t *testing.T) {
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
	_, err := client.Machines.Previews.Delete(context.TODO(), dedalus.MachinePreviewDeleteParams{
		MachineID: "machine_id",
		PreviewID: "preview_id",
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
