package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lixoi/survey/internal/app"
	config "github.com/lixoi/survey/internal/config"
	"github.com/lixoi/survey/internal/logger"
	internalhttp "github.com/lixoi/survey/internal/server/http"
	storage "github.com/lixoi/survey/internal/storage"
	sqlstorage "github.com/lixoi/survey/internal/storage/sql"
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

	strg := sqlstorage.New(config, *logg)
	survey := app.New(logg, strg)

	databaseTests(strg)

	server := internalhttp.NewServer(logg, *survey)

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

	if err = strg.Connect(ctx); err != nil {
		return
	}

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

func databaseTests(strq *sqlstorage.Storage) error {
	c := context.Background()
	strq.Connect(c)
	user := storage.User{
		ID: rand.Int63n(300),
	}
	_ = user
	strq.AddUser(c, user)
	//strq.UpdateSurvey(255, 2, "answer")
	//strq.DeleteUser(284)
	//strq.AddUser(user)
	//strq.UpdateUser(284, true)
	strq.GetSurveyForUser(c, 161)

	return nil
}
