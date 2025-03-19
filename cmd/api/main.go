package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/radityacandra/banking-challenge/internal/core"
	"github.com/radityacandra/banking-challenge/internal/server"
	"github.com/radityacandra/banking-challenge/pkg/database"
	"github.com/radityacandra/banking-challenge/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger, err := logger.LoadLogger()
	if err != nil {
		log.Fatal("failed to load logger")
		return
	}

	config, err := core.LoadConfig(logger)
	if err != nil {
		log.Fatal("failed to load config")
		return
	}

	logger.Info("starting application, press CTRL + C to gracefully shutdown")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	db, err := database.Init(ctx, config.PostgresUri)
	if err != nil {
		logger.Fatal("failed to establish db connection", zap.Error(err))
		return
	}
	logger.Info("db connection established")

	deps := core.NewDependency(logger, db, config)

	server.InitServer(ctx, deps)

	// process blocked here
	code := deps.GracefulShutdown(ctx)
	os.Exit(code)
}
