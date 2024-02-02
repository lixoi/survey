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
	// addSurvey(e survey.Event) error
	// getQuestions(table string, size int)
	// getQuestion(id int64, table string)
	// addSurvey(user storage.User, questions []storage.Question)
	UpdateUser(id int64, done bool) error
	DeleteUser(id int64) error
	// deleteSurvey(id int64) error
	UpdateSurvey(userId int64, index int64, answer string) error
	GetSurveyForUser(id int64) []survey.Survey
	Close() error
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
