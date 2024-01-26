package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lixoi/survey/internal/app"
	config "github.com/lixoi/survey/internal/config"
	"github.com/lixoi/survey/internal/logger"
	internalhttp "github.com/lixoi/survey/internal/server/http"
	memorystorage "github.com/lixoi/survey/internal/storage/memory"
	migrations "github.com/lixoi/survey/migrations"
)

var (
	migration  string
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.json", "Path to configuration file")
	flag.StringVar(&migration, "migration", "", "Up or Down flag to migration DB")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := config.NewConfig(configFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logg := logger.New(config.Logger.Level)

	if migration != "" && migrations.UpDown(config.PSQL, migration, *logg) != nil {
		return
	}

	storage := memorystorage.New()
	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, *calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
