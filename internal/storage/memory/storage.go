package memorystorage

import (
	"fmt"
	"sort"
	"sync"
	"time"

	storage "github.com/lixoi/survey/internal/storage"
)

const (
	capacity = 100
)

type Storage struct {
	listEvents []storage.Event
	mu         sync.RWMutex //nolint:unused
}

func New() *Storage {
	return &Storage{
		listEvents: make([]storage.Event, 0, capacity),
	}
}

func (s *Storage) search(id int64) int {
	for i := 0; i < len(s.listEvents); i++ {
		if s.listEvents[i].ID == id {
			return i
		}
	}

	return -1
}

func (s *Storage) Create(e storage.Event) error {
	if s.search(e.ID) != -1 {
		return fmt.Errorf("Event with ID %d is exist in memory", e.ID)
	}

	s.mu.Lock()
	s.listEvents = append(s.listEvents, e)
	sort.Slice(s.listEvents, func(i, j int) bool {
		return s.listEvents[i].CreatedAt.Compare(s.listEvents[j].CreatedAt) == -1
	})
	s.mu.Unlock()

	return nil
}

func (s *Storage) Close() error {
	s.listEvents = nil

	return nil
}

func (s *Storage) Update(e storage.Event) error {
	i := s.search(e.ID)
	if i == -1 {
		return fmt.Errorf("Event with ID %d is not exist in memory", e.ID)
	}

	s.mu.Lock()
	s.listEvents[i] = e
	sort.Slice(s.listEvents, func(i, j int) bool {
		return s.listEvents[i].CreatedAt.Compare(s.listEvents[j].CreatedAt) == -1
	})
	s.mu.Unlock()

	return nil
}

func (s *Storage) Delete(id int64) error {

	i := s.search(id)
	if i == -1 {
		return fmt.Errorf("Event with ID %d is not exist in memory", id)
	}

	s.mu.Lock()
	s.listEvents = append(s.listEvents[:i], s.listEvents[i+1:]...)
	s.mu.Unlock()

	return nil
}

func (s *Storage) GetListForDay(date time.Time) []storage.Event {
	i := sort.Search(len(s.listEvents), func(i int) bool {
		return s.listEvents[i].CreatedAt.Compare(date) >= 0
	})
	if i == len(s.listEvents) {
		return nil
	}
	finishDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, 1)
	j := sort.Search(len(s.listEvents), func(i int) bool {
		return s.listEvents[i].CreatedAt.Compare(finishDate) >= 0
	})

	return s.listEvents[i:j]
}

func (s *Storage) GetListForWeek(date time.Time) []storage.Event {
	i := sort.Search(len(s.listEvents), func(i int) bool {
		return s.listEvents[i].CreatedAt.Year() == date.Year() && s.listEvents[i].CreatedAt.YearDay() == date.YearDay()
	})
	if i == len(s.listEvents) {
		return nil
	}

	finishDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, 7)
	j := sort.Search(len(s.listEvents), func(i int) bool {
		return s.listEvents[i].CreatedAt.Compare(finishDate) >= 0
	})

	return s.listEvents[i:j]
}

func (s *Storage) GetListForMonth(date time.Time) []storage.Event {
	i := sort.Search(len(s.listEvents), func(i int) bool {
		return s.listEvents[i].CreatedAt.Year() == date.Year() && s.listEvents[i].CreatedAt.YearDay() == date.YearDay()
	})
	if i == len(s.listEvents) {
		return nil
	}
	finishDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
	j := sort.Search(len(s.listEvents), func(i int) bool {
		return s.listEvents[i].CreatedAt.Compare(finishDate) >= 0
	})

	return s.listEvents[i:j]
}
