package api

import "context"

// FileStore
type FileStore interface {
	SaveHTML(ctx context.Context, fn string, body []byte) error
	Close() error
}
