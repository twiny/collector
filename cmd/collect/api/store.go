package api

import (
	"context"

	//
	"collector/pkg/collector/v1"
)

// Store
type Store interface {
	StoreData(ctx context.Context, d *collector.Details) error
	Close() error
}
