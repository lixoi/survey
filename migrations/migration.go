package migrations

import (
	"embed"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	goose "github.com/pressly/goose/v3"

	config "github.com/lixoi/survey/internal/config"
	log "github.com/lixoi/survey/internal/logger"
)

//go:embed 13022024_init.sql
var embedMigrations embed.FS

func UpDown(dbparams config.PSQLConfig, migration string, log log.Logger) error {
	dbstring := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbparams.DSN, dbparams.Port, dbparams.User, dbparams.Pass, dbparams.DB)
	db, err := goose.OpenDBWithDriver("pgx", dbstring)
	if err != nil {
		log.Error("goose: failed to open DB: " + err.Error())
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Error("goose: failed to close DB: " + err.Error())
		}
	}()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Error("goose: failed set dialect: " + err.Error())
		return err
	}

	if migration == "Up" {
		err = goose.Up(db, ".")
	}
	if migration == "Down" {
		err = goose.Down(db, ".")
	}
	if err != nil {
		log.Error("goose: failed to migration DB: " + err.Error())
	}

	return err
}
