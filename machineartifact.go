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
	"github.com/dedalus-labs/dedalus-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-go/option"
	"github.com/dedalus-labs/dedalus-go/packages/pagination"
	"github.com/dedalus-labs/dedalus-go/packages/param"
	"github.com/dedalus-labs/dedalus-go/packages/respjson"
)

// MachineArtifactService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMachineArtifactService] method instead.
type MachineArtifactService struct {
	Options []option.RequestOption
}

// NewMachineArtifactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMachineArtifactService(opts ...option.RequestOption) (r MachineArtifactService) {
	r = MachineArtifactService{}
	r.Options = opts
	return
}

// Get artifact
func (r *MachineArtifactService) Get(ctx context.Context, query MachineArtifactGetParams, opts ...option.RequestOption) (res *Artifact, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if query.ArtifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/artifacts/%s", url.PathEscape(query.MachineID), url.PathEscape(query.ArtifactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List artifacts
func (r *MachineArtifactService) List(ctx context.Context, params MachineArtifactListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Artifact], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/artifacts", url.PathEscape(params.MachineID))
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

// List artifacts
func (r *MachineArtifactService) ListAutoPaging(ctx context.Context, params MachineArtifactListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Artifact] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Delete artifact
func (r *MachineArtifactService) Delete(ctx context.Context, body MachineArtifactDeleteParams, opts ...option.RequestOption) (res *Artifact, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MachineID == "" {
		err = errors.New("missing required machine_id parameter")
		return nil, err
	}
	if body.ArtifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/machines/%s/artifacts/%s", url.PathEscape(body.MachineID), url.PathEscape(body.ArtifactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Artifact struct {
	ArtifactID  string    `json:"artifact_id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	MachineID   string    `json:"machine_id" api:"required"`
	Name        string    `json:"name" api:"required"`
	SizeBytes   int64     `json:"size_bytes" api:"required"`
	DownloadURL string    `json:"download_url"`
	ExecutionID string    `json:"execution_id"`
	ExpiresAt   time.Time `json:"expires_at" format:"date-time"`
	MimeType    string    `json:"mime_type"`
	Sha256      string    `json:"sha256"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArtifactID  respjson.Field
		CreatedAt   respjson.Field
		MachineID   respjson.Field
		Name        respjson.Field
		SizeBytes   respjson.Field
		DownloadURL respjson.Field
		ExecutionID respjson.Field
		ExpiresAt   respjson.Field
		MimeType    respjson.Field
		Sha256      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Artifact) RawJSON() string { return r.JSON.raw }
func (r *Artifact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArtifactList struct {
	Items      []Artifact `json:"items" api:"required"`
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
func (r ArtifactList) RawJSON() string { return r.JSON.raw }
func (r *ArtifactList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineArtifactGetParams struct {
	MachineID  string `path:"machine_id" api:"required" json:"-"`
	ArtifactID string `path:"artifact_id" api:"required" json:"-"`
	paramObj
}

type MachineArtifactListParams struct {
	MachineID string            `path:"machine_id" api:"required" json:"-"`
	Cursor    param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit     param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MachineArtifactListParams]'s query parameters as
// `url.Values`.
func (r MachineArtifactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MachineArtifactDeleteParams struct {
	MachineID  string `path:"machine_id" api:"required" json:"-"`
	ArtifactID string `path:"artifact_id" api:"required" json:"-"`
	paramObj
}
