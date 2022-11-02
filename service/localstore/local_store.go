package localstore

import "context"

// LocalStore
type LocalStore struct {
	//
}

// NewLocalStore
func NewLocalStore(dir string) (*LocalStore, error) {
	return nil, nil
}

func (ls *LocalStore) SaveHTML(ctx context.Context, fn string, body []byte) error {
	//
	return nil
}
func (ls *LocalStore) Close() error {
	//
	return nil
}
