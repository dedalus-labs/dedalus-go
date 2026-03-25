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

// TerminalClientEventUnion contains all possible properties and values from
// [TerminalInputEvent], [TerminalResizeEvent].
//
// Use the [TerminalClientEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TerminalClientEventUnion struct {
	// This field is from variant [TerminalInputEvent].
	Data string `json:"data"`
	// Any of "input", "resize".
	Type string `json:"type"`
	// This field is from variant [TerminalResizeEvent].
	Height int64 `json:"height"`
	// This field is from variant [TerminalResizeEvent].
	Width int64 `json:"width"`
	JSON  struct {
		Data   respjson.Field
		Type   respjson.Field
		Height respjson.Field
		Width  respjson.Field
		raw    string
	} `json:"-"`
}

// anyTerminalClientEvent is implemented by each variant of
// [TerminalClientEventUnion] to add type safety for the return type of
// [TerminalClientEventUnion.AsAny]
type anyTerminalClientEvent interface {
	implTerminalClientEventUnion()
}

func (TerminalInputEvent) implTerminalClientEventUnion()  {}
func (TerminalResizeEvent) implTerminalClientEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TerminalClientEventUnion.AsAny().(type) {
//	case dedalus.TerminalInputEvent:
//	case dedalus.TerminalResizeEvent:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TerminalClientEventUnion) AsAny() anyTerminalClientEvent {
	switch u.Type {
	case "input":
		return u.AsInput()
	case "resize":
		return u.AsResize()
	}
	return nil
}

func (u TerminalClientEventUnion) AsInput() (v TerminalInputEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TerminalClientEventUnion) AsResize() (v TerminalResizeEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TerminalClientEventUnion) RawJSON() string { return u.JSON.raw }

