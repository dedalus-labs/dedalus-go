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

// MachineSSHService contains methods and other services that help with interacting
// with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMachineSSHService] method instead.
type MachineSSHService struct {
	Options []option.RequestOption
}

// NewMachineSSHService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMachineSSHService(opts ...option.RequestOption) (r MachineSSHService) {
	r = MachineSSHService{}
	r.Options = opts
	return
}

// Create SSH session
func (r *MachineSSHService) New(ctx context.Context, params MachineSSHNewParams, opts ...option.RequestOption) (res *SSHSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/ssh", url.PathEscape(params.MachineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get SSH session
func (r *MachineSSHService) Get(ctx context.Context, query MachineSSHGetParams, opts ...option.RequestOption) (res *SSHSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/ssh/%s", url.PathEscape(query.MachineID), url.PathEscape(query.SessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List SSH sessions
func (r *MachineSSHService) List(ctx context.Context, params MachineSSHListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SSHSession], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/ssh", url.PathEscape(params.MachineID))
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

// List SSH sessions
func (r *MachineSSHService) ListAutoPaging(ctx context.Context, params MachineSSHListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SSHSession] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Delete SSH session
func (r *MachineSSHService) Delete(ctx context.Context, body MachineSSHDeleteParams, opts ...option.RequestOption) (res *SSHSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if body.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/ssh/%s", url.PathEscape(body.MachineID), url.PathEscape(body.SessionID))
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
	MachineID string    `json:"machine_id" api:"required"`
	SessionID string    `json:"session_id" api:"required"`
	// Any of "wake_in_progress", "ready", "closed", "expired", "failed".
	Status       SSHSessionStatus `json:"status" api:"required"`
	Connection   SSHConnection    `json:"connection"`
	ErrorCode    string           `json:"error_code"`
	ErrorMessage string           `json:"error_message"`
	ExpiresAt    time.Time        `json:"expires_at" format:"date-time"`
	ReadyAt      time.Time        `json:"ready_at" format:"date-time"`
	RetryAfterMs int64            `json:"retry_after_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		MachineID    respjson.Field
		SessionID    respjson.Field
		Status       respjson.Field
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

type MachineSSHNewParams struct {
	MachineID              string `path:"machine_id" api:"required" json:"-"`
	SSHSessionCreateParams SSHSessionCreateParams
	paramObj
}

func (r MachineSSHNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SSHSessionCreateParams)
}
func (r *MachineSSHNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineSSHGetParams struct {
	MachineID string `path:"machine_id" api:"required" json:"-"`
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type MachineSSHListParams struct {
	MachineID string            `path:"machine_id" api:"required" json:"-"`
	Cursor    param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit     param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MachineSSHListParams]'s query parameters as `url.Values`.
func (r MachineSSHListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MachineSSHDeleteParams struct {
	MachineID string `path:"machine_id" api:"required" json:"-"`
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}
