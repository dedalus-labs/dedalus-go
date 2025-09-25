// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalus

import (
	"github.com/dedalus-labs/dedalus-go/internal/apierror"
	"github.com/dedalus-labs/dedalus-go/packages/param"
	"github.com/dedalus-labs/dedalus-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Order = shared.Order

// Order Status
//
// This is an alias to an internal type.
type OrderStatus = shared.OrderStatus

// Equals "placed"
const OrderStatusPlaced = shared.OrderStatusPlaced

// Equals "approved"
const OrderStatusApproved = shared.OrderStatusApproved

// Equals "delivered"
const OrderStatusDelivered = shared.OrderStatusDelivered

// This is an alias to an internal type.
type OrderParam = shared.OrderParam
