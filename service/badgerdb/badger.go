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

func (b *BadgerDB) StoreData(ctx context.Context, d *collector.Details) error {
	//
	return nil
}
func (b *BadgerDB) Close() error {
	//
	return nil
}
