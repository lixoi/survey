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
	Close(c context.Context) error
	//
	AddUser(ctx context.Context, user storage.User) error
	// getQuestions(table string, size int)
	// getQuestion(id int64, table string)
	// addSurvey(user storage.User, questions []storage.Question)
	FinishSurveyFor(ctx context.Context, userId uint64) error
	DeleteUser(ctx context.Context, userId uint64) error
	// deleteSurveyFor(ctx context.Context, userId int64) error
	StartSurveyFor(ctx context.Context, userId uint64) (*storage.Survey, error)
	SetAnswerFor(ctx context.Context, userId uint64, index uint32, answer string) (*storage.Survey, error)
	GetSurveyFor(ctx context.Context, userId uint64) ([]storage.Survey, error)
	GetInfoFor(ctx context.Context, userId uint64) (*storage.User, error)
	// isSurveyStartedFor(ctx context.Context, userId int64) bool
	// isExistQuestionFor(ctx context.Context, userId int64, index int64) int64
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
