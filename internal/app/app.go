package app

import (
	"context"

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
	Connect(c context.Context) error
	Create(c context.Context) error
	AddUser(ctx context.Context, e survey.User) error
	// addSurvey(e survey.Event) error
	// getQuestions(table string, size int)
	// getQuestion(id int64, table string)
	// addSurvey(user storage.User, questions []storage.Question)
	UpdateUser(ctx context.Context, id int64, done bool) error
	DeleteUser(ctx context.Context, id int64) error
	// deleteSurvey(id int64) error
	UpdateSurvey(ctx context.Context, userId int64, index int64, answer string) error
	GetSurveyForUser(ctx context.Context, id int64) []survey.Survey
	Close(c context.Context) error
}

func New(logger Logger, storage Storage) *App {
	return &App{logger: logger, storage: storage}
}

func (a *App) Create(ctx context.Context) error {

	err := a.storage.Create(ctx)
	if err != nil {
		a.logger.Error(ErrExistEvent + " : " + err.Error())
	}

	return err
}
