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
func (r *WorkspaceService) List(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) (res *pagination.WorkspaceList[WorkspaceListItem], err error) {
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
func (r *WorkspaceService) ListAutoPaging(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) *pagination.WorkspaceListAutoPager[WorkspaceListItem] {
	return pagination.NewWorkspaceListAutoPager(r.List(ctx, query, opts...))
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

// The properties ImageVersion, MemoryMiB, StorageGiB, VCPU are required.
type CreateParams struct {
	ImageVersion string `json:"image_version" api:"required"`
	// Memory in MiB.
	MemoryMiB  int64 `json:"memory_mib" api:"required"`
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
	ObservedRevision string    `json:"observed_revision" api:"required"`
	// Any of "accepted", "placement_pending", "starting", "running", "stopping",
	// "sleeping", "destroying", "destroyed", "failed".
	Phase             LifecycleStatusPhase `json:"phase" api:"required"`
	Reason            string               `json:"reason" api:"required"`
	Retryable         bool                 `json:"retryable" api:"required"`
	Revision          string               `json:"revision" api:"required"`
	AssignedHost      string               `json:"assigned_host"`
	LastError         string               `json:"last_error"`
	MemoryAssignedMiB int64                `json:"memory_assigned_mib"`
	MemoryResizeState string               `json:"memory_resize_state"`
	MemoryTargetMiB   int64                `json:"memory_target_mib"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastProgressAt    respjson.Field
		LastTransitionAt  respjson.Field
		ObservedRevision  respjson.Field
		Phase             respjson.Field
		Reason            respjson.Field
		Retryable         respjson.Field
		Revision          respjson.Field
		AssignedHost      respjson.Field
		LastError         respjson.Field
		MemoryAssignedMiB respjson.Field
		MemoryResizeState respjson.Field
		MemoryTargetMiB   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	MemoryMiB  param.Opt[int64] `json:"memory_mib,omitzero"`
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
	// Any of "active", "inactive", "destroyed".
	DesiredState WorkspaceDesiredState `json:"desired_state" api:"required"`
	// Memory in MiB.
	MemoryMiB  int64           `json:"memory_mib" api:"required"`
	Status     LifecycleStatus `json:"status" api:"required"`
	StorageGiB int64           `json:"storage_gib" api:"required"`
	// CPU in vCPUs.
	VCPU        float64 `json:"vcpu" api:"required"`
	WorkspaceID string  `json:"workspace_id" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema       string `json:"$schema" format:"uri"`
	ImageVersion string `json:"image_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DesiredState respjson.Field
		MemoryMiB    respjson.Field
		Status       respjson.Field
		StorageGiB   respjson.Field
		VCPU         respjson.Field
		WorkspaceID  respjson.Field
		Schema       respjson.Field
		ImageVersion respjson.Field
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
	WorkspaceDesiredStateActive    WorkspaceDesiredState = "active"
	WorkspaceDesiredStateInactive  WorkspaceDesiredState = "inactive"
	WorkspaceDesiredStateDestroyed WorkspaceDesiredState = "destroyed"
)

type WorkspaceList struct {
	Items []WorkspaceListItem `json:"items" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema     string `json:"$schema" format:"uri"`
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Schema      respjson.Field
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
	// Any of "active", "inactive", "destroyed".
	DesiredState string `json:"desired_state" api:"required"`
	// Memory in MiB.
	MemoryMiB  int64           `json:"memory_mib" api:"required"`
	Status     LifecycleStatus `json:"status" api:"required"`
	StorageGiB int64           `json:"storage_gib" api:"required"`
	// CPU in vCPUs.
	VCPU         float64 `json:"vcpu" api:"required"`
	WorkspaceID  string  `json:"workspace_id" api:"required"`
	ImageVersion string  `json:"image_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		DesiredState respjson.Field
		MemoryMiB    respjson.Field
		Status       respjson.Field
		StorageGiB   respjson.Field
		VCPU         respjson.Field
		WorkspaceID  respjson.Field
		ImageVersion respjson.Field
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
