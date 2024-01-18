package main

import (
	"context"

	"github.com/mkaiho/go-ws-sample/util"
)

func main() {
	ctx := context.Background()
	logger := util.FromContext(ctx)
	logger.Begin("main.main", "begin")("end")

	logger.Info("hello world")
}
