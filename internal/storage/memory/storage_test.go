package memorystorage

import (
	"testing"
	"time"

	storage "github.com/lixoi/survey/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	//  Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
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
		store := New()
		for i := range el {
			require.Equal(t, nil, store.Create(el[i]))
		}

		require.Equal(t, 6, len(store.GetListForDay(date)), "don't compare events for list of day")
		require.Equal(t, 8, len(store.GetListForWeek(date)), "don't compare events for list of week")
		require.Equal(t, 9, len(store.GetListForMonth(date)), "don't compare events for list of month")
	})

	t.Run("opertion update ", func(t *testing.T) {
		store := New()
		for i := range el {
			require.Equal(t, nil, store.Create(el[i]))
		}

		require.NotEqual(t, nil, store.Update(storage.Event{ID: 11, CreatedAt: time.Now().AddDate(0, 0, 6)}))
		require.Equal(t, nil, store.Update(storage.Event{ID: 3, CreatedAt: time.Now().AddDate(0, 0, 6)}))
		require.Equal(t, 5, len(store.GetListForDay(date)), "don't compare events for list of day")
	})

	t.Run("opertion delete ", func(t *testing.T) {
		store := New()
		for i := range el {
			require.Equal(t, nil, store.Create(el[i]))
		}

		require.NotEqual(t, nil, store.Delete(11))
		require.Equal(t, nil, store.Delete(3))
		require.Equal(t, 5, len(store.GetListForDay(date)), "don't compare events for list of day")
		require.Equal(t, nil, store.Delete(6))
		require.Equal(t, 5, len(store.GetListForDay(date)), "don't compare events for list of day")
	})

}
