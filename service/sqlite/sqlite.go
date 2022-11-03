package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"os"

	"github.com/twiny/collector/pkg/collector/v1"

	//
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var dbSchema []byte

// SQLiteDB
type SQLiteDB struct {
	db *sql.DB
}

// NewSQLiteDB something like "./tmp/db/"
func NewSQLiteDB(dir string) (*SQLiteDB, error) {
	// create dir if not exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// create file if not exist
	storefile := dir + "/store.db"
	if _, err := os.Stat(storefile); os.IsNotExist(err) {
		f, err := os.Create(storefile)
		if err != nil {
			return nil, err
		}
		f.Close()
	}

	db, err := sql.Open("sqlite3", storefile+"?cache=shared_sync=1&_cache_size=25000")
	if err != nil {
		return nil, err
	}

	// check connectivity
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// init schema
	if _, err := db.ExecContext(context.Background(), string(dbSchema)); err != nil {
		return nil, err
	}

	return &SQLiteDB{
		db: db,
	}, nil
}

// StoreDetails
func (s *SQLiteDB) StoreDetails(ctx context.Context, d *collector.Details) error {
	stmt, err := s.db.PrepareContext(ctx, insertDetail)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		d.ID,
		d.Website,
		d.URL,
		d.PageTitle,
		d.HTMLFile,
		d.FirstVisit,
		d.LastVisit,
	)

	return err
}

// Close
func (s *SQLiteDB) Close() error {
	return s.db.Close()
}

const insertDetail = `
INSERT INTO "details"
	(id, website, url, page_title, html_file, first_visit, last_visit)
VALUES 
	(?,?,?,?,?,?,?);
`
