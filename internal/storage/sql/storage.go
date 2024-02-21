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
	MAX_QUESTIONS     = 3
	BASE_CLASS        = "security_questions"
	PROFILE_CLASS_ONE = "linux_questions"
	PROFILE_CLASS_TWO = "network_questions"
)

type Storage struct { // TODO
	connectParams string
	logg          log.Logger
	db            *sqlx.DB
}

func getRandList(questionsId []int64, size int) []int64 {
	if len(questionsId) < size {
		return nil
	}
	retQuestionsId := make([]int64, size)
	for i := 0; i < size; i++ {
		n := rand.Intn(len(questionsId))
		retQuestionsId[i] = questionsId[n]
		questionsId = append(questionsId[:n], questionsId[n+1:]...)
	}

	return retQuestionsId
}

func New(dbparams config.Config, logg log.Logger) *Storage {
	params := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbparams.PSQL.DSN, dbparams.PSQL.Port, dbparams.PSQL.User, dbparams.PSQL.Pass, dbparams.PSQL.DB)
	return &Storage{
		connectParams: params,
		logg:          logg,
	}
}

func (s *Storage) Connect(ctx context.Context) (err error) {
	s.db, err = sqlx.Open("pgx", s.connectParams)
	if err != nil {
		s.logg.Error("cannot open pgx driver: " + err.Error())
		return errors.New("Error conneect to DB")
	}

	if err = s.db.PingContext(ctx); err != nil {
		s.logg.Error("cannot ping to DB")
	}

	return nil
}

func (s *Storage) Create(ctx context.Context) error {
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Storage) AddUser(ctx context.Context, user storage.User) error {
	row := s.db.QueryRowxContext(ctx, `
		SELECT 1 FROM users WHERE id = $1
	`, strconv.FormatInt(user.ID, 10))

	var id int64
	if err := row.Scan(&id); err == nil {
		s.logg.Error("Event with ID " + strconv.FormatInt(user.ID, 10) + " is exist in DB")
		return errors.New("This user is exists in DB")
	}
	if user.BaseQ == "" {
		user.BaseQ = BASE_CLASS
	}
	if user.FirstProfileQ == "" {
		user.FirstProfileQ = PROFILE_CLASS_ONE
	}
	if user.SecProfileQ == "" {
		user.SecProfileQ = PROFILE_CLASS_TWO
	}
	type table struct {
		name string
		size int
	}
	index := 1
	for _, t := range []table{
		{name: user.BaseQ, size: MAX_QUESTIONS / 2},
		{name: user.FirstProfileQ, size: MAX_QUESTIONS / 4},
		{name: user.SecProfileQ, size: MAX_QUESTIONS / 4}} {
		questions := s.getQuestions(ctx, t.name, t.size)
		if questions == nil {
			s.logg.Error("AddUser " + strconv.FormatInt(user.ID, 10) + ": Not questions " + t.name)
			return errors.New("Not questions " + t.name)
		}
		if err := s.addSurvey(ctx, user, questions, index); err != nil {
			s.logg.Error("AddUser " + strconv.FormatInt(user.ID, 10) + ": " + err.Error())
			return errors.New("Not add survey")
		}
		index += len(questions)
	}

	if user.ExistTo.IsZero() {
		user.ExistTo = time.Now().AddDate(0, 0, 3)
	}

	query := `
		INSERT INTO users (id, base_questions, first_profile_questions, sec_profile_questions, exist_to)
		VALUES (:id, :base_questions, :first_profile_questions, :sec_profile_questions, :exist_to)
	`
	_, err := s.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":                      user.ID,
		"base_questions":          user.BaseQ,
		"first_profile_questions": user.FirstProfileQ,
		"sec_profile_questions":   user.SecProfileQ,
		"exist_to":                user.ExistTo,
	})

	if err != nil {
		s.logg.Error("Not add user " + strconv.FormatInt(user.ID, 10) + "to users table: " + err.Error())
		return err
	}

	return nil
}

func (s *Storage) getQuestions(ctx context.Context, table string, size int) []string {
	questionsId := []int64{}
	if err := s.db.SelectContext(ctx, &questionsId, `SELECT id FROM `+table); err != nil {
		s.logg.Error("Do not select from " + table + " : " + err.Error())
		return nil
	}
	if len(questionsId) == 0 || len(questionsId) < size {
		s.logg.Error("Size of questions in " + table + " is not current")
		return nil
	}

	questionsId = getRandList(questionsId, size)

	res := []string{}
	for _, v := range questionsId {
		res = append(res, s.getQuestion(ctx, v, table))
	}

	return res
}

func (s *Storage) getQuestion(ctx context.Context, id int64, table string) string {
	var question string
	row := s.db.QueryRowxContext(ctx, `
		SELECT question FROM `+table+` WHERE id = $1 LIMIT 1
	`, strconv.FormatInt(id, 10))

	if err := row.Scan(&question); err == sql.ErrNoRows {
		s.logg.Error("Not questions in table " + table)
		return ""
	}

	return question
}

func (s *Storage) addSurvey(ctx context.Context, user storage.User, questions []string, index int) error {
	query := `
		INSERT INTO survey (user_id, title, question, question_number)
		VALUES (:user_id, :title, :question, :question_number)
	`
	for i, q := range questions {
		if _, err := s.db.NamedExecContext(ctx, query, map[string]interface{}{
			"user_id":         user.ID,
			"title":           "",
			"question":        q,
			"question_number": index + i,
		}); err != nil {
			s.logg.Error("Not insert question in table survey: " + err.Error())
			return err
		}

	}
	return nil
}

