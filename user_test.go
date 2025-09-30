// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalusgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-go/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalusgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.New(context.TODO(), dedalusgo.UserNewParams{
		User: dedalusgo.UserParam{
			ID:         dedalusgo.Int(10),
			Email:      dedalusgo.String("john@email.com"),
			FirstName:  dedalusgo.String("John"),
			LastName:   dedalusgo.String("James"),
			Password:   dedalusgo.String("12345"),
			Phone:      dedalusgo.String("12345"),
			Username:   dedalusgo.String("theUser"),
			UserStatus: dedalusgo.Int(1),
		},
	})
	if err != nil {
		var apierr *dedalusgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalusgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.Get(context.TODO(), "username")
	if err != nil {
		var apierr *dedalusgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalusgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Users.Update(
		context.TODO(),
		"username",
		dedalusgo.UserUpdateParams{
			User: dedalusgo.UserParam{
				ID:         dedalusgo.Int(10),
				Email:      dedalusgo.String("john@email.com"),
				FirstName:  dedalusgo.String("John"),
				LastName:   dedalusgo.String("James"),
				Password:   dedalusgo.String("12345"),
				Phone:      dedalusgo.String("12345"),
				Username:   dedalusgo.String("theUser"),
				UserStatus: dedalusgo.Int(1),
			},
		},
	)
	if err != nil {
		var apierr *dedalusgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalusgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Users.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *dedalusgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalusgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.NewWithList(context.TODO(), dedalusgo.UserNewWithListParams{
		Items: []dedalusgo.UserParam{{
			ID:         dedalusgo.Int(10),
			Email:      dedalusgo.String("john@email.com"),
			FirstName:  dedalusgo.String("John"),
			LastName:   dedalusgo.String("James"),
			Password:   dedalusgo.String("12345"),
			Phone:      dedalusgo.String("12345"),
			Username:   dedalusgo.String("theUser"),
			UserStatus: dedalusgo.Int(1),
		}},
	})
	if err != nil {
		var apierr *dedalusgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalusgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.Login(context.TODO(), dedalusgo.UserLoginParams{
		Password: dedalusgo.String("password"),
		Username: dedalusgo.String("username"),
	})
	if err != nil {
		var apierr *dedalusgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalusgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Users.Logout(context.TODO())
	if err != nil {
		var apierr *dedalusgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
