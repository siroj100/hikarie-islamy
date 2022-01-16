package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/siroj100/hikarie-islamy/cmd/internal"
	"github.com/siroj100/hikarie-islamy/internal/config"
)

var (
	buildTime, gitRevision, gitBranch string
	buildID                           string
)

func main() {
	var isConfigTest bool
	flag.BoolVar(&isConfigTest, "test", false, "Enable config test mode")
	flag.Parse()

	buildID = fmt.Sprintf("%s-%s:%s", buildTime, gitBranch, gitRevision)
	fmt.Println(os.Args[0], buildID)
	conf := config.Init()
	dbs := internal.InitGormDb(conf.Database)
	ucase := internal.InitUseCase(conf, dbs)
	if isConfigTest {
		// do config test
		fmt.Println("test mode requested, exiting")
		os.Exit(0)
	}
	r := initRoutes(conf, ucase)
	servAddr := fmt.Sprintf("%s:%d", conf.Server.Ip, conf.Server.Port)
	srv := http.Server{
		Addr:         servAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      r,
	}

	fmt.Println("ready. listening on", servAddr)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
