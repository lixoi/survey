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
	MAX_QUESTIONS     = 4
	BASE_CLASS        = "security_questions"
	PROFILE_CLASS_ONE = "linux_questions"
	PROFILE_CLASS_TWO = "network_questions"
)

type Storage struct { // TODO
	connectParams string
	logg          log.Logger
	db            *sqlx.DB
}

func getRandList(questionsId []uint64, size int) []uint64 {
	if len(questionsId) < size {
		return nil
	}
	retQuestionsId := make([]uint64, size)
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

/*
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
*/

func (s *Storage) Create(ctx context.Context) (err error) {
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

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

// добавление пользователя в БД, автоматическое формирование опросного листа в БД
func (s *Storage) AddUser(ctx context.Context, user storage.User) error {
	// проверка, что пользотавеля с идентификатором нет в БД
	row := s.db.QueryRowxContext(ctx, `
		SELECT id FROM users WHERE id = $1 LIMIT 1
	`, strconv.FormatUint(user.ID, 10))

	var id uint64
	if err := row.Scan(&id); err == nil {
		s.logg.Error("User " + strconv.FormatUint(user.ID, 10) + " already exists in DB")
		return errors.New("This user already exists in DB")
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
		// генерация списка вопросов для каждого класса
		questions := s.getQuestions(ctx, t.name, t.size)
		if questions == nil {
			s.logg.Error("AddUser " + strconv.FormatUint(user.ID, 10) + ": Not questions " + t.name)
			return fmt.Errorf("Not questions in %s", t.name)
		}
		// запись полученного рандомного списка в БД
		if err := s.addSurvey(ctx, user, questions, index); err != nil {
			s.logg.Error("Not add survey for " + strconv.FormatUint(user.ID, 10) + ": " + err.Error())
			return fmt.Errorf("Not add survey")
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
		s.logg.Error("Not add user " + strconv.FormatUint(user.ID, 10) + "to users table: " + err.Error())
		return fmt.Errorf("Not add user")
	}

	return nil
}

func (s *Storage) getQuestions(ctx context.Context, table string, size int) []string {
	questionsId := []uint64{}
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
		q := s.getQuestion(ctx, v, table)
		if q != "" {
			res = append(res, q)
		}
	}

	return res
}

func (s *Storage) getQuestion(ctx context.Context, id uint64, table string) string {
	var question string
	row := s.db.QueryRowxContext(ctx, `
		SELECT question FROM `+table+` WHERE id = $1 LIMIT 1
	`, strconv.FormatUint(id, 10))

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
			return fmt.Errorf("Not insert question in table survey")
		}

	}
	return nil
}

// удаление пользователя из БД
func (s *Storage) DeleteUser(ctx context.Context, userId uint64) error {
	row := s.db.QueryRowxContext(ctx, `
		SELECT id FROM users WHERE id = $1 LIMIT 1
	`, strconv.FormatUint(userId, 10))

	var id uint64
	if err := row.Scan(&id); err == sql.ErrNoRows {
		s.logg.Error("Not user " + strconv.FormatUint(userId, 10) + " in users table: " + err.Error())
		return fmt.Errorf("User is not exist in DB")
	}

	if _, err := s.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, userId); err != nil {
		s.logg.Error("Not delete user " + strconv.FormatUint(userId, 10) + " : " + err.Error())
		return fmt.Errorf("Not delete user")
	}

	// удаление всех вопросов, привязаннывх к пользователю
	return s.deleteSurveyFor(ctx, userId)
}

func (s *Storage) deleteSurveyFor(ctx context.Context, userId uint64) error {
	surveyUserId := []uint64{}
	if err := s.db.SelectContext(ctx, &surveyUserId, `SELECT id FROM survey WHERE user_id = $1`, userId); err != nil {
		s.logg.Error("Not select surveys for user " + strconv.FormatUint(userId, 10) + " : " + err.Error())
		return fmt.Errorf("Not select survey for user")
	}

	if len(surveyUserId) == 0 {
		s.logg.Wirning("Not surveys for user " + strconv.FormatUint(userId, 10) + " in DB")
		return fmt.Errorf("Not surveys for user in DB")
	}

	if _, err := s.db.ExecContext(ctx, `DELETE FROM survey WHERE user_id = $1`, userId); err != nil {
		s.logg.Error("Not delete surveys for user " + strconv.FormatUint(userId, 10) + " : " + err.Error())
		return fmt.Errorf("Not delete surveys for user in DB")
	}

	return nil
}

