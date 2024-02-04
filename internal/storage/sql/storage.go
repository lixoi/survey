package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
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

func getRandList(questionsId []int64, size int) []int64 {
	if len(questionsId) < size {
		return nil
	}
	retQuestionsId := make([]int64, size)
	// rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		n := rand.Intn(len(questionsId))
		retQuestionsId[i] = questionsId[n]
		questionsId = append(questionsId[:n], questionsId[n+1:]...)
	}

	return retQuestionsId
}

func New(dbparams config.Config, log log.Logger) *Storage {
	params := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbparams.PSQL.DSN, dbparams.PSQL.Port, dbparams.PSQL.User, dbparams.PSQL.Pass, dbparams.PSQL.DB)
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

func (s *Storage) Create(ctx context.Context) error {
	return nil
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
	type table struct {
		name string
		size int
	}
	for _, t := range []table{
		{name: user.BaseQ, size: MAX_QUESTIONS / 2},
		{name: user.FirstFrofileQ, size: MAX_QUESTIONS / 4},
		{name: user.SecProfileQ, size: MAX_QUESTIONS / 4}} {
		questions := s.getQuestions(t.name, t.size)
		if questions == nil {
			return errors.New("")
		}
		if s.addSurvey(user, questions) != nil {
			return errors.New("")
		}
	}

	if user.ExistTo.IsZero() {
		user.ExistTo = time.Now().AddDate(0, 0, 3)
	}

	query := `
		INSERT INTO users (id, base_questions, first_profile_questions, sec_profile_questions, exist_to)
		VALUES (:id, :base_questions, :first_profile_questions, :sec_profile_questions, :exist_to)
	`
	_, err := s.db.NamedExecContext(s.ctx, query, map[string]interface{}{
		"id":                      user.ID,
		"base_questions":          user.BaseQ,
		"first_profile_questions": user.FirstFrofileQ,
		"sec_profile_questions":   user.SecProfileQ,
		"exist_to":                user.ExistTo,
	})

	return err
}

func (s *Storage) getQuestions(table string, size int) []storage.Question {
	questionsId := []int64{}
	if s.db.SelectContext(s.ctx, &questionsId, `SELECT id FROM $1`, table) != nil {
		return nil
	}
	if len(questionsId) == 0 || len(questionsId) < size {
		return nil
	}

	randQuestions := getRandList(questionsId, size)

	res := []storage.Question{}
	for _, v := range randQuestions {
		res = append(res, s.getQuestion(v, table))
	}

	return res
}

func (s *Storage) getQuestion(id int64, table string) storage.Question {
	res := storage.Question{}
	if s.db.SelectContext(s.ctx,
		&res,
		`SELECT 1 FROM $1 WHERE id = $2`,
		table, strconv.FormatInt(id, 10)) != nil {
		// error log
		return res
	}

	return res
}

func (s *Storage) addSurvey(user storage.User, questions []storage.Question) error {
	query := `
		INSERT INTO survey (user_id, title, question, question_number)
		VALUES (:user_id, :title, :question, :question_number)
	`
	for i, q := range questions {
		if _, err := s.db.NamedExecContext(s.ctx, query, map[string]interface{}{
			"user_id":         user.ID,
			"title":           "",
			"question":        q.QuestionText,
			"question_number": i,
		}); err != nil {
			return err
		}

	}
	return nil
}

func (s *Storage) UpdateUser(id int64, done bool) error {
	row := s.db.QueryRowxContext(s.ctx, `
		SELECT 1 FROM users WHERE id = $1
	`, strconv.FormatInt(id, 10))

	if err := row.Scan(&id); err == sql.ErrNoRows {
		return fmt.Errorf("User  %d is not exist in DB", id)
	}

	query := `
		UPDATE users SET survey_done=:survey_done
		WHERE id=:id
	`
	_, err := s.db.NamedExecContext(s.ctx, query, map[string]interface{}{
		"survey_done": done,
		"id":          id,
	})

	return err
}

func (s *Storage) DeleteUser(id int64) error {
	row := s.db.QueryRowxContext(s.ctx, `
		SELECT 1 FROM events WHERE id = $1
	`, strconv.FormatInt(id, 10))

	var getId int64
	if err := row.Scan(&getId); err == sql.ErrNoRows {
		return fmt.Errorf("User %d is not exist in DB", id)
	}

	_, err := s.db.NamedExecContext(s.ctx, `
		DELETE FROM users WHERE id = :id
	`, map[string]interface{}{"id": id})
	if err != nil {
		return errors.New("")
	}

	return s.deleteSurvey(id)
}

func (s *Storage) deleteSurvey(id int64) error {
	surveyUserId := []int64{}
	if s.db.SelectContext(s.ctx, &surveyUserId, `SELECT id FROM survey WHERE user_id = $1`, id) != nil {
		return nil
	}
	if len(surveyUserId) == 0 {
		return nil
	}

	_, err := s.db.NamedExecContext(s.ctx, `
		DELETE FROM survey WHERE user_id = :user_id
	`, map[string]interface{}{"user_id": id})

	return err
}

func (s *Storage) UpdateSurvey(userId int64, index int64, answer string) error {
	row := s.db.QueryRowxContext(s.ctx, `
		SELECT 1 FROM survey WHERE user_id = $1 AND question_number = $2
	`, strconv.FormatInt(userId, 10), strconv.FormatInt(index, 10))

	var id int64
	if err := row.Scan(&id); err == sql.ErrNoRows {
		return fmt.Errorf("Survey whith index %d for user %d is not exist in DB", index, userId)
	}

	query := `
		UPDATE survey SET answer=:answer, answered_at=:answered_at
		WHERE id = :id
	`
	_, err := s.db.NamedExecContext(s.ctx, query, map[string]interface{}{
		"id":          id,
		"answer":      answer,
		"answered_at": time.Now(),
	})

	return err
}

func (s *Storage) GetSurveyForUser(id int64) []storage.Survey {
	var res []storage.Survey
	query := `SELECT * FROM survey WHERE user_id = $1`
	if s.db.SelectContext(s.ctx, &res, query, id) != nil {
		return nil
	}

	return res
}
