//go:generate wire
//go:build wireinject
// +build wireinject

package wire

import (
	"net"
	"github.com/ekowdd89/go-gin-boilerplate/pkg/cmd"
	"github.com/ekowdd89/go-gin-boilerplate/pkg/httpserver"
	"github.com/ekowdd89/go-gin-boilerplate/pkg/postgres"
	"github.com/ekowdd89/go-gin-boilerplate/utils"
	"github.com/google/wire"
)

func ProvidePostgres() (*postgres.Postgres, error) {
	return postgres.New()
}

func ProvideListener() (net.Listener, error) {
	return net.Listen("tcp", ":8080")
}
func ProvideHttpServer(
	p *postgres.Postgres,
	l net.Listener,
) (*httpserver.HttpServer, error) {
	return httpserver.New(
		httpserver.WithListener(l),
		httpserver.WithUserRepository(p),
		httpserver.WithMemberRepository(p),
	)
}


func ProvideCmd(
	s *httpserver.HttpServer,
	p *postgres.Postgres,
)( c *cmd.Cmd, err error) {
	c = &cmd.Cmd{
		Dotenv:      true,
		EnvPrefx:    "GO_GIN_",
		H:           s,
		P:           p,
	}
	err = utils.LoadEnv(c, utils.LoadEnvOptions{
			EnvPrefix: c.EnvPrefx,
			Dotenv:    c.Dotenv,
	})
	if err != nil {
		return nil, err
	}
	return
}

func InitializeCmd() (c *cmd.Cmd, err error) {
	wire.Build(
		ProvidePostgres,
		ProvideListener,
		ProvideHttpServer,
		// ProvideDotEnv,
		// ProvideEnvPrefix,
		ProvideCmd,
	)
	return
}