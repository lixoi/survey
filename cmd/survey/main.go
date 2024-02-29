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

	internalgrpc "github.com/lixoi/survey/internal/server/grpc"
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

	// databaseTests(strg)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	grpcServer := internalgrpc.New(strg, *logg)
	httpServer := internalhttp.NewServer(logg, *survey)

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpServer.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	if err = strg.Create(ctx); err != nil {
		return
	}

	logg.Info("calendar is running...")

	go func() {
		if err := grpcServer.Start(ctx, "50051"); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	}()

	if err := httpServer.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

func databaseTests(strq *sqlstorage.Storage) error {
	c := context.Background()
	strq.Create(c)
	user := storage.User{
		ID: uint64(rand.Int63n(30000)),
	}
	_ = user
	//strq.AddUser(c, user)
	//strq.StartSurveyFor(c, 2369)
	strq.SetAnswerFor(c, 2369, 2, "answer")
	strq.FinishSurveyFor(c, 2369)
	strq.GetInfoFor(c, 2369)
	strq.GetSurveyFor(c, 2369)
	//strq.UpdateSurvey(255, 2, "answer")
	strq.DeleteUser(c, 2369)
	//strq.AddUser(user)
	//strq.UpdateUser(284, true)
	//strq.GetSurveyForUser(c, 161)

	return nil
}