// завершение опроса, установка флага окончания опроса в параметрах пользователя
func (s *Storage) FinishSurveyFor(ctx context.Context, userId uint64) error {
	row := s.db.QueryRowxContext(ctx, `
		SELECT id FROM users WHERE id = $1 LIMIT 1
	`, strconv.FormatUint(userId, 10))

	var id uint64
	if err := row.Scan(&id); err == sql.ErrNoRows {
		s.logg.Error("Not user " + strconv.FormatUint(userId, 10) + " in users table: " + err.Error())
		return fmt.Errorf("Not user in DB")
	}
	query := `UPDATE users SET survey_done = true WHERE id= $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		s.logg.Error("Not update user " + strconv.FormatUint(id, 10) + " : " + err.Error())
		return fmt.Errorf("Not update user")
	}

	return nil
}

// Старт опроса
// В респонсе возвращается первый вопрос
func (s *Storage) StartSurveyFor(ctx context.Context, userId uint64) (*storage.Survey, error) {
	// если пользователь уже проходил опрос, то возвращается ошибка
	if s.isSurveyStartedFor(ctx, userId) {
		return nil, fmt.Errorf("Survey is already started or finished for user %d", userId)
	}
	// если первого вопроса нет или на вопрос уже был дан ответ, то возвращается ошибка
	id := s.isExistQuestionFor(ctx, userId, 1)
	if id == 0 {
		return nil, fmt.Errorf("Survey is finished")
	}
	// чтение первого вопроса из БД
	res := s.getNextQuestion(ctx, id)
	if res == nil {
		return nil,
			fmt.Errorf("Not questions for user " +
				strconv.FormatUint(userId, 10) +
				", you need to call ICh")
	}

	return res, nil
}

// Запись ответа кандидата в БД
// В респонсе возвращается следующий вопрос
func (s *Storage) SetAnswerFor(ctx context.Context, userId uint64, index uint32, answer string) (*storage.Survey, error) {
	// если вопроса нет или на вопрос уже был дан ответ, то возвращается ошибка
	id := s.isExistQuestionFor(ctx, userId, index)
	if id == 0 {
		return nil,
			fmt.Errorf("Question whith index %d is not exist in DB or you have alraedy answered it", index)
	}
	// запись ответа в БД
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
			strconv.FormatUint(uint64(index), 10) +
			" or user " +
			strconv.FormatUint(userId, 10) +
			" is not updated in DB: " + err.Error())
		return nil, err
	}

	// проверка наличия следующего вопроса
	id = s.isExistQuestionFor(ctx, userId, index+1)
	if id == 0 {
		return nil, fmt.Errorf("Survey is finished")
	}
	// чтение следующего вопроса из БД
	res := s.getNextQuestion(ctx, id)
	if res == nil {
		return nil,
			fmt.Errorf("Not questions for user " +
				strconv.FormatUint(userId, 10) +
				", you need to call ICh")
	}

	return res, nil
}

// получение списка ответов для пользователя
func (s *Storage) GetSurveyFor(ctx context.Context, userId uint64) ([]storage.Survey, error) {
	row := s.db.QueryRowxContext(ctx, `
		SELECT survey_done FROM users WHERE id = $1 LIMIT 1
		`, strconv.FormatUint(userId, 10))
	done := false
	err := row.Scan(&done)
	// проверка наличия пользователя в БД
	if err == sql.ErrNoRows {
		s.logg.Error("There is not user " + strconv.FormatUint(userId, 10) + " in DB")
		return nil, errors.New("There is not user in DB")
	}
	// проверка прохождения опроса пользователем
	if done == false {
		s.logg.Info("User " + strconv.FormatUint(userId, 10) + " have not passed survey yet")
		return nil, errors.New("User have not passed survey yet")
	}
	var res []storage.Survey
	query := `SELECT * FROM survey WHERE user_id = $1`
	if err := s.db.SelectContext(ctx, &res, query, userId); err != nil {
		s.logg.Error("Not select survey for user " + strconv.FormatUint(userId, 10))
		return nil, errors.New("Not select survey for user")
	}

	return res, nil
}

func (s *Storage) GetInfoFor(ctx context.Context, userId uint64) (*storage.User, error) {
	res := []*storage.User{}
	query := `SELECT * FROM users WHERE id = $1 LIMIT 1`
	if err := s.db.SelectContext(ctx, &res, query, userId); err != nil || len(res) != 1 {
		s.logg.Error("Not question with id " + strconv.FormatUint(userId, 10))
		return nil, err
	}

	return res[0], nil
}

func (s *Storage) isSurveyStartedFor(ctx context.Context, userId uint64) bool {
	row := s.db.QueryRowxContext(ctx, `
		SELECT survey_start FROM users WHERE id = $1 LIMIT 1
		`, strconv.FormatUint(userId, 10))
	var isStart interface{}
	if err := row.Scan(&isStart); err == sql.ErrNoRows || isStart != nil {
		s.logg.Error("Survey is already started or finished for user " + strconv.FormatUint(userId, 10))
		return true
	}
	// init start survey for user
	query := `UPDATE users SET survey_start = $1 WHERE id= $2`
	_, err := s.db.ExecContext(ctx, query, time.Now(), userId)
	if err != nil {
		s.logg.Error("Not update user " + strconv.FormatUint(userId, 10) + " : " + err.Error())
		return true
	}

	return false
}

func (s *Storage) isExistQuestionFor(ctx context.Context, userId uint64, index uint32) uint64 {
	row := s.db.QueryRowxContext(ctx, `
		SELECT id, answered_at FROM survey WHERE user_id = $1 AND question_number = $2 LIMIT 1
	`, strconv.FormatUint(userId, 10), strconv.FormatUint(uint64(index), 10))
	var id uint64
	var answeredAt interface{}
	err := row.Scan(&id, &answeredAt)
	if err == sql.ErrNoRows {
		s.logg.Info("Question whith index " +
			strconv.FormatUint(uint64(index), 10) +
			" for user " +
			strconv.FormatUint(userId, 10) +
			" is not exist in DB: " + err.Error())
		return 0
	}
	if answeredAt != nil {
		s.logg.Error("Question whith index " +
			strconv.FormatUint(uint64(index), 10) +
			" for user " +
			strconv.FormatUint(userId, 10) +
			" is already answered")
		return 0
	}

	return id
}

func (s *Storage) getNextQuestion(ctx context.Context, id uint64) *storage.Survey {
	res := []*storage.Survey{}
	query := `SELECT * FROM survey WHERE id = $1 LIMIT 1`
	if err := s.db.SelectContext(ctx, &res, query, id); err != nil || len(res) != 1 {
		s.logg.Error("Not question with id " + strconv.FormatUint(id, 10))
		return nil
	}

	return res[0]
}
