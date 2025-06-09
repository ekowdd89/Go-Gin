package postgres

import (
	// "context"
	"database/sql"
	"log"
	// "os"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

type OptsFunc func(*Postgres) error
type Postgres struct {
	// dbUrl   string
	db      *sql.DB
	logger  *log.Logger
	// closers []func(context.Context) error
}

func New(opts ...OptsFunc) (p *Postgres, err error) {
	p = &Postgres{
		logger: log.Default(),
	}
	for _ , opt := range opts {
		if err := opt(p); err != nil {
			return nil, err
		}
	}
	p.db, err = otelsql.Open("postgres", "postgres://postgres:cempakamasp23@localhost:5432/sertif_test?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return
}
