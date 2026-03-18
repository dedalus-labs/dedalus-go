// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/dedalus-labs/dedalus-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-go/option"
	"github.com/dedalus-labs/dedalus-go/packages/param"
	"github.com/dedalus-labs/dedalus-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type WorkspaceList[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r WorkspaceList[T]) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceList[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *WorkspaceList[T]) GetNextPage() (res *WorkspaceList[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *WorkspaceList[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &WorkspaceList[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type WorkspaceListAutoPager[T any] struct {
	page *WorkspaceList[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewWorkspaceListAutoPager[T any](page *WorkspaceList[T], err error) *WorkspaceListAutoPager[T] {
	return &WorkspaceListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *WorkspaceListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *WorkspaceListAutoPager[T]) Current() T {
	return r.cur
}

func (r *WorkspaceListAutoPager[T]) Err() error {
	return r.err
}

func (r *WorkspaceListAutoPager[T]) Index() int {
	return r.run
}

type SSHSessionList[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r SSHSessionList[T]) RawJSON() string { return r.JSON.raw }
func (r *SSHSessionList[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SSHSessionList[T]) GetNextPage() (res *SSHSessionList[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SSHSessionList[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SSHSessionList[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SSHSessionListAutoPager[T any] struct {
	page *SSHSessionList[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewSSHSessionListAutoPager[T any](page *SSHSessionList[T], err error) *SSHSessionListAutoPager[T] {
	return &SSHSessionListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SSHSessionListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SSHSessionListAutoPager[T]) Current() T {
	return r.cur
}

func (r *SSHSessionListAutoPager[T]) Err() error {
	return r.err
}

func (r *SSHSessionListAutoPager[T]) Index() int {
	return r.run
}

type ExecutionList[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ExecutionList[T]) RawJSON() string { return r.JSON.raw }
func (r *ExecutionList[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ExecutionList[T]) GetNextPage() (res *ExecutionList[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ExecutionList[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ExecutionList[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ExecutionListAutoPager[T any] struct {
	page *ExecutionList[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewExecutionListAutoPager[T any](page *ExecutionList[T], err error) *ExecutionListAutoPager[T] {
	return &ExecutionListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ExecutionListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ExecutionListAutoPager[T]) Current() T {
	return r.cur
}

func (r *ExecutionListAutoPager[T]) Err() error {
	return r.err
}

func (r *ExecutionListAutoPager[T]) Index() int {
	return r.run
}

type ExecutionEvents[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ExecutionEvents[T]) RawJSON() string { return r.JSON.raw }
func (r *ExecutionEvents[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ExecutionEvents[T]) GetNextPage() (res *ExecutionEvents[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ExecutionEvents[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ExecutionEvents[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ExecutionEventsAutoPager[T any] struct {
	page *ExecutionEvents[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewExecutionEventsAutoPager[T any](page *ExecutionEvents[T], err error) *ExecutionEventsAutoPager[T] {
	return &ExecutionEventsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ExecutionEventsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ExecutionEventsAutoPager[T]) Current() T {
	return r.cur
}

func (r *ExecutionEventsAutoPager[T]) Err() error {
	return r.err
}

func (r *ExecutionEventsAutoPager[T]) Index() int {
	return r.run
}

type ArtifactList[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ArtifactList[T]) RawJSON() string { return r.JSON.raw }
func (r *ArtifactList[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ArtifactList[T]) GetNextPage() (res *ArtifactList[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ArtifactList[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ArtifactList[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ArtifactListAutoPager[T any] struct {
	page *ArtifactList[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewArtifactListAutoPager[T any](page *ArtifactList[T], err error) *ArtifactListAutoPager[T] {
	return &ArtifactListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ArtifactListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ArtifactListAutoPager[T]) Current() T {
	return r.cur
}

func (r *ArtifactListAutoPager[T]) Err() error {
	return r.err
}

func (r *ArtifactListAutoPager[T]) Index() int {
	return r.run
}

type PreviewList[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PreviewList[T]) RawJSON() string { return r.JSON.raw }
func (r *PreviewList[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PreviewList[T]) GetNextPage() (res *PreviewList[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PreviewList[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PreviewList[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PreviewListAutoPager[T any] struct {
	page *PreviewList[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPreviewListAutoPager[T any](page *PreviewList[T], err error) *PreviewListAutoPager[T] {
	return &PreviewListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PreviewListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PreviewListAutoPager[T]) Current() T {
	return r.cur
}

func (r *PreviewListAutoPager[T]) Err() error {
	return r.err
}

func (r *PreviewListAutoPager[T]) Index() int {
	return r.run
}

type TerminalList[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TerminalList[T]) RawJSON() string { return r.JSON.raw }
func (r *TerminalList[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TerminalList[T]) GetNextPage() (res *TerminalList[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TerminalList[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TerminalList[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TerminalListAutoPager[T any] struct {
	page *TerminalList[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTerminalListAutoPager[T any](page *TerminalList[T], err error) *TerminalListAutoPager[T] {
	return &TerminalListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TerminalListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TerminalListAutoPager[T]) Current() T {
	return r.cur
}

func (r *TerminalListAutoPager[T]) Err() error {
	return r.err
}

func (r *TerminalListAutoPager[T]) Index() int {
	return r.run
}
