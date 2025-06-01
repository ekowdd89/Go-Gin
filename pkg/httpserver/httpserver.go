package httpserver

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/ekowdd89/go-gin-boilerplate/pkg/postgres"
	"github.com/gin-gonic/gin"
	// "go.opentelemetry.io/otel"
	// "golang.org/x/net/trace"
)

type OptsFunc func(*HttpServer) error

func WithListener(l net.Listener) OptsFunc {
	return func(hs *HttpServer) (err error) {
		hs.listener = l
		return
	}
}

func WithUserRepository(ur postgres.UserRepository) OptsFunc {
	return func(hs *HttpServer) (err error) {
		hs.userRepo = ur
		return
	}
}

func WithMemberRepository(mb postgres.MemberRepository) OptsFunc {
	return func(hs *HttpServer) (err error) {
		hs.memberRepo = mb
		return
	}
}

type HttpServer struct {
	listener   net.Listener
	handler    *gin.Engine
	server     *http.Server
	userRepo   postgres.UserRepository
	memberRepo postgres.MemberRepository
	// trace    trace.Trace
}

func New(opts ...OptsFunc) (hs *HttpServer, err error) {
	hs = &HttpServer{
		handler: gin.Default(),
		// trace:   otel.Tracer("httpserver"),
	}
	for _, opt := range opts {
		if err := opt(hs); err != nil {
			return nil, err
		}
	}
	err = hs.buildServer()
	return
}

func (h *HttpServer) buildServer() (err error) {
	h.handler.Use(gin.Recovery())
	h.server = &http.Server{
		Handler:  h.handler,
		ErrorLog: log.New(log.Writer(), "", log.Lshortfile),
	}
	h.registerHealtCheck()
	h.registerUsers()
	h.registerMembers()
	return
}

func (h *HttpServer) registerHealtCheck() *HttpServer {
	// h.handler
	// h.handler
	h.handler.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	return h
}
func (h *HttpServer) registerUsers() *HttpServer {
	h.handler.GET("/users", func(c *gin.Context) {
		users, err := h.userRepo.FindUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})
	return h
}

func (h *HttpServer) registerMembers() *HttpServer {
	h.handler.GET("/members", func(c *gin.Context) {
		members, err := h.memberRepo.FindMembers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, members)
	})
	return h
}

func (h *HttpServer) Run(cxt context.Context) (err error) {
	go func() {
		<-cxt.Done()
		err = errors.Join(err, h.server.Shutdown(cxt))
	}()

	if h.listener == nil {
		return h.server.ListenAndServe()
	}
	return errors.Join(err, h.server.Serve(h.listener))
}
func (h *HttpServer) Close(cxt context.Context) (err error) {
	return h.server.Shutdown(cxt)
}
