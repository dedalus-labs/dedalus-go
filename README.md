# Dedalus

This library provides convenient access to the Dedalus REST API from Go.

The full API of this library can be found in [api.md](./api.md).

<br />

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Image import](#image-import)
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

## Image import

Import and readiness polling are available in the `helpers/images` package:

```go
import (
    "context"

    sdk "github.com/dedalus-labs/dedalus-go"
    "github.com/dedalus-labs/dedalus-go/helpers/images"
)

func importArchive(ctx context.Context) (*images.Image, error) {
    client := sdk.NewClient()
    return images.Import(ctx, &client, "my-app.tar", images.ImportOptions{
        Name: "my-app",
        Platform: "linux/amd64",
    })
}
```

`Import` hashes and streams a `docker save` archive, registers it, and waits until
ready by default. Set `NoWait: true` to return after registration; call
`images.Poll(ctx, &client, image.ImageID, images.PollOptions{})` to wait later.

`ImportOptions` embeds `PollOptions`: `Timeout` defaults to one hour for the entire
import, `PollInterval` defaults to two seconds, and `OnProgress` receives each
observed image. A caller's earlier context deadline takes precedence. Cancellation
stops the client without deleting the image or cancelling accepted preparation.
`PreparationError.Image` preserves the image ID on preparation failure or interrupted
polling, and `errors.Is` recognizes context cancellation and deadline expiry.

Use the same `IdempotencyKey` when retrying unchanged input. The SDK's API retry
policy is retained. Archives must be regular, nonempty files of at most 5 GiB.
Storage uploads use a separate HTTP client without API credentials and reject
redirects. The image API must be deployed and enabled for your organization;
these helpers do not bypass access controls.

Custom image imports currently require an Ubuntu 24.04 `linux/amd64` Docker-save
archive. On Apple silicon, build or pull with `--platform linux/amd64` before
saving the archive. The image provides the filesystem; container entrypoint,
CMD, environment, working directory, and user settings are not applied. Start
processes using Machine exec or SSH after the Machine reaches running.

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
		option.WithBrowserToken(os.Getenv("DEDALUS_BROWSER_TOKEN")),
	)

	machine, err := client.Machines.New(context.Background(), sdk.MachineNewParams{
		CreateParams: sdk.CreateParams{
			Autosleep:  sdk.String("300s"),
			MemoryMiB:  sdk.Int(4096),
			StorageGiB: sdk.Int(10),
			VCPU:       sdk.Float(1),
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
| `option.WithAPIKey` | `func(string) option.RequestOption` | - | Dedalus API key for Bearer token authentication. Defaults to DEDALUS_API_KEY. |
| `option.WithXAPIKey` | `func(string) option.RequestOption` | - | Dedalus API key for X-API-Key header authentication. Defaults to DEDALUS_X_API_KEY. |
| `option.WithAsBaseURL` | `func(string) option.RequestOption` | - | MCP Authorization Server URL. Defaults to DEDALUS_AS_URL. |
| `option.WithDedalusOrgID` | `func(string) option.RequestOption` | - | Optional organization assertion. Must match the organization identified by the API credential. Defaults to DEDALUS_ORG_ID. |
| `option.WithBrowserToken` | `func(string) option.RequestOption` | - | Browser session access token managed by dedalus login. Defaults to DEDALUS_BROWSER_TOKEN. |

Declared schemes:

- `ApiKeyAuth` API key in header `x-api-key`
- `BearerAuth` bearer token
- `BrowserAuth` OAuth2/OpenID Connect

<br />

## Errors

Non-success responses return generated API errors. Error objects expose status, headers, response body, and request metadata where the target runtime supports it.

```go
machine, err := client.Machines.New(context.Background(), sdk.MachineNewParams{
	CreateParams: sdk.CreateParams{
		Autosleep:  sdk.String("300s"),
		MemoryMiB:  sdk.Int(4096),
		StorageGiB: sdk.Int(10),
		VCPU:       sdk.Float(1),
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

Documented error statuses: `400`, `401`, `402`, `403`, `404`, `409`, `422`, `429`, `500`, `503`, `504`, `default`.

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
| `option.WithDedalusOrgID` | `func(string) option.RequestOption` | `os.Getenv("DEDALUS_ORG_ID")` | Optional organization assertion. Must match the organization identified by the API credential. |
| `option.WithBrowserToken` | `func(string) option.RequestOption` | `os.Getenv("DEDALUS_BROWSER_TOKEN")` | Browser session access token managed by dedalus login. |
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

Paginated list methods return one page per call. A method that returns a page type has a sibling `<Method>AutoPaging` method returning an iterator that fetches the next page as the loop reaches it:

```go
iter := client.Machines.ListAutoPaging(context.Background(), sdk.MachineListParams{})
for iter.Next() {
	fmt.Printf("%+v\n", iter.Current())
}
if err := iter.Err(); err != nil {
	panic(err)
}
```

To walk the pages yourself, range over the page's items and advance it with `GetNextPage()`, which returns a nil page after the last one:

```go
page, err := client.Machines.List(context.Background(), sdk.MachineListParams{})
for page != nil {
	for _, item := range page.Items {
		fmt.Printf("%+v\n", item)
	}
	page, err = page.GetNextPage()
}
if err != nil {
	panic(err)
}
```

<br />

## Helpers

- Pass `option.WithResponseInto(&raw)` to capture the underlying `*http.Response` for a request.
- Use the generated `String`, `Int`, `Bool`, `Float`, `Time`, `Opt`, `Ptr`, `IntPtr`, `BoolPtr`, `FloatPtr`, `StringPtr`, `TimePtr`, and `File` helpers when setting optional params.

<br />

## Logging

- Wrap the HTTP client with `option.WithMiddleware(...)` to add request logging or tracing.

<br />

## Requirements

- Go 1.22 or newer

Powered by Scalar.
