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

// WorkspaceExecutionService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceExecutionService] method instead.
type WorkspaceExecutionService struct {
	Options []option.RequestOption
}

// NewWorkspaceExecutionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceExecutionService(opts ...option.RequestOption) (r WorkspaceExecutionService) {
	r = WorkspaceExecutionService{}
	r.Options = opts
	return
}

// Create execution
func (r *WorkspaceExecutionService) New(ctx context.Context, workspaceID string, body WorkspaceExecutionNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/executions", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get execution
func (r *WorkspaceExecutionService) Get(ctx context.Context, executionID string, query WorkspaceExecutionGetParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/executions/%s", url.PathEscape(query.WorkspaceID), url.PathEscape(executionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List executions
func (r *WorkspaceExecutionService) List(ctx context.Context, workspaceID string, query WorkspaceExecutionListParams, opts ...option.RequestOption) (res *pagination.ExecutionList[Execution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/executions", url.PathEscape(workspaceID))
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

// List executions
func (r *WorkspaceExecutionService) ListAutoPaging(ctx context.Context, workspaceID string, query WorkspaceExecutionListParams, opts ...option.RequestOption) *pagination.ExecutionListAutoPager[Execution] {
	return pagination.NewExecutionListAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Delete execution
func (r *WorkspaceExecutionService) Delete(ctx context.Context, executionID string, body WorkspaceExecutionDeleteParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/executions/%s", url.PathEscape(body.WorkspaceID), url.PathEscape(executionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List execution events
func (r *WorkspaceExecutionService) Events(ctx context.Context, executionID string, params WorkspaceExecutionEventsParams, opts ...option.RequestOption) (res *pagination.ExecutionEvents[ExecutionEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/executions/%s/events", url.PathEscape(params.WorkspaceID), url.PathEscape(executionID))
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
func (r *WorkspaceExecutionService) EventsAutoPaging(ctx context.Context, executionID string, params WorkspaceExecutionEventsParams, opts ...option.RequestOption) *pagination.ExecutionEventsAutoPager[ExecutionEvent] {
	return pagination.NewExecutionEventsAutoPager(r.Events(ctx, executionID, params, opts...))
}

// Get execution output
func (r *WorkspaceExecutionService) Output(ctx context.Context, executionID string, query WorkspaceExecutionOutputParams, opts ...option.RequestOption) (res *ExecutionOutput, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required execution_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/executions/%s/output", url.PathEscape(query.WorkspaceID), url.PathEscape(executionID))
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
	// Any of "wake_in_progress", "queued", "running", "succeeded", "failed",
	// "cancelled", "expired".
	Status      ExecutionStatus `json:"status" api:"required"`
	WorkspaceID string          `json:"workspace_id" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema          string        `json:"$schema" format:"uri"`
	Artifacts       []ArtifactRef `json:"artifacts" api:"nullable"`
	CompletedAt     time.Time     `json:"completed_at" format:"date-time"`
	Cwd             string        `json:"cwd"`
	EnvKeys         []string      `json:"env_keys" api:"nullable"`
	ErrorCode       string        `json:"error_code"`
	ErrorMessage    string        `json:"error_message"`
	ExitCode        int64         `json:"exit_code"`
	ExpiresAt       time.Time     `json:"expires_at" format:"date-time"`
	RetryAfterMs    int64         `json:"retry_after_ms"`
	Signal          int64         `json:"signal"`
	StartedAt       time.Time     `json:"started_at" format:"date-time"`
	StderrBytes     int64         `json:"stderr_bytes"`
	StderrTruncated bool          `json:"stderr_truncated"`
	StdoutBytes     int64         `json:"stdout_bytes"`
	StdoutTruncated bool          `json:"stdout_truncated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command         respjson.Field
		CreatedAt       respjson.Field
		ExecutionID     respjson.Field
		Status          respjson.Field
		WorkspaceID     respjson.Field
		Schema          respjson.Field
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
	Command      []string          `json:"command,omitzero" api:"required"`
	Cwd          param.Opt[string] `json:"cwd,omitzero"`
	Stdin        param.Opt[string] `json:"stdin,omitzero"`
	TimeoutMs    param.Opt[int64]  `json:"timeout_ms,omitzero"`
	WakeIfNeeded param.Opt[bool]   `json:"wake_if_needed,omitzero"`
	Env          map[string]string `json:"env,omitzero"`
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
	Items []ExecutionEvent `json:"items" api:"required"`
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
func (r ExecutionEvents) RawJSON() string { return r.JSON.raw }
func (r *ExecutionEvents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionList struct {
	Items []Execution `json:"items" api:"required"`
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
func (r ExecutionList) RawJSON() string { return r.JSON.raw }
func (r *ExecutionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionOutput struct {
	ExecutionID string `json:"execution_id" api:"required"`
	// A URL to the JSON Schema for this object.
	Schema          string `json:"$schema" format:"uri"`
	Stderr          string `json:"stderr"`
	StderrBytes     int64  `json:"stderr_bytes"`
	StderrTruncated bool   `json:"stderr_truncated"`
	Stdout          string `json:"stdout"`
	StdoutBytes     int64  `json:"stdout_bytes"`
	StdoutTruncated bool   `json:"stdout_truncated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExecutionID     respjson.Field
		Schema          respjson.Field
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

type WorkspaceExecutionNewParams struct {
	ExecutionCreateParams ExecutionCreateParams
	paramObj
}

func (r WorkspaceExecutionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExecutionCreateParams)
}
func (r *WorkspaceExecutionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExecutionCreateParams)
}

type WorkspaceExecutionGetParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}

type WorkspaceExecutionListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceExecutionListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceExecutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceExecutionDeleteParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}

type WorkspaceExecutionEventsParams struct {
	WorkspaceID string            `path:"workspace_id" api:"required" json:"-"`
	Cursor      param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit       param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceExecutionEventsParams]'s query parameters as
// `url.Values`.
func (r WorkspaceExecutionEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceExecutionOutputParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}
