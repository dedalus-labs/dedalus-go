// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalus

import (
	"context"
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

// MachineService contains methods and other services that help with interacting
// with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMachineService] method instead.
type MachineService struct {
	Options    []option.RequestOption
	Artifacts  MachineArtifactService
	Previews   MachinePreviewService
	SSH        MachineSSHService
	Executions MachineExecutionService
	Terminals  MachineTerminalService
}

// NewMachineService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMachineService(opts ...option.RequestOption) (r MachineService) {
	r = MachineService{}
	r.Options = opts
	r.Artifacts = NewMachineArtifactService(opts...)
	r.Previews = NewMachinePreviewService(opts...)
	r.SSH = NewMachineSSHService(opts...)
	r.Executions = NewMachineExecutionService(opts...)
	r.Terminals = NewMachineTerminalService(opts...)
	return
}

// Create machine
func (r *MachineService) New(ctx context.Context, body MachineNewParams, opts ...option.RequestOption) (res *Machine, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/machines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get machine
func (r *MachineService) Get(ctx context.Context, query MachineGetParams, opts ...option.RequestOption) (res *Machine, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s", url.PathEscape(query.MachineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update machine
func (r *MachineService) Update(ctx context.Context, params MachineUpdateParams, opts ...option.RequestOption) (res *Machine, err error) {
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s", url.PathEscape(params.MachineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List machines
func (r *MachineService) List(ctx context.Context, query MachineListParams, opts ...option.RequestOption) (res *pagination.CursorPage[MachineListItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/machines"
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

// List machines
func (r *MachineService) ListAutoPaging(ctx context.Context, query MachineListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[MachineListItem] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Destroy machine
func (r *MachineService) Delete(ctx context.Context, params MachineDeleteParams, opts ...option.RequestOption) (res *Machine, err error) {
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s", url.PathEscape(params.MachineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Streams machine lifecycle updates over Server-Sent Events. Each `status` event
// contains a full `LifecycleResponse` payload. The stream closes after the machine
// reaches its current desired state.
func (r *MachineService) WatchStreaming(ctx context.Context, params MachineWatchParams, opts ...option.RequestOption) (stream *ssestream.Stream[Machine]) {
	var (
		raw *http.Response
		err error
	)
	if !param.IsOmitted(params.LastEventID) {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", params.LastEventID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return ssestream.NewStream[Machine](nil, err)
	}
	path := fmt.Sprintf("v1/machines/%s/status/stream", url.PathEscape(params.MachineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[Machine](ssestream.NewDecoder(raw), err)
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

type Machine struct {
	// Any of "running", "sleeping", "destroyed".
	DesiredState MachineDesiredState `json:"desired_state" api:"required"`
	MachineID    string              `json:"machine_id" api:"required"`
	// Memory in MiB.
	MemoryMiB  int64           `json:"memory_mib" api:"required"`
	Status     LifecycleStatus `json:"status" api:"required"`
	StorageGiB int64           `json:"storage_gib" api:"required"`
	// CPU in vCPUs.
	VCPU float64 `json:"vcpu" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DesiredState respjson.Field
		MachineID    respjson.Field
		MemoryMiB    respjson.Field
		Status       respjson.Field
		StorageGiB   respjson.Field
		VCPU         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Machine) RawJSON() string { return r.JSON.raw }
func (r *Machine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineDesiredState string

const (
	MachineDesiredStateRunning   MachineDesiredState = "running"
	MachineDesiredStateSleeping  MachineDesiredState = "sleeping"
	MachineDesiredStateDestroyed MachineDesiredState = "destroyed"
)

type MachineList struct {
	Items      []MachineListItem `json:"items" api:"required"`
	NextCursor string            `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineList) RawJSON() string { return r.JSON.raw }
func (r *MachineList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineListItem struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "running", "sleeping", "destroyed".
	DesiredState MachineListItemDesiredState `json:"desired_state" api:"required"`
	MachineID    string                      `json:"machine_id" api:"required"`
	// Memory in MiB.
	MemoryMiB  int64           `json:"memory_mib" api:"required"`
	Status     LifecycleStatus `json:"status" api:"required"`
	StorageGiB int64           `json:"storage_gib" api:"required"`
	// CPU in vCPUs.
	VCPU float64 `json:"vcpu" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		DesiredState respjson.Field
		MachineID    respjson.Field
		MemoryMiB    respjson.Field
		Status       respjson.Field
		StorageGiB   respjson.Field
		VCPU         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineListItem) RawJSON() string { return r.JSON.raw }
func (r *MachineListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineListItemDesiredState string

const (
	MachineListItemDesiredStateRunning   MachineListItemDesiredState = "running"
	MachineListItemDesiredStateSleeping  MachineListItemDesiredState = "sleeping"
	MachineListItemDesiredStateDestroyed MachineListItemDesiredState = "destroyed"
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

type MachineNewParams struct {
	CreateParams CreateParams
	paramObj
}

func (r MachineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateParams)
}
func (r *MachineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineGetParams struct {
	MachineID string `path:"machine_id" api:"required" json:"-"`
	paramObj
}

type MachineUpdateParams struct {
	MachineID    string `path:"machine_id" api:"required" json:"-"`
	UpdateParams UpdateParams
	IfMatch      string `header:"If-Match" api:"required" json:"-"`
	paramObj
}

func (r MachineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateParams)
}
func (r *MachineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MachineListParams]'s query parameters as `url.Values`.
func (r MachineListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MachineDeleteParams struct {
	MachineID string `path:"machine_id" api:"required" json:"-"`
	IfMatch   string `header:"If-Match" api:"required" json:"-"`
	paramObj
}

type MachineWatchParams struct {
	MachineID   string            `path:"machine_id" api:"required" json:"-"`
	LastEventID param.Opt[string] `header:"Last-Event-ID,omitzero" json:"-"`
	paramObj
}
