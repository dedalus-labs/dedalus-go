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

// WorkspaceArtifactService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceArtifactService] method instead.
type WorkspaceArtifactService struct {
	Options []option.RequestOption
}

// NewWorkspaceArtifactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceArtifactService(opts ...option.RequestOption) (r WorkspaceArtifactService) {
	r = WorkspaceArtifactService{}
	r.Options = opts
	return
}

// Get artifact
func (r *WorkspaceArtifactService) Get(ctx context.Context, artifactID string, query WorkspaceArtifactGetParams, opts ...option.RequestOption) (res *Artifact, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if artifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/artifacts/%s", url.PathEscape(query.WorkspaceID), url.PathEscape(artifactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List artifacts
func (r *WorkspaceArtifactService) List(ctx context.Context, workspaceID string, query WorkspaceArtifactListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Artifact], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/artifacts", url.PathEscape(workspaceID))
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

// List artifacts
func (r *WorkspaceArtifactService) ListAutoPaging(ctx context.Context, workspaceID string, query WorkspaceArtifactListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Artifact] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Delete artifact
func (r *WorkspaceArtifactService) Delete(ctx context.Context, artifactID string, body WorkspaceArtifactDeleteParams, opts ...option.RequestOption) (res *Artifact, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.WorkspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	if artifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/artifacts/%s", url.PathEscape(body.WorkspaceID), url.PathEscape(artifactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Artifact struct {
	ArtifactID  string    `json:"artifact_id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Name        string    `json:"name" api:"required"`
	SizeBytes   int64     `json:"size_bytes" api:"required"`
	WorkspaceID string    `json:"workspace_id" api:"required"`
	DownloadURL string    `json:"download_url"`
	ExecutionID string    `json:"execution_id"`
	ExpiresAt   time.Time `json:"expires_at" format:"date-time"`
	MimeType    string    `json:"mime_type"`
	Sha256      string    `json:"sha256"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArtifactID  respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		SizeBytes   respjson.Field
		WorkspaceID respjson.Field
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

type WorkspaceArtifactGetParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}

type WorkspaceArtifactListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceArtifactListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceArtifactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceArtifactDeleteParams struct {
	WorkspaceID string `path:"workspace_id" api:"required" json:"-"`
	paramObj
}
