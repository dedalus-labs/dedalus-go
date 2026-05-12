// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalus

import (
	"github.com/dedalus-labs/dedalus-go/option"
)

// OrgService contains methods and other services that help with interacting with
// the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgService] method instead.
type OrgService struct {
	Options []option.RequestOption
	Usage   OrgUsageService
}

// NewOrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgService(opts ...option.RequestOption) (r OrgService) {
	r = OrgService{}
	r.Options = opts
	r.Usage = NewOrgUsageService(opts...)
	return
}
