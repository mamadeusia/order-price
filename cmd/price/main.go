package main

import (
	"context"
	"log"
	"preh/config"
	"preh/core"
	"preh/server"
	"preh/server/rabbit"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func main() {
	if err := config.InitConfig("config", []string{"/home/mohammad/Documents/preHiring"}); err != nil {
		log.Fatal(err)
	}
	core.InitCore()
	if err := rabbit.InitMq(); err != nil {
		log.Fatal(err)
	}
	rabbit.DeclareQ()

	ctx, cancel = context.WithCancel(context.Background())
	rabbit.RunConsumer(ctx)

	server.StartGrpcServer()
}
