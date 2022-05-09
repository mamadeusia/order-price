package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"preh/config"
	"syscall"

	"preh/server/db"
	"preh/server/http"
	"preh/server/rabbit"
)

func main() {
	if err := config.InitConfig("config", []string{"/home/mohammad/Documents/preHiring"}); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	db.InitDb(ctx)
	// db.SetToRedis(ctx, "psdk", "hosseinoo")

	// fmt.Println(db.GetFromRedis(ctx, "psdk"))
	rabbit.InitMq()
	rabbit.DeclareQ()
	signalForExit := make(chan os.Signal, 1)
	signal.Notify(signalForExit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	http.InitEngine()
	stop := <-signalForExit

	fmt.Println(stop)
}
