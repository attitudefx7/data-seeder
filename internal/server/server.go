package server

import (
	"context"
	"data-seeder/internal/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var config *Config

func Init(cfg string) error {
	var err error
	config,err = NewConfigFromFile(cfg)

	if err != nil {
		return err
	}

	if err = db.Init(config.Db); err != nil {
		return err
	}

	return err
}

func Run() {
	addr := fmt.Sprintf(":%d",config.Port)

	router := http.NewServeMux()

	registerRouter(router)

	srv := &http.Server{
		Addr: addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v", err)
	}

	log.Println("Server exiting")
}