func (s *Storage) UpdateUser(ctx context.Context, id int64, done bool) error {
	row := s.db.QueryRowxContext(ctx, `
		SELECT id FROM users WHERE id = $1 LIMIT 1
	`, strconv.FormatInt(id, 10))

	var getId int64
	if err := row.Scan(&getId); err == sql.ErrNoRows {
		s.logg.Error("Not user " + strconv.FormatInt(id, 10) + " in users table: " + err.Error())
		return fmt.Errorf("Not user in users table")
	}

	query := `UPDATE users SET survey_done = $1 WHERE id= $2`
	_, err := s.db.ExecContext(ctx, query, done, id)
	if err != nil {
		s.logg.Error("Not update user " + strconv.FormatInt(id, 10) + " : " + err.Error())
		return fmt.Errorf("Not update user")
	}

	return nil
}

func (s *Storage) DeleteUser(ctx context.Context, id int64) error {
	row := s.db.QueryRowxContext(ctx, `
		SELECT id FROM users WHERE id = $1 LIMIT 1
	`, strconv.FormatInt(id, 10))

	var getId int64
	if err := row.Scan(&getId); err == sql.ErrNoRows {
		s.logg.Error("Not user " + strconv.FormatInt(id, 10) + " in users table: " + err.Error())
		return fmt.Errorf("User %d is not exist in DB", id)
	}

	if _, err := s.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id); err != nil {
		s.logg.Error("Not delete user " + strconv.FormatInt(id, 10) + " : " + err.Error())
		return fmt.Errorf("Not delete user")
	}

	return s.deleteSurvey(ctx, id)
}

func (s *Storage) deleteSurvey(ctx context.Context, id int64) error {
	surveyUserId := []int64{}
	if err := s.db.SelectContext(ctx, &surveyUserId, `SELECT id FROM survey WHERE user_id = $1`, id); err != nil {
		s.logg.Error("Not select surveys for user " + strconv.FormatInt(id, 10) + " : " + err.Error())
		return err
	}

	if len(surveyUserId) == 0 {
		s.logg.Wirning("Not surveys for user " + strconv.FormatInt(id, 10))
		return nil
	}

	if _, err := s.db.ExecContext(ctx, `DELETE FROM survey WHERE user_id = $1`, id); err != nil {
		s.logg.Error("Not delete surveys for user " + strconv.FormatInt(id, 10) + " : " + err.Error())
		return err
	}

	return nil
}

func (s *Storage) UpdateSurvey(ctx context.Context, userId int64, index int64, answer string) error {
	// if start survey
	if index == 1 && s.isSurveyStartedFor(ctx, userId) {
		return fmt.Errorf("Survey is already started or finished for user %d", userId)
	}

	row := s.db.QueryRowxContext(ctx, `
		SELECT id FROM survey WHERE user_id = $1 AND question_number = $2 LIMIT 1
	`, strconv.FormatInt(userId, 10), strconv.FormatInt(index, 10))
	var id int64
	if err := row.Scan(&id); err == sql.ErrNoRows {
		s.logg.Error("Survey whith index " +
			strconv.FormatInt(index, 10) +
			" for user " +
			strconv.FormatInt(userId, 10) +
			" is not exist in DB: " + err.Error())
		return fmt.Errorf("Survey whith index %d for user %d is not exist in DB", index, userId)
	}

	query := `
		UPDATE survey SET answer=:answer, answered_at=:answered_at
		WHERE id = :id
	`
	if _, err := s.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":          id,
		"answer":      answer,
		"answered_at": time.Now(),
	}); err != nil {
		s.logg.Error("Survey whith index " +
			strconv.FormatInt(index, 10) +
			" or user " +
			strconv.FormatInt(userId, 10) +
			" is not updated in DB: " + err.Error())
		return err
	}

	return nil
}

func (s *Storage) isSurveyStartedFor(ctx context.Context, userId int64) bool {
	row := s.db.QueryRowxContext(ctx, `
		SELECT survey_start FROM users WHERE id = $1 LIMIT 1
		`, strconv.FormatInt(userId, 10))
	var isStart time.Time
	if err := row.Scan(&isStart); err == sql.ErrNoRows || isStart.IsZero() == true {
		s.logg.Error("Survey is already started or finished for user " + strconv.FormatInt(userId, 10))
		return true
	}

	// init start survey for user
	query := `UPDATE users SET survey_start = $1 WHERE id= $2`
	_, err := s.db.ExecContext(ctx, query, time.Now(), userId)
	if err != nil {
		s.logg.Error("Not update user " + strconv.FormatInt(userId, 10) + " : " + err.Error())
		return true
	}

	return false
}

func (s *Storage) GetSurveyForUser(ctx context.Context, id int64) []storage.Survey {
	var res []storage.Survey
	query := `SELECT * FROM survey WHERE user_id = $1`
	if err := s.db.SelectContext(ctx, &res, query, id); err != nil {
		s.logg.Error("Not select survey for user " + strconv.FormatInt(id, 10))
		return nil
	}

	return res
}
