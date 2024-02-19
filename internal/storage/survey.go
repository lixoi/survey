package storage

import "time"

type Survey struct {
	ID             int64
	UserID         string `db:"user_id"`
	Title          string
	Question       string `db:"question"`
	Answer         string
	AnsweredAt     time.Time `db:"answered_at"`
	QuestionNumber int64     `db:"question_number"`
}

type User struct {
	ID            int64
	BaseQ         string    `db:"base_questions"`
	FirstFrofileQ string    `db:"first_profile_questions"`
	SecProfileQ   string    `db:"sec_profile_questions"`
	SurveyDone    bool      `db:"survey_done"`
	CreatedAt     time.Time `db:"created_at"`
	ExistTo       time.Time `db:"exist_to"`
}

type Question struct {
	ID           int64
	QuestionText string `db:"question"`
	Description  string
}