func (r *TerminalClientEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TerminalClientEventUnion to a
// TerminalClientEventUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TerminalClientEventUnionParam.Overrides()
func (r TerminalClientEventUnion) ToParam() TerminalClientEventUnionParam {
	return param.Override[TerminalClientEventUnionParam](json.RawMessage(r.RawJSON()))
}

func TerminalClientEventParamOfInput(data string) TerminalClientEventUnionParam {
	var input TerminalInputEventParam
	input.Data = data
	return TerminalClientEventUnionParam{OfInput: &input}
}

func TerminalClientEventParamOfResize(height int64, type_ TerminalResizeEventType, width int64) TerminalClientEventUnionParam {
	var resize TerminalResizeEventParam
	resize.Height = height
	resize.Type = type_
	resize.Width = width
	return TerminalClientEventUnionParam{OfResize: &resize}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TerminalClientEventUnionParam struct {
	OfInput  *TerminalInputEventParam  `json:",omitzero,inline"`
	OfResize *TerminalResizeEventParam `json:",omitzero,inline"`
	paramUnion
}

func (u TerminalClientEventUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInput, u.OfResize)
}
func (u *TerminalClientEventUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[TerminalClientEventUnionParam](
		"type",
		apijson.Discriminator[TerminalInputEventParam]("input"),
		apijson.Discriminator[TerminalResizeEventParam]("resize"),
	)
}

type TerminalClosedEvent struct {
	// Any of "closed".
	Type TerminalClosedEventType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TerminalClosedEvent) RawJSON() string { return r.JSON.raw }
func (r *TerminalClosedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TerminalClosedEventType string

const (
	TerminalClosedEventTypeClosed TerminalClosedEventType = "closed"
)

// The properties Height, Width are required.
type TerminalCreateParams struct {
	Height int64             `json:"height" api:"required"`
	Width  int64             `json:"width" api:"required"`
	Cwd    param.Opt[string] `json:"cwd,omitzero"`
	Shell  param.Opt[string] `json:"shell,omitzero"`
	Env    map[string]string `json:"env,omitzero"`
	paramObj
}

func (r TerminalCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow TerminalCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TerminalCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TerminalErrorEvent struct {
	// Any of "error".
	Type         TerminalErrorEventType `json:"type" api:"required"`
	ErrorCode    string                 `json:"error_code"`
	ErrorMessage string                 `json:"error_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type         respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TerminalErrorEvent) RawJSON() string { return r.JSON.raw }
func (r *TerminalErrorEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TerminalErrorEventType string

const (
	TerminalErrorEventTypeError TerminalErrorEventType = "error"
)

type TerminalInputEvent struct {
	// Base64-encoded terminal input.
	Data string `json:"data" api:"required" format:"byte"`
	// Any of "input".
	Type TerminalInputEventType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TerminalInputEvent) RawJSON() string { return r.JSON.raw }
func (r *TerminalInputEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TerminalInputEvent to a TerminalInputEventParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TerminalInputEventParam.Overrides()
func (r TerminalInputEvent) ToParam() TerminalInputEventParam {
	return param.Override[TerminalInputEventParam](json.RawMessage(r.RawJSON()))
}

type TerminalInputEventType string

const (
	TerminalInputEventTypeInput TerminalInputEventType = "input"
)

// The properties Data, Type are required.
type TerminalInputEventParam struct {
	// Base64-encoded terminal input.
	Data string `json:"data" api:"required" format:"byte"`
	// Any of "input".
	Type TerminalInputEventType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r TerminalInputEventParam) MarshalJSON() (data []byte, err error) {
	type shadow TerminalInputEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TerminalInputEventParam) UnmarshalJSON(data []byte) error {
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

type TerminalOutputEvent struct {
	// Base64-encoded terminal output.
	Data string `json:"data" api:"required" format:"byte"`
	// Any of "output".
	Type TerminalOutputEventType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TerminalOutputEvent) RawJSON() string { return r.JSON.raw }
func (r *TerminalOutputEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TerminalOutputEventType string

const (
	TerminalOutputEventTypeOutput TerminalOutputEventType = "output"
)

type TerminalResizeEvent struct {
	Height int64 `json:"height" api:"required"`
	// Any of "resize".
	Type  TerminalResizeEventType `json:"type" api:"required"`
	Width int64                   `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Type        respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TerminalResizeEvent) RawJSON() string { return r.JSON.raw }
func (r *TerminalResizeEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TerminalResizeEvent to a TerminalResizeEventParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TerminalResizeEventParam.Overrides()
func (r TerminalResizeEvent) ToParam() TerminalResizeEventParam {
	return param.Override[TerminalResizeEventParam](json.RawMessage(r.RawJSON()))
}

type TerminalResizeEventType string

const (
	TerminalResizeEventTypeResize TerminalResizeEventType = "resize"
)

// The properties Height, Type, Width are required.
type TerminalResizeEventParam struct {
	Height int64 `json:"height" api:"required"`
	// Any of "resize".
	Type  TerminalResizeEventType `json:"type,omitzero" api:"required"`
	Width int64                   `json:"width" api:"required"`
	paramObj
}

func (r TerminalResizeEventParam) MarshalJSON() (data []byte, err error) {
	type shadow TerminalResizeEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TerminalResizeEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TerminalServerEventUnion contains all possible properties and values from
// [TerminalOutputEvent], [TerminalErrorEvent], [TerminalClosedEvent].
//
// Use the [TerminalServerEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TerminalServerEventUnion struct {
	// This field is from variant [TerminalOutputEvent].
	Data string `json:"data"`
	// Any of "output", "error", "closed".
	Type string `json:"type"`
	// This field is from variant [TerminalErrorEvent].
	ErrorCode string `json:"error_code"`
	// This field is from variant [TerminalErrorEvent].
	ErrorMessage string `json:"error_message"`
	JSON         struct {
		Data         respjson.Field
		Type         respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		raw          string
	} `json:"-"`
}

// anyTerminalServerEvent is implemented by each variant of
// [TerminalServerEventUnion] to add type safety for the return type of
// [TerminalServerEventUnion.AsAny]
type anyTerminalServerEvent interface {
	implTerminalServerEventUnion()
}

func (TerminalOutputEvent) implTerminalServerEventUnion() {}
func (TerminalErrorEvent) implTerminalServerEventUnion()  {}
func (TerminalClosedEvent) implTerminalServerEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TerminalServerEventUnion.AsAny().(type) {
//	case dedalus.TerminalOutputEvent:
//	case dedalus.TerminalErrorEvent:
//	case dedalus.TerminalClosedEvent:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TerminalServerEventUnion) AsAny() anyTerminalServerEvent {
	switch u.Type {
	case "output":
		return u.AsOutput()
	case "error":
		return u.AsError()
	case "closed":
		return u.AsClosed()
	}
	return nil
}

func (u TerminalServerEventUnion) AsOutput() (v TerminalOutputEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TerminalServerEventUnion) AsError() (v TerminalErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TerminalServerEventUnion) AsClosed() (v TerminalClosedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TerminalServerEventUnion) RawJSON() string { return u.JSON.raw }

func (r *TerminalServerEventUnion) UnmarshalJSON(data []byte) error {
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
