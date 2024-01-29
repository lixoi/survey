package app

import (
	"context"
	"strconv"

	survey "github.com/lixoi/survey/internal/storage"
)

const (
	ErrExistEvent = "ErrExistEvent"
)

type App struct { // TODO
	logger  Logger
	storage Storage
}

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Wirning(msg string)
	Error(msg string)
}

type Storage interface { // TODO
	AddUser(e survey.User) error
	// AddSurvey(e survey.Event) error
	Close() error
	UpdateUser(e survey.User) error
	UpdateSurvey(e survey.Survey) error
	DeleteUser(id int64) error
	// DeleteSurvey(id int64) error
	GetSurveyForUser(id int64) []survey.Survey
}

func New(logger Logger, storage Storage) *App {
	return &App{logger: logger, storage: storage}
}

func (a *App) Create(ctx context.Context, id, title string) error {

	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	err = a.storage.Create(event.Event{ID: intID, Title: title})
	if err != nil {
		a.logger.Error(ErrExistEvent + " : " + err.Error())
	}

	return err
}
