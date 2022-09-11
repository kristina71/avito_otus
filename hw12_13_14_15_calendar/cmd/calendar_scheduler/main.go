package calendar_scheduler

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	storage2 "github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/config"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/logger"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/mq"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/initstorage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/config_schedule.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	schConfig := config.NewConfigScheduler()

	if err := schConfig.BuildConfigScheduler(configFile); err != nil {
		log.Fatalf("Config error: %v", err)
	}

	logg, err := logger.New(schConfig.Logger.Level, schConfig.Logger.Path)
	if err != nil {
		log.Fatalf("Logger error: %v", err)
	}
	logg.Info("scheduler start")
	defer logg.Info("\nscheduler end")

	period, err := time.ParseDuration(schConfig.Schedule.Period)
	if err != nil {
		logg.Error("failed to parse period from config")
	}

	remindFor, err := time.ParseDuration(schConfig.Schedule.RemindFor)
	if err != nil {
		logg.Error("failed to parse remind_for")
	}

	ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", schConfig.Database.Username,
		schConfig.Database.Password, schConfig.Database.Host, schConfig.Database.Port, schConfig.Database.Name,
		schConfig.Database.SSLMode)

	storage, err := initstorage.NewStorage(ctx, schConfig.Storage, dsn)
	if err != nil {
		logg.Error("failed to connect DB: " + err.Error())
	}
	defer storage.Close(ctx)

	producer := mq.NewProducer(schConfig.Schedule.Uri, schConfig.Schedule.Queue, logg)
	err = producer.Connect()
	if err != nil {
		logg.Error("failed to connect RabbitMQ")
	}
	defer producer.Close()
	fmt.Println(remindFor, period)

	sigChan := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context, remindFor time.Duration,
		period time.Duration, db storage2.Storage, producer *mq.Producer, logg *logger.Logger) {
		for {
			start := time.Now()

			notices, err := db.ListForScheduler(ctx, remindFor, period)
			if err != nil {
				logg.Error(fmt.Sprintf("can't get events: %v", err))
			}

			for _, v := range notices {
				b, err := json.Marshal(v)
				if err != nil {
					logg.Error(fmt.Sprintf("can't marshal notice: %v", err))
				}

				err = producer.Publish(ctx, b)
				if err != nil {
					logg.Error(fmt.Sprintf("can't publish: %v", err))
				}
			}

			err = db.DeleteAll(ctx)
			if err != nil {
				logg.Error(fmt.Sprintf("can't clear old events: %v", err))
			}

			timer := time.NewTimer(period - time.Since(start))
			select {
			case <-timer.C:
				continue
			case <-ctx.Done():
				return
			}
		}
	}(ctx, remindFor, period, storage, producer, logg)

	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
