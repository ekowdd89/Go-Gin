package main

import (
	"context"

	"github.com/ekowdd89/go-gin-boilerplate/pkg/cmd"
)

func main() {
	c, err := cmd.New()
	if err != nil {
		panic(err)
	}

	c.Run(context.Background())
}
