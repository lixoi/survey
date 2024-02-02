package sqlstorage

import (
	"context"
	"testing"
	"time"

	storage "github.com/lixoi/survey/internal/storage"
	"github.com/stretchr/testify/require"

	config "github.com/lixoi/survey/internal/config"
	"github.com/lixoi/survey/internal/logger"
	migrations "github.com/lixoi/survey/migrations"
)

func TestStorage(t *testing.T) {
	//  Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	configFile := "/home/user/otus/lixoi-hw1/hw12_13_14_15_calendar/config.json"

	config, _ := config.NewConfig(configFile)
	logg := logger.New(config.Logger.Level)
	migrations.UpDown(config.PSQL, "", *logg)
	store := New(config.PSQL, *logg)
	store.Connect(context.Background())

	date := time.Now()
	el := []storage.Event{
		{
			ID:             1,
			Title:          "event1",
			CreatedAt:      time.Date(date.Year(), date.Month(), date.Day(), 12, 12, 30, 0, time.UTC).AddDate(0, 0, 1),
			ExistTo:        time.Date(date.Year(), date.Month(), date.Day(), 13, 0, 0, 1000, time.UTC),
			Description:    "test1",
			UserID:         "1",
			TimeSendReport: date,
		},
		{ID: 2, CreatedAt: time.Now()},
		{ID: 3, CreatedAt: time.Now()},
		{ID: 4, CreatedAt: time.Now().AddDate(0, 0, 10)},
		{ID: 5, CreatedAt: time.Now()},
		{ID: 6, CreatedAt: time.Now().AddDate(0, 1, 0)},
		{ID: 7, CreatedAt: time.Now()},
		{ID: 8, CreatedAt: time.Now().AddDate(0, 0, 4)},
		{ID: 9, CreatedAt: time.Now()},
		{ID: 10, CreatedAt: time.Now()},
	}

	t.Run("list operations", func(t *testing.T) {
		for i := range el {
			require.Equal(t, nil, store.Create(el[i]))
		}

		require.Equal(t, 6, len(store.GetListForDay(date)), "don't compare events for list of day")
		require.Equal(t, 8, len(store.GetListForWeek(date)), "don't compare events for list of week")
		require.Equal(t, 9, len(store.GetListForMonth(date)), "don't compare events for list of month")
	})

	t.Run("delete operation", func(t *testing.T) {
		require.Equal(t, nil, store.Delete(el[3].ID))
		require.Equal(t, 8, len(store.GetListForMonth(date)), "don't compare events for list of month")
	})

	t.Run("update operation", func(t *testing.T) {
		require.Equal(t, nil, store.Update(storage.Event{ID: 5, CreatedAt: time.Now().AddDate(0, 2, 0)}))
		require.Equal(t, 7, len(store.GetListForMonth(date)), "don't compare events for list of month")
	})
}
