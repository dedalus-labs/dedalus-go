# Dedalus

This library provides convenient access to the Dedalus REST API from Go.

The full API of this library can be found in [api.md](./api.md).

<br />

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Reference](./api.md)
- [Authentication](#authentication)
- [Errors](#errors)
- [Client Options](#client-options)
- [Request Options](#request-options)
- [Retries and Timeouts](#retries-and-timeouts)
- [Pagination](#pagination)
- [Helpers](#helpers)
- [Logging](#logging)
- [Requirements](#requirements)

<br />

## Installation

```sh
go get github.com/dedalus-labs/dedalus-go
```

<br />

## Usage

```go
package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/option"
)

func main() {
	client := sdk.NewClient(
		option.WithXAPIKey(os.Getenv("DEDALUS_X_API_KEY")),
	)

	machine, err := client.Machines.New(context.Background(), sdk.MachineNewParams{
		CreateParams: sdk.CreateParams{
			Autosleep:  sdk.F[string]("300s"),
			MemoryMib:  sdk.F[int64](4096),
			StorageGib: sdk.F[int64](10),
			Vcpu:       sdk.F[float64](1),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(machine.MachineID)
}
```

The examples in the following sections assume a `client` configured as shown above.

See the [API reference](./api.md) for every available operation.

<br />

## Authentication

Pass credentials to the generated client constructor. Environment variables are read automatically when supported by the target runtime.

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `option.WithAPIKey` | `string \| provider` | - | Dedalus API key for Bearer token authentication. Defaults to DEDALUS_API_KEY. |
| `option.WithXAPIKey` | `string \| provider` | - | Dedalus API key for X-API-Key header authentication. Defaults to DEDALUS_X_API_KEY. |
| `option.WithAsBaseURL` | `string \| provider` | - | MCP Authorization Server URL. Defaults to DEDALUS_AS_URL. |
| `option.WithDedalusOrgID` | `string \| provider` | - | Organization ID for request scoping. Defaults to DEDALUS_ORG_ID. |

Declared schemes:

- `ApiKeyAuth` API key in header `x-api-key`
- `BearerAuth` bearer token

<br />

## Errors

Non-success responses return generated API errors. Error objects expose status, headers, response body, and request metadata where the target runtime supports it.

```go
machine, err := client.Machines.New(context.Background(), sdk.MachineNewParams{
	CreateParams: sdk.CreateParams{
		Autosleep:  sdk.F[string]("300s"),
		MemoryMib:  sdk.F[int64](4096),
		StorageGib: sdk.F[int64](10),
		Vcpu:       sdk.F[float64](1),
	},
})
if err != nil {
	var apiErr *sdk.Error
	if errors.As(err, &apiErr) {
		fmt.Println(apiErr.StatusCode, apiErr.RawJSON())
	}
	panic(err)
}

// imports: "context", "errors", "fmt", sdk "github.com/dedalus-labs/dedalus-go"
```

Documented error statuses: `401`, `403`, `409`, `429`, `503`, `default`.

<br />

## Client Options

Configure the generated client by setting any of these options when you create it.

```go
client := sdk.NewClient(
	option.WithBaseURL("https://api.example.com"),
	option.WithMaxRetries(2),
	option.WithRequestTimeout(60*time.Second),
)

// imports: sdk "github.com/dedalus-labs/dedalus-go", "github.com/dedalus-labs/dedalus-go/option", "time"
```

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `option.WithAPIKey` | `func(string) option.RequestOption` | `os.Getenv("DEDALUS_API_KEY")` | Dedalus API key for Bearer token authentication. |
| `option.WithXAPIKey` | `func(string) option.RequestOption` | `os.Getenv("DEDALUS_X_API_KEY")` | Dedalus API key for X-API-Key header authentication. |
| `option.WithAsBaseURL` | `func(string) option.RequestOption` | `os.Getenv("DEDALUS_AS_URL")` | MCP Authorization Server URL. |
| `option.WithDedalusOrgID` | `func(string) option.RequestOption` | `os.Getenv("DEDALUS_ORG_ID")` | Organization ID for request scoping. |
| `option.WithEnvironmentProduction` | `func() option.RequestOption` | - | Select the production API environment. |
| `option.WithBaseURL` | `func(string) option.RequestOption` | `os.Getenv("DEDALUS_BASE_URL")` | Override the default API base URL. |
| `option.WithRequestTimeout` | `func(time.Duration) option.RequestOption` | - | Maximum time to wait for each request attempt. |
| `option.WithMaxRetries` | `func(int) option.RequestOption` | `2` | Number of retries for temporary failures. |
| `option.WithHTTPClient` | `func(option.HTTPClient) option.RequestOption` | - | Custom HTTP client or transport implementation. |

<br />

## Request Options

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `option.WithHeader` | `func(string, string) option.RequestOption` | - | Set a per-request header. |
| `option.WithQuery` | `func(string, string) option.RequestOption` | - | Set a per-request query parameter. |
| `option.WithRequestBody` | `func(string, any) option.RequestOption` | - | Override the serialized request body and content type. |
| `option.WithResponseInto` | `func(**http.Response) option.RequestOption` | - | Capture the raw HTTP response. |
| `option.WithResponseBodyInto` | `func(any) option.RequestOption` | - | Override the response deserialization target. |

<br />

## Retries and Timeouts

Generated clients support request timeouts and retry temporary failures such as network errors, 408, 409, 429, and 5xx responses. Retry delays honor `Retry-After` headers when present. Tune the retry and timeout client options shown above, or override them per request.

<br />

## Pagination

List endpoints return paginated results you can iterate directly; the SDK fetches subsequent pages for you.

```go
page, err := client.Machines.List(context.Background(), sdk.MachineListParams{})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

<br />

## Helpers

- Pass `option.WithResponseInto(&raw)` to capture the underlying `*http.Response` for a request.
- Use the generated `String`, `Int`, `Bool`, `Float`, `Time`, `Opt`, and `Ptr` helpers when setting optional params.

<br />

## Logging

- Wrap the HTTP client with `option.WithMiddleware(...)` to add request logging or tracing.

<br />

## Requirements

- Go 1.22 or newer

Powered by Scalar.
