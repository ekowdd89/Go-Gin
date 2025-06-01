package cmd

import (
	"context"
	"fmt"
	"net"
	"os"

	httpServer "github.com/ekowdd89/go-gin-boilerplate/pkg/httpserver"
	"github.com/ekowdd89/go-gin-boilerplate/pkg/kafka"
	"github.com/ekowdd89/go-gin-boilerplate/pkg/postgres"
	"github.com/ekowdd89/go-gin-boilerplate/utils"
)

type OptsFunc func(*Cmd) error

type Cmd struct {
	HttpAddr    string `env:"HTTP_LISTENER_ADDR,expand" envDefault:":8080" json:"http_addr"`
	PostgresUrl string `env:"DATABASE_URL,required,notEmpty,expand" json:"postgres_url"`

	dotenv   bool
	envPrefx string

	h *httpServer.HttpServer
	p *postgres.Postgres
}

func New(opts ...OptsFunc) (c *Cmd, err error) {
	c = &Cmd{
		dotenv:   true,
		envPrefx: "GO_GIN_",
	}
	for _, opt := range opts {
		opt(c)
	}

	err = utils.LoadEnv(c, utils.LoadEnvOptions{
		EnvPrefix: c.envPrefx,
		Dotenv:    c.dotenv,
	})
	if err != nil {
		return nil, err
	}

	if err = c.initPostgres(); err != nil {
		return nil, err
	}
	if err = c.initServer(); err != nil {
		return nil, err
	}

	return
}

func (c *Cmd) initPostgres() (err error) {
	opts := []postgres.OptsFunc{
		postgres.WithConnectionString(os.Getenv("GO_GIN_DATABASE_URL")),
	}
	c.p, err = postgres.New(opts...)
	if err != nil {
		return fmt.Errorf("failed to init postgres: %w", err)
	}
	return
}

func (c *Cmd) initServer() (err error) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	h, err := httpServer.New(
		httpServer.WithListener(listener),
		httpServer.WithUserRepository(c.p),
		httpServer.WithMemberRepository(c.p),
	)
	if err != nil {
		panic(err)
	}
	err = h.Run(context.Background())
	return
}

func (c *Cmd) initKafka() (err error) {
	err = kafka.New(
		kafka.WithBrokers([]string{"localhost:9092"}),
		kafka.WithDefaultTopic("test"),
	)
	if err != nil {
		return fmt.Errorf("failed to init kafka: %w", err)
	}
	return
}

func (c *Cmd) Run(ctx context.Context) (err error) {
	return c.h.Run(ctx)
}
