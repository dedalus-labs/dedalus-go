---
name: dedalus-go-sdk
description: "Go SDK for Dedalus API. Use when writing Go code that calls Dedalus API with the github.com/dedalus-labs/dedalus-go package: installing it, constructing and authenticating the client, and calling API operations."
---

# Dedalus Go SDK

Generated Go client for Dedalus API, published as `github.com/dedalus-labs/dedalus-go`. Use the generated client instead of hand-writing HTTP requests.

## Install

```sh
go get github.com/dedalus-labs/dedalus-go
```

## Client setup and authentication

```go
import (
	"context"
	"fmt"

	sdk "github.com/dedalus-labs/dedalus-go"
)

client := sdk.NewClient()
```

Provide credentials using the options below. Environment variables are read automatically when the target runtime supports them:

- `option.WithAPIKey` (env: `DEDALUS_API_KEY`) — Dedalus API key for Bearer token authentication.
- `option.WithXAPIKey` (env: `DEDALUS_X_API_KEY`) — Dedalus API key for X-API-Key header authentication.
- `option.WithAsBaseURL` (env: `DEDALUS_AS_URL`) — MCP Authorization Server URL.
- `option.WithDedalusOrgID` (env: `DEDALUS_ORG_ID`) — Organization ID for request scoping.

## Calling operations

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

Method names, parameter shapes, and response types are generated from the API description — do not guess them. Look up the exact call signature in [api.md](./api.md) before writing a call.

## Pagination

List endpoints return paginated results you can iterate directly; the SDK fetches subsequent pages for you.

```go
page, err := client.Machines.List(context.Background(), sdk.MachineListParams{})
if err != nil {
	panic(err)
}

fmt.Println(page)
```

## Error handling

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

## Requirements

- Go 1.22 or newer

## Reference files

- [README.md](./README.md) — full feature tour: client options, request options, retries and timeouts, logging.
- [api.md](./api.md) — complete catalogue of every operation with request and response types.
