package app

import (
	"context"

	storage "github.com/lixoi/survey/internal/storage"
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
	AddUser(ctx context.Context, user storage.User) error
	// addSurvey(e survey.Event) error
	// getQuestions(table string, size int)
	// getQuestion(id int64, table string)
	// addSurvey(user storage.User, questions []storage.Question)
	FinishSurveyFor(ctx context.Context, userId int64, done bool) error
	DeleteUser(ctx context.Context, userId int64) error
	// deleteSurvey(id int64) error
	StartSurveyFor(ctx context.Context, userId int64) (*storage.Survey, error)
	SetAnswerFor(ctx context.Context, userId int64, index int64, answer string) (*storage.Survey, error)
	GetSurveyFor(ctx context.Context, userId int64) ([]storage.Survey, error)
	// isSurveyStartedFor(ctx context.Context, userId int64) bool
	// isExistQuestionFor(ctx context.Context, userId int64, index int64) int64

	UpdateSurvey(ctx context.Context, userId int64, index int64, answer string) error
	GetSurveyForUser(ctx context.Context, userId int64) []storage.Survey
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
