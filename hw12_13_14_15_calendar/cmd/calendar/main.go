package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/app"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/config"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/server/http"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/initstorage"

	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := config.NewConfig(configFile)

	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	logger, err := logger.New(config.Logger.Level, config.Logger.File)
	if err != nil {
		log.Fatalf("Logger error: %v", err)
	}

	ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Database.Username,
		config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name,
		config.Database.SSLMode)

	storage, err := initstorage.NewStorage(ctx, config.Storage, dsn)
	if err != nil {
		logger.Error("failed to connect DB: " + err.Error())
	}

	logger.Info("DB connected...")

	calendar := app.New(logger, storage)

	server := internalhttp.NewServer(logger, calendar)

	//grpc := internalgrpc.New(logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err = server.Stop(ctx); err != nil {
			logger.Error("failed to stop http server: " + err.Error())
		}

		/*if err = internalgrpc.Stop(ctx); err != nil {
			logg.Error("failed to stop grpc server: " + err.Error())
		}

		if err = calendar.Close(ctx); err != nil {
			logg.Error("failed close storage: " + err.Error())
		}*/
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		addrServer := net.JoinHostPort(config.Server.Host, config.Server.HTTPPort)
		if err = server.Start(ctx, addrServer); err != nil {
			logger.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	}()

	/*go func() {
		defer wg.Done()
		addrServer := net.JoinHostPort(config.Server.Host, config.Server.GrpcPort)
		if err = internalgrpc.Start(ctx, addrServer); err != nil {
			logg.Error("failed to start gRPC server: " + err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	}()*/

	<-ctx.Done()
	wg.Wait()
}
