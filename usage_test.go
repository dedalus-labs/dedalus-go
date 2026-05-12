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

func TestUsageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Usage.Get(context.TODO(), dedalus.UsageGetParams{
		PeriodStart: dedalus.String("period_start"),
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUsageMachineComputeWithOptionalParams(t *testing.T) {
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
	_, err := client.Usage.MachineCompute(context.TODO(), dedalus.UsageMachineComputeParams{
		Granularity: dedalus.String("granularity"),
		MachineID:   dedalus.String("machine_id"),
		PeriodEnd:   dedalus.String("period_end"),
		PeriodStart: dedalus.String("period_start"),
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUsageMachineStorageWithOptionalParams(t *testing.T) {
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
	_, err := client.Usage.MachineStorage(context.TODO(), dedalus.UsageMachineStorageParams{
		MachineID:   dedalus.String("machine_id"),
		PeriodEnd:   dedalus.String("period_end"),
		PeriodStart: dedalus.String("period_start"),
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
