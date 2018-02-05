package main

import (
	"My-project/db"
	"context"
	"My-project/api"
	"My-project/conf"

	"log"

	"My-project/setup"

	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
	_ "github.com/lib/pq"
)

func main() {
	var (
		cnf conf.Conf
		err error
	)
	if err = env.Parse(&cnf); err != nil {
		log.Fatal("configuration parsed with err", err)
	}

	dbCon := db.StartDB(cnf.DBUser, cnf.DBPassword, cnf.DBName, cnf.DBHost)
	setup.CreateTable(dbCon)

	//Init API
	ctx, cancel := context.WithCancel(context.Background())

	api := api.New(ctx, cnf, dbCon)

	go api.Start()

	//System signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	select {
	case <-c:
		log.Println("Interrupted", api.Stop())
	case <-ctx.Done():
		log.Println("Exited ", ctx.Err())
	}

	// will use for all cases for satisfy vet linter
	cancel()

}
