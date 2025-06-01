package postgres

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

type OptsFunc func(*Postgres) error

func WithConnectionString(dbUrl string) OptsFunc {
	return func(p *Postgres) (err error) {
		p.dbUrl = dbUrl
		p.db, err = otelsql.Open("postgres", dbUrl)
		return
	}
}

type Postgres struct {
	dbUrl   string
	db      *sql.DB
	logger  *log.Logger
	closers []func(context.Context) error
}

func New(opts ...OptsFunc) (p *Postgres, err error) {
	p = &Postgres{
		logger: log.Default(),
	}
	for _, opt := range opts {
		opt(p)
	}
	return
}
