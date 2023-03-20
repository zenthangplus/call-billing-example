package port

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"time"
)

type BillingRepository interface {
	// Begin a transaction
	Begin(ctx context.Context) (BillingRepository, error)

	// Commit the changes in a transaction
	Commit(ctx context.Context) error

	// Rollback the changes in a transaction
	Rollback(ctx context.Context) error

	// ExistsByUsername checks bill is exists or not by username
	ExistsByUsername(ctx context.Context, username string) (bool, error)

	// Create a bill with given data
	Create(ctx context.Context, username string, callDuration time.Duration, callCount int64) error

	// IncreaseByUsername is an atomic action to
	// increase billing data for a given username.
	// * Safe for concurrency *
	IncreaseByUsername(ctx context.Context, username string, increasedDuration time.Duration, increasedCount int64) error

	// FindOneByUsername find a billing record by username
	FindOneByUsername(ctx context.Context, username string) (*entity.Billing, error)
}
