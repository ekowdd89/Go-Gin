package main

import (
	"context"

	"github.com/ekowdd89/go-gin-boilerplate/pkg/wire"
)

func main() {
	c, err := wire.InitializeCmd()
	if err != nil {
		panic(err)
	}
	errRun := c.Run(context.Background())
	if errRun != nil {
		panic(errRun)
	}

}
