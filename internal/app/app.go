package app

import (
	"context"
	"strconv"
	"time"

	event "github.com/lixoi/survey/internal/storage"
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
	Create(e event.Event) error
	Close() error
	Update(e event.Event) error
	Delete(id int64) error
	GetListForDay(date time.Time) []event.Event
	GetListForWeek(date time.Time) []event.Event
	GetListForMonth(date time.Time) []event.Event
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
