package pg

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/pr1sonmike/rest-api-example/config"
)

// Timeout is a Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*pg.DB
}

// Dial creates new database connection to postgres
func Dial() (*DB, error) {
	cfg := config.Get()
	if cfg.PgURL == "" {
		return nil, nil
	}
	pgOpts, err := pg.ParseURL(cfg.PgURL)
	if err != nil {
		return nil, err
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	pgDB.WithTimeout(time.Second * time.Duration(Timeout))

	return &DB{pgDB}, nil
}
