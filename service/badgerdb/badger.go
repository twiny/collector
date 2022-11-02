package badgerdb

import (
	"context"

	//
	"collector/pkg/collector/v1"
)

// BadgerDB
type BadgerDB struct {
	//
}

// NewBadgerDB
func NewBadgerDB(dir string) (*BadgerDB, error) {
	return nil, nil
}

// StoreDetails
func (b *BadgerDB) StoreDetails(ctx context.Context, d *collector.Details) error {
	//
	return nil
}

// Close
func (b *BadgerDB) Close() error {
	//
	return nil
}
