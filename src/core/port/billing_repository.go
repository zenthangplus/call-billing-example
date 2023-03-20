package port

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"time"
)

type BillingRepository interface {
	// IncreaseByUsername is an atomic action to
	// increase billing data for a given username.
	// * Safe for concurrency *
	IncreaseByUsername(ctx context.Context, username string, increasedDuration time.Duration, increasedCount int) error

	// FindOneByUsername find a billing record by username
	FindOneByUsername(ctx context.Context, username string) (*entity.Billing, error)
}
