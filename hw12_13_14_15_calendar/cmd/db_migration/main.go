package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	config "github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

var configFile string

var (
	flagSet = flag.NewFlagSet("goose", flag.ExitOnError)
	dir     = flagSet.String("dir", ".", "dir with migration sql files")
)

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
}

func main() {
	flagSet.Parse(os.Args[1:])
	args := flagSet.Args()

	if len(args) < 1 {
		flagSet.Usage()
	}

	config, err := config.NewConfig(configFile)

	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	if config.Storage != "SQL" {
		log.Fatalf("in memory storage is used")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Database.Username,
		config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name,
		config.Database.SSLMode)

	driver := "postgres"

	db, err := goose.OpenDBWithDriver(driver, dsn)
	if err != nil {
		log.Fatalf("failed to open DB with the error: %v", err)
	}

	var arguments []string
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err = goose.Run(args[0], db, *dir, arguments...); err != nil {
		log.Fatalf("goose migrator run: %v", err)
	}
}
