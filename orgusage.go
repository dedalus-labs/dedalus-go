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
	"github.com/dedalus-labs/dedalus-go/packages/param"
	"github.com/dedalus-labs/dedalus-go/packages/respjson"
)

// OrgUsageService contains methods and other services that help with interacting
// with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgUsageService] method instead.
type OrgUsageService struct {
	Options []option.RequestOption
}

// NewOrgUsageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgUsageService(opts ...option.RequestOption) (r OrgUsageService) {
	r = OrgUsageService{}
	r.Options = opts
	return
}

// Get org billed machine usage
func (r *OrgUsageService) Get(ctx context.Context, params OrgUsageGetParams, opts ...option.RequestOption) (res *OrgUsage, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.OrgID == "" {
		err = errors.New("missing required org_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orgs/%s/usage", url.PathEscape(params.OrgID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List machine storage usage evidence
func (r *OrgUsageService) GetMachineStorageUsage(ctx context.Context, params OrgUsageGetMachineStorageUsageParams, opts ...option.RequestOption) (res *MachineStorageUsage, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.OrgID == "" {
		err = errors.New("missing required org_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orgs/%s/usage/storage/machines", url.PathEscape(params.OrgID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List machine usage evidence
func (r *OrgUsageService) GetMachineUsage(ctx context.Context, params OrgUsageGetMachineUsageParams, opts ...option.RequestOption) (res *MachineUsage, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.OrgID == "" {
		err = errors.New("missing required org_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/orgs/%s/usage/machines", url.PathEscape(params.OrgID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type MachineStorageUsage struct {
	// Exclusive evidence period end.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date-time"`
	// Inclusive evidence period start.
	PeriodStart time.Time `json:"period_start" api:"required" format:"date-time"`
	// Machine-level storage usage evidence rows.
	Rows []MachineStorageUsageEvidence `json:"rows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineStorageUsage) RawJSON() string { return r.JSON.raw }
func (r *MachineStorageUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineStorageUsageEvidence struct {
	// Exclusive evidence bucket end.
	BucketEnd time.Time `json:"bucket_end" api:"required" format:"date-time"`
	// Inclusive evidence bucket start.
	BucketStart time.Time `json:"bucket_start" api:"required" format:"date-time"`
	// Machine logical bytes observed for storage allocation.
	LogicalStorageBytes int64 `json:"logical_storage_bytes" api:"required"`
	// Machine identifier.
	MachineID string `json:"machine_id" api:"required"`
	// Org storage bucket ID this evidence row contributes to.
	OrgMeteringBucketID string `json:"org_metering_bucket_id" api:"required"`
	// Allocated logical MiB-seconds for this machine.
	StorageMiBSeconds int64 `json:"storage_mib_seconds" api:"required"`
	// Stripe storage meter event identifier linked to that org bucket.
	StripeStorageIdentifier string `json:"stripe_storage_identifier" api:"required"`
	// Latest Stripe emission timestamp for the linked org bucket, when emitted.
	LatestStripeEmittedAt time.Time `json:"latest_stripe_emitted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketEnd               respjson.Field
		BucketStart             respjson.Field
		LogicalStorageBytes     respjson.Field
		MachineID               respjson.Field
		OrgMeteringBucketID     respjson.Field
		StorageMiBSeconds       respjson.Field
		StripeStorageIdentifier respjson.Field
		LatestStripeEmittedAt   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineStorageUsageEvidence) RawJSON() string { return r.JSON.raw }
func (r *MachineStorageUsageEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineUsage struct {
	// Evidence granularity used for rows: hour or day.
	Granularity string `json:"granularity" api:"required"`
	// Exclusive evidence period end.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date-time"`
	// Inclusive evidence period start.
	PeriodStart time.Time `json:"period_start" api:"required" format:"date-time"`
	// Machine-level usage evidence rows.
	Rows []MachineUsageEvidence `json:"rows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Granularity respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineUsage) RawJSON() string { return r.JSON.raw }
func (r *MachineUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MachineUsageEvidence struct {
	// Machine-awake seconds in this bucket.
	AwakeSeconds int64 `json:"awake_seconds" api:"required"`
	// Exclusive evidence bucket end.
	BucketEnd time.Time `json:"bucket_end" api:"required" format:"date-time"`
	// Inclusive evidence bucket start.
	BucketStart time.Time `json:"bucket_start" api:"required" format:"date-time"`
	// Requested vCPU millicores multiplied by active CPU seconds.
	CPUMillicoreSeconds int64 `json:"cpu_millicore_seconds" api:"required"`
	// Latest raw window_end represented by this row.
	LastWindowEnd time.Time `json:"last_window_end" api:"required" format:"date-time"`
	// Machine identifier.
	MachineID string `json:"machine_id" api:"required"`
	// Requested memory MiB multiplied by awake seconds.
	MemoryMiBSeconds int64 `json:"memory_mib_seconds" api:"required"`
	// Org compute bucket IDs this evidence row contributes to.
	OrgMeteringBucketIDs []string `json:"org_metering_bucket_ids" api:"required"`
	// Requested memory for this shape, in MiB.
	RequestedMemoryMiB int64 `json:"requested_memory_mib" api:"required"`
	// Requested storage for this shape, in GiB.
	RequestedStorageGiB int64 `json:"requested_storage_gib" api:"required"`
	// Requested vCPU for this shape.
	RequestedVCPU float64 `json:"requested_vcpu" api:"required"`
	// Stable fingerprint for the requested machine shape.
	SpecFingerprint string `json:"spec_fingerprint" api:"required"`
	// Stripe CPU meter event identifiers linked to those org buckets.
	StripeCPUIdentifiers []string `json:"stripe_cpu_identifiers" api:"required"`
	// Stripe memory meter event identifiers linked to those org buckets.
	StripeMemoryIdentifiers []string `json:"stripe_memory_identifiers" api:"required"`
	// Raw usage windows compacted into this evidence row.
	WindowCount int64 `json:"window_count" api:"required"`
	// Latest Stripe emission timestamp for linked org buckets, when emitted.
	LatestStripeEmittedAt time.Time `json:"latest_stripe_emitted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwakeSeconds            respjson.Field
		BucketEnd               respjson.Field
		BucketStart             respjson.Field
		CPUMillicoreSeconds     respjson.Field
		LastWindowEnd           respjson.Field
		MachineID               respjson.Field
		MemoryMiBSeconds        respjson.Field
		OrgMeteringBucketIDs    respjson.Field
		RequestedMemoryMiB      respjson.Field
		RequestedStorageGiB     respjson.Field
		RequestedVCPU           respjson.Field
		SpecFingerprint         respjson.Field
		StripeCPUIdentifiers    respjson.Field
		StripeMemoryIdentifiers respjson.Field
		WindowCount             respjson.Field
		LatestStripeEmittedAt   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineUsageEvidence) RawJSON() string { return r.JSON.raw }
func (r *MachineUsageEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrgUsage struct {
	// Closed awake seconds in billed org buckets for the period.
	BilledAwakeSeconds int64 `json:"billed_awake_seconds" api:"required"`
	// Closed requested vCPU millicores multiplied by active CPU seconds for the
	// period.
	BilledCPUMillicoreSeconds int64 `json:"billed_cpu_millicore_seconds" api:"required"`
	// Closed billable logical MiB-seconds for the period, matching the Stripe storage
	// meter.
	BilledLogicalStorageMiBSeconds int64 `json:"billed_logical_storage_mib_seconds" api:"required"`
	// Closed requested memory MiB multiplied by awake seconds for the period.
	BilledMemoryMiBSeconds int64 `json:"billed_memory_mib_seconds" api:"required"`
	// Plan-included storage in GiB, used as a local guardrail only.
	IncludedStorageGiB int64 `json:"included_storage_gib" api:"required"`
	// Billing plan in effect for the organization.
	PlanSlug string `json:"plan_slug" api:"required"`
	// Current provisioned storage summed across machines in GiB.
	ProvisionedStorageGiB int64 `json:"provisioned_storage_gib" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BilledAwakeSeconds             respjson.Field
		BilledCPUMillicoreSeconds      respjson.Field
		BilledLogicalStorageMiBSeconds respjson.Field
		BilledMemoryMiBSeconds         respjson.Field
		IncludedStorageGiB             respjson.Field
		PlanSlug                       respjson.Field
		ProvisionedStorageGiB          respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrgUsage) RawJSON() string { return r.JSON.raw }
func (r *OrgUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrgUsageGetParams struct {
	OrgID string `path:"org_id" api:"required" json:"-"`
	// Billing period start (YYYY-MM-DD). Defaults to first of current month.
	PeriodStart param.Opt[string] `query:"period_start,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrgUsageGetParams]'s query parameters as `url.Values`.
func (r OrgUsageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUsageGetMachineStorageUsageParams struct {
	OrgID string `path:"org_id" api:"required" json:"-"`
	// Optional machine ID filter.
	MachineID param.Opt[string] `query:"machine_id,omitzero" json:"-"`
	// Last UTC evidence date to include (YYYY-MM-DD). Defaults to current time.
	PeriodEnd param.Opt[string] `query:"period_end,omitzero" json:"-"`
	// Evidence period start (YYYY-MM-DD). Defaults to first of current month.
	PeriodStart param.Opt[string] `query:"period_start,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrgUsageGetMachineStorageUsageParams]'s query parameters as
// `url.Values`.
func (r OrgUsageGetMachineStorageUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUsageGetMachineUsageParams struct {
	OrgID string `path:"org_id" api:"required" json:"-"`
	// Evidence granularity: hour or day. Defaults to hour.
	Granularity param.Opt[string] `query:"granularity,omitzero" json:"-"`
	// Optional machine ID filter.
	MachineID param.Opt[string] `query:"machine_id,omitzero" json:"-"`
	// Last UTC evidence date to include (YYYY-MM-DD). Defaults to current time.
	PeriodEnd param.Opt[string] `query:"period_end,omitzero" json:"-"`
	// Evidence period start (YYYY-MM-DD). Defaults to first of current month.
	PeriodStart param.Opt[string] `query:"period_start,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrgUsageGetMachineUsageParams]'s query parameters as
// `url.Values`.
func (r OrgUsageGetMachineUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
