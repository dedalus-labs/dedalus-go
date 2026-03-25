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

// WorkspacePreviewService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspacePreviewService] method instead.
type WorkspacePreviewService struct {
	Options []option.RequestOption
}

// NewWorkspacePreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspacePreviewService(opts ...option.RequestOption) (r WorkspacePreviewService) {
	r = WorkspacePreviewService{}
	r.Options = opts
	return
}

// Create preview
func (r *WorkspacePreviewService) New(ctx context.Context, workspaceID string, body WorkspacePreviewNewParams, opts ...option.RequestOption) (res *Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/previews", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get preview
func (r *WorkspacePreviewService) Get(ctx context.Context, previewID string, query WorkspacePreviewGetParams, opts ...option.RequestOption) (res *Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if previewID == "" {
		err = errors.New("missing required preview_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/previews/%s", url.PathEscape(query.WorkspaceID), url.PathEscape(previewID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List previews
func (r *WorkspacePreviewService) List(ctx context.Context, workspaceID string, query WorkspacePreviewListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Preview], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/previews", url.PathEscape(workspaceID))
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

// List previews
func (r *WorkspacePreviewService) ListAutoPaging(ctx context.Context, workspaceID string, query WorkspacePreviewListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Preview] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Delete preview
func (r *WorkspacePreviewService) Delete(ctx context.Context, previewID string, body WorkspacePreviewDeleteParams, opts ...option.RequestOption) (res *Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if previewID == "" {
		err = errors.New("missing required preview_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/previews/%s", url.PathEscape(body.WorkspaceID), url.PathEscape(previewID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Preview struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Port      int64     `json:"port" api:"required"`
	PreviewID string    `json:"preview_id" api:"required"`
	// Any of "wake_in_progress", "ready", "closed", "expired", "failed".
	Status       PreviewStatus `json:"status" api:"required"`
	WorkspaceID  string        `json:"workspace_id" api:"required"`
	ErrorCode    string        `json:"error_code"`
	ErrorMessage string        `json:"error_message"`
	ExpiresAt    time.Time     `json:"expires_at" format:"date-time"`
	// Any of "http", "https".
	Protocol     PreviewProtocol `json:"protocol"`
	ReadyAt      time.Time       `json:"ready_at" format:"date-time"`
	RetryAfterMs int64           `json:"retry_after_ms"`
	URL          string          `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Port         respjson.Field
		PreviewID    respjson.Field
		Status       respjson.Field
		WorkspaceID  respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		ExpiresAt    respjson.Field
		Protocol     respjson.Field
		ReadyAt      respjson.Field
		RetryAfterMs respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Preview) RawJSON() string { return r.JSON.raw }
func (r *Preview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewStatus string

const (
	PreviewStatusWakeInProgress PreviewStatus = "wake_in_progress"
	PreviewStatusReady          PreviewStatus = "ready"
	PreviewStatusClosed         PreviewStatus = "closed"
	PreviewStatusExpired        PreviewStatus = "expired"
	PreviewStatusFailed         PreviewStatus = "failed"
)

type PreviewProtocol string

const (
	PreviewProtocolHTTP  PreviewProtocol = "http"
	PreviewProtocolHTTPS PreviewProtocol = "https"
)

// The property Port is required.
type PreviewCreateParams struct {
	Port int64 `json:"port" api:"required"`
	// Any of "http", "https".
	Protocol PreviewCreateParamsProtocol `json:"protocol,omitzero"`
	paramObj
}

func (r PreviewCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow PreviewCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewCreateParamsProtocol string

const (
	PreviewCreateParamsProtocolHTTP  PreviewCreateParamsProtocol = "http"
	PreviewCreateParamsProtocolHTTPS PreviewCreateParamsProtocol = "https"
)

type PreviewList struct {
	Items      []Preview `json:"items" api:"required"`
	NextCursor string    `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewList) RawJSON() string { return r.JSON.raw }
func (r *PreviewList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspacePreviewNewParams struct {
	PreviewCreateParams PreviewCreateParams
	paramObj
}

func (r WorkspacePreviewNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreviewCreateParams)
}
func (r *WorkspacePreviewNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PreviewCreateParams)
}

type WorkspacePreviewGetParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}

type WorkspacePreviewListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspacePreviewListParams]'s query parameters as
// `url.Values`.
func (r WorkspacePreviewListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspacePreviewDeleteParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}
