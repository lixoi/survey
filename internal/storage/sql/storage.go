package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	config "github.com/lixoi/survey/internal/config"
	log "github.com/lixoi/survey/internal/logger"
	storage "github.com/lixoi/survey/internal/storage"
)

const (
	MAX_QUESTIONS     = 20
	BASE_CLASS        = "information_security_questions"
	PROFILE_CLASS_ONE = "linux_questions"
	PROFILE_CLASS_TWO = "network_questions"
)

type Storage struct { // TODO
	connectParams string
	db            *sqlx.DB
	ctx           context.Context
}

func getRandList(questionsId []int, size int) []int {
	if len(questionsId) < size {
		return nil
	}
	retQuestionsId := make([]int, size)
	// rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		n := rand.Intn(len(questionsId))
		retQuestionsId[i] = questionsId[n]
		questionsId = append(questionsId[:n], questionsId[n+1:]...)
	}

	return retQuestionsId
}

func New(dbparams config.PSQLConfig, log log.Logger) *Storage {
	params := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbparams.DSN, dbparams.Port, dbparams.User, dbparams.Pass, dbparams.DB)
	return &Storage{connectParams: params}
}

func (s *Storage) Connect(c context.Context) (err error) {
	s.db, err = sqlx.Open("pgx", s.connectParams)
	if err != nil {
		return fmt.Errorf("cannot open pgx driver: %w", err)
	}
	s.ctx = c

	return s.db.PingContext(s.ctx)
}

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Storage) AddUser(user storage.User) error {
	row := s.db.QueryRowxContext(s.ctx, `
		SELECT 1 FROM users WHERE id = $1
	`, strconv.FormatInt(user.ID, 10))

	var id int64
	if err := row.Scan(&id); err == nil {
		return fmt.Errorf("Event with ID %d is exist in DB", user.ID)
	}

	if user.BaseQ == "" {
		user.BaseQ = BASE_CLASS
	}
	if user.FirstFrofileQ == "" {
		user.FirstFrofileQ = PROFILE_CLASS_ONE
	}
	if user.SecProfileQ == "" {
		user.SecProfileQ = PROFILE_CLASS_TWO
	}
	if user.ExistTo.IsZero() {
		user.ExistTo = time.Now().AddDate(0, 0, 3)
	}

	query := `
		INSERT INTO events (id, base_questions, first_profile_questions, sec_profile_questions, exist_to)
		VALUES (:id, :title, :created_at, :exist_to, :description, :user_id, :time_send_report)
	`
	_, err := s.db.NamedExecContext(s.ctx, query, map[string]interface{}{
		"id":                      user.ID,
		"base_questions":          user.BaseQ,
		"first_profile_questions": user.FirstFrofileQ,
		"sec_profile_questions":   user.SecProfileQ,
		"exist_to":                user.ExistTo,
	})

	query2 := `
		INSERT INTO events (id, title, created_at, exist_to, description, user_id, time_send_report)
		VALUES (:id, :title, :created_at, :exist_to, :description, :user_id, :time_send_report)
	`
	_, err = s.db.NamedExecContext(s.ctx, query2, map[string]interface{}{
		"id":               e.ID,
		"title":            e.Title,
		"created_at":       e.CreatedAt,
		"exist_to":         e.ExistTo,
		"description":      e.Description,
		"user_id":          e.UserID,
		"time_send_report": e.TimeSendReport,
	})

	return err
}

func (s *Storage) Update(e storage.Event) error {
	row := s.db.QueryRowxContext(s.ctx, `
		SELECT 1 FROM events WHERE id = $1
	`, strconv.FormatInt(e.ID, 10))

	var id int64
	if err := row.Scan(&id); err == sql.ErrNoRows {
		return fmt.Errorf("Event with ID %d is not exist in DB", e.ID)
	}

	query := `
		UPDATE events SET id=:id, title=:title, created_at=:created_at, exist_to=:exist_to, description=:description, user_id=:user_id, time_send_report=:time_send_report
		WHERE id = :id
	`
	_, err := s.db.NamedExecContext(s.ctx, query, map[string]interface{}{
		"id":               e.ID,
		"title":            e.Title,
		"created_at":       e.CreatedAt,
		"exist_to":         e.ExistTo,
		"description":      e.Description,
		"user_id":          e.UserID,
		"time_send_report": e.TimeSendReport,
	})

	return err
}

func (s *Storage) Delete(id int64) error {
	row := s.db.QueryRowxContext(s.ctx, `
		SELECT 1 FROM events WHERE id = $1
	`, strconv.FormatInt(id, 10))

	var getID int64
	if err := row.Scan(&getID); err == sql.ErrNoRows {
		return fmt.Errorf("Event with ID %d is not exist in DB", id)
	}

	_, err := s.db.NamedExecContext(s.ctx, `
		DELETE FROM events WHERE id = :id
	`, map[string]interface{}{"id": id})

	return err
}

func (s *Storage) GetListForDay(date time.Time) []storage.Event {
	currentDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	nextDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, 1)
	var res []storage.Event
	query := `
		SELECT
			*
		FROM events
		WHERE created_at BETWEEN $1 AND $2
	`
	if s.db.SelectContext(s.ctx, &res, query, currentDay, nextDay) != nil {
		return nil
	}

	return res
}

func (s *Storage) GetListForWeek(date time.Time) []storage.Event {
	currentDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	finishDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, 7)
	var res []storage.Event
	query := `
			SELECT 
				* 
			FROM events
			WHERE created_at BETWEEN $1 AND $2
	`
	if err := s.db.SelectContext(s.ctx, &res, query, currentDate, finishDate); err != nil {
		return nil
	}

	return res
}

func (s *Storage) GetListForMonth(date time.Time) []storage.Event {
	currentDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	finishDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
	var res []storage.Event
	query := `
		SELECT
			*
		FROM events
		WHERE created_at BETWEEN $1 AND $2
	`
	if s.db.SelectContext(s.ctx, &res, query, currentDay, finishDate) != nil {
		return nil
	}

	return res
}
