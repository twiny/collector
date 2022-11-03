package localstore

import (
	"context"
	"io/ioutil"
	"os"
)

// LocalStore
type LocalStore struct {
	dir string
}

// NewLocalStore
func NewLocalStore(dir string) (*LocalStore, error) {
	// create dir if not exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	return &LocalStore{
		dir: dir,
	}, nil
}

func (ls *LocalStore) SaveHTML(ctx context.Context, fn string, body []byte) error {
	return ioutil.WriteFile(ls.dir+fn, body, 0)
}

func (ls *LocalStore) Close() error {
	return nil
}
