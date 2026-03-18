// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/dedalus-labs/dedalus-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-go/internal/apiquery"
	shimjson "github.com/dedalus-labs/dedalus-go/internal/encoding/json"
	"github.com/dedalus-labs/dedalus-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-go/option"
	"github.com/dedalus-labs/dedalus-go/packages/pagination"
	"github.com/dedalus-labs/dedalus-go/packages/param"
	"github.com/dedalus-labs/dedalus-go/packages/respjson"
)

// WorkspaceTerminalService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceTerminalService] method instead.
type WorkspaceTerminalService struct {
	Options []option.RequestOption
}

// NewWorkspaceTerminalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceTerminalService(opts ...option.RequestOption) (r WorkspaceTerminalService) {
	r = WorkspaceTerminalService{}
	r.Options = opts
	return
}

// Create terminal
func (r *WorkspaceTerminalService) New(ctx context.Context, workspaceID string, body WorkspaceTerminalNewParams, opts ...option.RequestOption) (res *Terminal, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/terminals", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get terminal
func (r *WorkspaceTerminalService) Get(ctx context.Context, terminalID string, query WorkspaceTerminalGetParams, opts ...option.RequestOption) (res *Terminal, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if terminalID == "" {
		err = errors.New("missing required terminal_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/terminals/%s", url.PathEscape(query.WorkspaceID), url.PathEscape(terminalID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List terminals
func (r *WorkspaceTerminalService) List(ctx context.Context, workspaceID string, query WorkspaceTerminalListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Terminal], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/terminals", url.PathEscape(workspaceID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List terminals
func (r *WorkspaceTerminalService) ListAutoPaging(ctx context.Context, workspaceID string, query WorkspaceTerminalListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Terminal] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Delete terminal
func (r *WorkspaceTerminalService) Delete(ctx context.Context, terminalID string, body WorkspaceTerminalDeleteParams, opts ...option.RequestOption) (res *Terminal, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if terminalID == "" {
		err = errors.New("missing required terminal_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/terminals/%s", url.PathEscape(body.WorkspaceID), url.PathEscape(terminalID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Terminal struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Height    int64     `json:"height" api:"required"`
	// Any of "wake_in_progress", "ready", "closed", "expired", "failed".
	Status       TerminalStatus `json:"status" api:"required"`
	TerminalID   string         `json:"terminal_id" api:"required"`
	Width        int64          `json:"width" api:"required"`
	WorkspaceID  string         `json:"workspace_id" api:"required"`
	ErrorCode    string         `json:"error_code"`
	ErrorMessage string         `json:"error_message"`
	ExpiresAt    time.Time      `json:"expires_at" format:"date-time"`
	// Any of "websocket".
	Protocol     TerminalProtocol `json:"protocol"`
	ReadyAt      time.Time        `json:"ready_at" format:"date-time"`
	RetryAfterMs int64            `json:"retry_after_ms"`
	StreamURL    string           `json:"stream_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Height       respjson.Field
		Status       respjson.Field
		TerminalID   respjson.Field
		Width        respjson.Field
		WorkspaceID  respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		ExpiresAt    respjson.Field
		Protocol     respjson.Field
		ReadyAt      respjson.Field
		RetryAfterMs respjson.Field
		StreamURL    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Terminal) RawJSON() string { return r.JSON.raw }
func (r *Terminal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TerminalStatus string

const (
	TerminalStatusWakeInProgress TerminalStatus = "wake_in_progress"
	TerminalStatusReady          TerminalStatus = "ready"
	TerminalStatusClosed         TerminalStatus = "closed"
	TerminalStatusExpired        TerminalStatus = "expired"
	TerminalStatusFailed         TerminalStatus = "failed"
)

type TerminalProtocol string

const (
	TerminalProtocolWebsocket TerminalProtocol = "websocket"
)

// The properties Height, Width are required.
type TerminalCreateParams struct {
	Height       int64             `json:"height" api:"required"`
	Width        int64             `json:"width" api:"required"`
	Cwd          param.Opt[string] `json:"cwd,omitzero"`
	Shell        param.Opt[string] `json:"shell,omitzero"`
	WakeIfNeeded param.Opt[bool]   `json:"wake_if_needed,omitzero"`
	Env          map[string]string `json:"env,omitzero"`
	paramObj
}

func (r TerminalCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow TerminalCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TerminalCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TerminalList struct {
	Items      []Terminal `json:"items" api:"required"`
	NextCursor string     `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TerminalList) RawJSON() string { return r.JSON.raw }
func (r *TerminalList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceTerminalNewParams struct {
	TerminalCreateParams TerminalCreateParams
	paramObj
}

func (r WorkspaceTerminalNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TerminalCreateParams)
}
func (r *WorkspaceTerminalNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TerminalCreateParams)
}

type WorkspaceTerminalGetParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}

type WorkspaceTerminalListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceTerminalListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceTerminalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceTerminalDeleteParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}
