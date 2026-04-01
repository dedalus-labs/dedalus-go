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
)

// MachineExecutionService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMachineExecutionService] method instead.
type MachineExecutionService struct {
	Options []option.RequestOption
}

// NewMachineExecutionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMachineExecutionService(opts ...option.RequestOption) (r MachineExecutionService) {
	r = MachineExecutionService{}
	r.Options = opts
	return
}

// Create execution
func (r *MachineExecutionService) New(ctx context.Context, params MachineExecutionNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/executions", url.PathEscape(params.MachineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get execution
func (r *MachineExecutionService) Get(ctx context.Context, query MachineExecutionGetParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if query.ExecutionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/executions/%s", url.PathEscape(query.MachineID), url.PathEscape(query.ExecutionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List executions
func (r *MachineExecutionService) List(ctx context.Context, params MachineExecutionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Execution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/executions", url.PathEscape(params.MachineID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List executions
func (r *MachineExecutionService) ListAutoPaging(ctx context.Context, params MachineExecutionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Execution] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Delete execution
func (r *MachineExecutionService) Delete(ctx context.Context, body MachineExecutionDeleteParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if body.ExecutionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/executions/%s", url.PathEscape(body.MachineID), url.PathEscape(body.ExecutionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List execution events
func (r *MachineExecutionService) Events(ctx context.Context, params MachineExecutionEventsParams, opts ...option.RequestOption) (res *pagination.CursorPage[ExecutionEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if params.ExecutionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/executions/%s/events", url.PathEscape(params.MachineID), url.PathEscape(params.ExecutionID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List execution events
func (r *MachineExecutionService) EventsAutoPaging(ctx context.Context, params MachineExecutionEventsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ExecutionEvent] {
	return pagination.NewCursorPageAutoPager(r.Events(ctx, params, opts...))
}

// Get execution output
func (r *MachineExecutionService) Output(ctx context.Context, query MachineExecutionOutputParams, opts ...option.RequestOption) (res *ExecutionOutput, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if query.ExecutionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/executions/%s/output", url.PathEscape(query.MachineID), url.PathEscape(query.ExecutionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ArtifactRef struct {
	ArtifactID string `json:"artifact_id" api:"required"`
	Name       string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArtifactID  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArtifactRef) RawJSON() string { return r.JSON.raw }
func (r *ArtifactRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Execution struct {
	Command     []string  `json:"command" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	ExecutionID string    `json:"execution_id" api:"required"`
	MachineID   string    `json:"machine_id" api:"required"`
	// Any of "wake_in_progress", "queued", "running", "succeeded", "failed",
	// "cancelled", "expired".
	Status          ExecutionStatus `json:"status" api:"required"`
	Artifacts       []ArtifactRef   `json:"artifacts" api:"nullable"`
	CompletedAt     time.Time       `json:"completed_at" format:"date-time"`
	Cwd             string          `json:"cwd"`
	EnvKeys         []string        `json:"env_keys" api:"nullable"`
	ErrorCode       string          `json:"error_code"`
	ErrorMessage    string          `json:"error_message"`
	ExitCode        int64           `json:"exit_code"`
	ExpiresAt       time.Time       `json:"expires_at" format:"date-time"`
	RetryAfterMs    int64           `json:"retry_after_ms"`
	Signal          int64           `json:"signal"`
	StartedAt       time.Time       `json:"started_at" format:"date-time"`
	StderrBytes     int64           `json:"stderr_bytes"`
	StderrTruncated bool            `json:"stderr_truncated"`
	StdoutBytes     int64           `json:"stdout_bytes"`
	StdoutTruncated bool            `json:"stdout_truncated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command         respjson.Field
		CreatedAt       respjson.Field
		ExecutionID     respjson.Field
		MachineID       respjson.Field
		Status          respjson.Field
		Artifacts       respjson.Field
		CompletedAt     respjson.Field
		Cwd             respjson.Field
		EnvKeys         respjson.Field
		ErrorCode       respjson.Field
		ErrorMessage    respjson.Field
		ExitCode        respjson.Field
		ExpiresAt       respjson.Field
		RetryAfterMs    respjson.Field
		Signal          respjson.Field
		StartedAt       respjson.Field
		StderrBytes     respjson.Field
		StderrTruncated respjson.Field
		StdoutBytes     respjson.Field
		StdoutTruncated respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Execution) RawJSON() string { return r.JSON.raw }
func (r *Execution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionStatus string

const (
	ExecutionStatusWakeInProgress ExecutionStatus = "wake_in_progress"
	ExecutionStatusQueued         ExecutionStatus = "queued"
	ExecutionStatusRunning        ExecutionStatus = "running"
	ExecutionStatusSucceeded      ExecutionStatus = "succeeded"
	ExecutionStatusFailed         ExecutionStatus = "failed"
	ExecutionStatusCancelled      ExecutionStatus = "cancelled"
	ExecutionStatusExpired        ExecutionStatus = "expired"
)

// The property Command is required.
type ExecutionCreateParams struct {
	Command   []string          `json:"command,omitzero" api:"required"`
	Cwd       param.Opt[string] `json:"cwd,omitzero"`
	Stdin     param.Opt[string] `json:"stdin,omitzero"`
	TimeoutMs param.Opt[int64]  `json:"timeout_ms,omitzero"`
	Env       map[string]string `json:"env,omitzero"`
	paramObj
}

func (r ExecutionCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExecutionCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExecutionCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionEvent struct {
	At       time.Time `json:"at" api:"required" format:"date-time"`
	Sequence int64     `json:"sequence" api:"required"`
	// Any of "lifecycle", "stdout", "stderr".
	Type         ExecutionEventType `json:"type" api:"required"`
	Chunk        string             `json:"chunk"`
	ErrorCode    string             `json:"error_code"`
	ErrorMessage string             `json:"error_message"`
	ExitCode     int64              `json:"exit_code"`
	Signal       int64              `json:"signal"`
	// Any of "wake_in_progress", "queued", "running", "succeeded", "failed",
	// "cancelled", "expired".
	Status ExecutionEventStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		At           respjson.Field
		Sequence     respjson.Field
		Type         respjson.Field
		Chunk        respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		ExitCode     respjson.Field
		Signal       respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionEvent) RawJSON() string { return r.JSON.raw }
func (r *ExecutionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionEventType string

const (
	ExecutionEventTypeLifecycle ExecutionEventType = "lifecycle"
	ExecutionEventTypeStdout    ExecutionEventType = "stdout"
	ExecutionEventTypeStderr    ExecutionEventType = "stderr"
)

type ExecutionEventStatus string

const (
	ExecutionEventStatusWakeInProgress ExecutionEventStatus = "wake_in_progress"
	ExecutionEventStatusQueued         ExecutionEventStatus = "queued"
	ExecutionEventStatusRunning        ExecutionEventStatus = "running"
	ExecutionEventStatusSucceeded      ExecutionEventStatus = "succeeded"
	ExecutionEventStatusFailed         ExecutionEventStatus = "failed"
	ExecutionEventStatusCancelled      ExecutionEventStatus = "cancelled"
	ExecutionEventStatusExpired        ExecutionEventStatus = "expired"
)

type ExecutionEvents struct {
	Items      []ExecutionEvent `json:"items" api:"required"`
	NextCursor string           `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionEvents) RawJSON() string { return r.JSON.raw }
func (r *ExecutionEvents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionList struct {
	Items      []Execution `json:"items" api:"required"`
	NextCursor string      `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionList) RawJSON() string { return r.JSON.raw }
func (r *ExecutionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionOutput struct {
	ExecutionID     string `json:"execution_id" api:"required"`
	Stderr          string `json:"stderr"`
	StderrBytes     int64  `json:"stderr_bytes"`
	StderrTruncated bool   `json:"stderr_truncated"`
	Stdout          string `json:"stdout"`
	StdoutBytes     int64  `json:"stdout_bytes"`
	StdoutTruncated bool   `json:"stdout_truncated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExecutionID     respjson.Field
		Stderr          respjson.Field
		StderrBytes     respjson.Field
		StderrTruncated respjson.Field
		Stdout          respjson.Field
		StdoutBytes     respjson.Field
		StdoutTruncated respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionOutput) RawJSON() string { return r.JSON.raw }
func (r *ExecutionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineExecutionNewParams struct {
	MachineID             string `path:"machine_id" api:"required" json:"-"`
	ExecutionCreateParams ExecutionCreateParams
	paramObj
}

func (r MachineExecutionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExecutionCreateParams)
}
func (r *MachineExecutionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineExecutionGetParams struct {
	MachineID   string `path:"machine_id" api:"required" json:"-"`
	ExecutionID string `path:"execution_id" api:"required" json:"-"`
	paramObj
}

type MachineExecutionListParams struct {
	MachineID string            `path:"machine_id" api:"required" json:"-"`
	Cursor    param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit     param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MachineExecutionListParams]'s query parameters as
// `url.Values`.
func (r MachineExecutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MachineExecutionDeleteParams struct {
	MachineID   string `path:"machine_id" api:"required" json:"-"`
	ExecutionID string `path:"execution_id" api:"required" json:"-"`
	paramObj
}

type MachineExecutionEventsParams struct {
	MachineID   string            `path:"machine_id" api:"required" json:"-"`
	ExecutionID string            `path:"execution_id" api:"required" json:"-"`
	Cursor      param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit       param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MachineExecutionEventsParams]'s query parameters as
// `url.Values`.
func (r MachineExecutionEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MachineExecutionOutputParams struct {
	MachineID   string `path:"machine_id" api:"required" json:"-"`
	ExecutionID string `path:"execution_id" api:"required" json:"-"`
	paramObj
}
