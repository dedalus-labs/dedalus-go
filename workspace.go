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
	"github.com/dedalus-labs/dedalus-go/packages/ssestream"
)

// WorkspaceService contains methods and other services that help with interacting
// with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceService] method instead.
type WorkspaceService struct {
	Options    []option.RequestOption
	Artifacts  WorkspaceArtifactService
	Previews   WorkspacePreviewService
	SSH        WorkspaceSSHService
	Executions WorkspaceExecutionService
	Terminals  WorkspaceTerminalService
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r WorkspaceService) {
	r = WorkspaceService{}
	r.Options = opts
	r.Artifacts = NewWorkspaceArtifactService(opts...)
	r.Previews = NewWorkspacePreviewService(opts...)
	r.SSH = NewWorkspaceSSHService(opts...)
	r.Executions = NewWorkspaceExecutionService(opts...)
	r.Terminals = NewWorkspaceTerminalService(opts...)
	return
}

// Create workspace
func (r *WorkspaceService) New(ctx context.Context, body WorkspaceNewParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get workspace
func (r *WorkspaceService) Get(ctx context.Context, workspaceID string, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update workspace
func (r *WorkspaceService) Update(ctx context.Context, workspaceID string, params WorkspaceUpdateParams, opts ...option.RequestOption) (res *Workspace, err error) {
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List workspaces
func (r *WorkspaceService) List(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[WorkspaceListItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workspaces"
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

// List workspaces
func (r *WorkspaceService) ListAutoPaging(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[WorkspaceListItem] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Destroy workspace
func (r *WorkspaceService) Delete(ctx context.Context, workspaceID string, body WorkspaceDeleteParams, opts ...option.RequestOption) (res *Workspace, err error) {
	if !param.IsOmitted(body.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", body.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Streams workspace lifecycle updates over Server-Sent Events. Each `status` event
// contains a full `LifecycleResponse` payload. The stream closes after the
// workspace reaches its current desired state.
func (r *WorkspaceService) WatchStreaming(ctx context.Context, workspaceID string, query WorkspaceWatchParams, opts ...option.RequestOption) (stream *ssestream.Stream[Workspace]) {
	var (
		raw *http.Response
		err error
	)
	if !param.IsOmitted(query.LastEventID) {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", query.LastEventID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return ssestream.NewStream[Workspace](nil, err)
	}
	path := fmt.Sprintf("v1/workspaces/%s/status/stream", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[Workspace](ssestream.NewDecoder(raw), err)
}

// The properties MemoryMiB, StorageGiB, VCPU are required.
type CreateParams struct {
	// Memory in MiB.
	MemoryMiB int64 `json:"memory_mib" api:"required"`
	// Storage in GiB.
	StorageGiB int64 `json:"storage_gib" api:"required"`
	// CPU in vCPUs.
	VCPU float64 `json:"vcpu" api:"required"`
	paramObj
}

func (r CreateParams) MarshalJSON() (data []byte, err error) {
	type shadow CreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecycleStatus struct {
	LastProgressAt   time.Time `json:"last_progress_at" api:"required" format:"date-time"`
	LastTransitionAt time.Time `json:"last_transition_at" api:"required" format:"date-time"`
	// Any of "accepted", "placement_pending", "starting", "running", "stopping",
	// "sleeping", "destroying", "destroyed", "failed".
	Phase     LifecycleStatusPhase `json:"phase" api:"required"`
	Reason    string               `json:"reason" api:"required"`
	Retryable bool                 `json:"retryable" api:"required"`
	Revision  string               `json:"revision" api:"required"`
	LastError string               `json:"last_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastProgressAt   respjson.Field
		LastTransitionAt respjson.Field
		Phase            respjson.Field
		Reason           respjson.Field
		Retryable        respjson.Field
		Revision         respjson.Field
		LastError        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LifecycleStatus) RawJSON() string { return r.JSON.raw }
func (r *LifecycleStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LifecycleStatusPhase string

const (
	LifecycleStatusPhaseAccepted         LifecycleStatusPhase = "accepted"
	LifecycleStatusPhasePlacementPending LifecycleStatusPhase = "placement_pending"
	LifecycleStatusPhaseStarting         LifecycleStatusPhase = "starting"
	LifecycleStatusPhaseRunning          LifecycleStatusPhase = "running"
	LifecycleStatusPhaseStopping         LifecycleStatusPhase = "stopping"
	LifecycleStatusPhaseSleeping         LifecycleStatusPhase = "sleeping"
	LifecycleStatusPhaseDestroying       LifecycleStatusPhase = "destroying"
	LifecycleStatusPhaseDestroyed        LifecycleStatusPhase = "destroyed"
	LifecycleStatusPhaseFailed           LifecycleStatusPhase = "failed"
)

type UpdateParams struct {
	// Memory in MiB.
	MemoryMiB param.Opt[int64] `json:"memory_mib,omitzero"`
	// Storage in GiB.
	StorageGiB param.Opt[int64] `json:"storage_gib,omitzero"`
	// CPU in vCPUs.
	VCPU param.Opt[float64] `json:"vcpu,omitzero"`
	paramObj
}

func (r UpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Workspace struct {
	// Any of "running", "sleeping", "destroyed".
	DesiredState WorkspaceDesiredState `json:"desired_state" api:"required"`
	// Memory in MiB.
	MemoryMiB  int64           `json:"memory_mib" api:"required"`
	Status     LifecycleStatus `json:"status" api:"required"`
	StorageGiB int64           `json:"storage_gib" api:"required"`
	// CPU in vCPUs.
	VCPU        float64 `json:"vcpu" api:"required"`
	WorkspaceID string  `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DesiredState respjson.Field
		MemoryMiB    respjson.Field
		Status       respjson.Field
		StorageGiB   respjson.Field
		VCPU         respjson.Field
		WorkspaceID  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workspace) RawJSON() string { return r.JSON.raw }
func (r *Workspace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceDesiredState string

const (
	WorkspaceDesiredStateRunning   WorkspaceDesiredState = "running"
	WorkspaceDesiredStateSleeping  WorkspaceDesiredState = "sleeping"
	WorkspaceDesiredStateDestroyed WorkspaceDesiredState = "destroyed"
)

type WorkspaceList struct {
	Items      []WorkspaceListItem `json:"items" api:"required"`
	NextCursor string              `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceList) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceListItem struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "running", "sleeping", "destroyed".
	DesiredState string `json:"desired_state" api:"required"`
	// Memory in MiB.
	MemoryMiB  int64           `json:"memory_mib" api:"required"`
	Status     LifecycleStatus `json:"status" api:"required"`
	StorageGiB int64           `json:"storage_gib" api:"required"`
	// CPU in vCPUs.
	VCPU        float64 `json:"vcpu" api:"required"`
	WorkspaceID string  `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		DesiredState respjson.Field
		MemoryMiB    respjson.Field
		Status       respjson.Field
		StorageGiB   respjson.Field
		VCPU         respjson.Field
		WorkspaceID  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceListItem) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceNewParams struct {
	CreateParams CreateParams
	paramObj
}

func (r WorkspaceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateParams)
}
func (r *WorkspaceNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateParams)
}

type WorkspaceUpdateParams struct {
	UpdateParams UpdateParams
	IfMatch      string `header:"If-Match" api:"required" json:"-"`
	paramObj
}

func (r WorkspaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateParams)
}
func (r *WorkspaceUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateParams)
}

type WorkspaceListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceListParams]'s query parameters as `url.Values`.
func (r WorkspaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceDeleteParams struct {
	IfMatch string `header:"If-Match" api:"required" json:"-"`
	paramObj
}

type WorkspaceWatchParams struct {
	LastEventID param.Opt[string] `header:"Last-Event-ID,omitzero" json:"-"`
	paramObj
}
