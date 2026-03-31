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

// WorkspaceSSHService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceSSHService] method instead.
type WorkspaceSSHService struct {
	Options []option.RequestOption
}

// NewWorkspaceSSHService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceSSHService(opts ...option.RequestOption) (r WorkspaceSSHService) {
	r = WorkspaceSSHService{}
	r.Options = opts
	return
}

// Create SSH session
func (r *WorkspaceSSHService) New(ctx context.Context, workspaceID string, body WorkspaceSSHNewParams, opts ...option.RequestOption) (res *SSHSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ssh", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get SSH session
func (r *WorkspaceSSHService) Get(ctx context.Context, sessionID string, query WorkspaceSSHGetParams, opts ...option.RequestOption) (res *SSHSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ssh/%s", url.PathEscape(query.WorkspaceID), url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List SSH sessions
func (r *WorkspaceSSHService) List(ctx context.Context, workspaceID string, query WorkspaceSSHListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SSHSession], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ssh", url.PathEscape(workspaceID))
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

// List SSH sessions
func (r *WorkspaceSSHService) ListAutoPaging(ctx context.Context, workspaceID string, query WorkspaceSSHListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SSHSession] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Delete SSH session
func (r *WorkspaceSSHService) Delete(ctx context.Context, sessionID string, body WorkspaceSSHDeleteParams, opts ...option.RequestOption) (res *SSHSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ssh/%s", url.PathEscape(body.WorkspaceID), url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type SSHConnection struct {
	Endpoint        string       `json:"endpoint" api:"required"`
	Port            int64        `json:"port" api:"required"`
	SSHUsername     string       `json:"ssh_username" api:"required"`
	HostTrust       SSHHostTrust `json:"host_trust"`
	UserCertificate string       `json:"user_certificate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoint        respjson.Field
		Port            respjson.Field
		SSHUsername     respjson.Field
		HostTrust       respjson.Field
		UserCertificate respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHConnection) RawJSON() string { return r.JSON.raw }
func (r *SSHConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHHostTrust struct {
	HostPattern string `json:"host_pattern" api:"required"`
	// Any of "cert_authority".
	Kind      SSHHostTrustKind `json:"kind" api:"required"`
	PublicKey string           `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HostPattern respjson.Field
		Kind        respjson.Field
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHHostTrust) RawJSON() string { return r.JSON.raw }
func (r *SSHHostTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHHostTrustKind string

const (
	SSHHostTrustKindCertAuthority SSHHostTrustKind = "cert_authority"
)

type SSHSession struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	SessionID string    `json:"session_id" api:"required"`
	// Any of "wake_in_progress", "ready", "closed", "expired", "failed".
	Status       SSHSessionStatus `json:"status" api:"required"`
	WorkspaceID  string           `json:"workspace_id" api:"required"`
	Connection   SSHConnection    `json:"connection"`
	ErrorCode    string           `json:"error_code"`
	ErrorMessage string           `json:"error_message"`
	ExpiresAt    time.Time        `json:"expires_at" format:"date-time"`
	ReadyAt      time.Time        `json:"ready_at" format:"date-time"`
	RetryAfterMs int64            `json:"retry_after_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		SessionID    respjson.Field
		Status       respjson.Field
		WorkspaceID  respjson.Field
		Connection   respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		ExpiresAt    respjson.Field
		ReadyAt      respjson.Field
		RetryAfterMs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHSession) RawJSON() string { return r.JSON.raw }
func (r *SSHSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHSessionStatus string

const (
	SSHSessionStatusWakeInProgress SSHSessionStatus = "wake_in_progress"
	SSHSessionStatusReady          SSHSessionStatus = "ready"
	SSHSessionStatusClosed         SSHSessionStatus = "closed"
	SSHSessionStatusExpired        SSHSessionStatus = "expired"
	SSHSessionStatusFailed         SSHSessionStatus = "failed"
)

// The property PublicKey is required.
type SSHSessionCreateParams struct {
	PublicKey string `json:"public_key" api:"required"`
	paramObj
}

func (r SSHSessionCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow SSHSessionCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSHSessionCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHSessionList struct {
	Items      []SSHSession `json:"items" api:"required"`
	NextCursor string       `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHSessionList) RawJSON() string { return r.JSON.raw }
func (r *SSHSessionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceSSHNewParams struct {
	SSHSessionCreateParams SSHSessionCreateParams
	paramObj
}

func (r WorkspaceSSHNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SSHSessionCreateParams)
}
func (r *WorkspaceSSHNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceSSHGetParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}

type WorkspaceSSHListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceSSHListParams]'s query parameters as `url.Values`.
func (r WorkspaceSSHListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceSSHDeleteParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}
