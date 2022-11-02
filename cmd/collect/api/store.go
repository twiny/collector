package api

import (
	"context"

	//
	"github.com/twiny/collector/pkg/collector/v1"
)

// Store
type Store interface {
	StoreDetails(ctx context.Context, d *collector.Details) error
	Close() error
}
