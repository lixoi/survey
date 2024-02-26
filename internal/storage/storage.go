package storage

import "time"

type Survey struct {
	ID             uint64
	UserID         uint64 `db:"user_id"`
	Title          string
	Question       string `db:"question"`
	Answer         string
	AnsweredAt     *time.Time `db:"answered_at"`
	QuestionNumber uint32     `db:"question_number"`
}

type User struct {
	ID            uint64
	BaseQ         string    `db:"base_questions"`
	FirstProfileQ string    `db:"first_profile_questions"`
	SecProfileQ   string    `db:"sec_profile_questions"`
	SurveyDone    bool      `db:"survey_done"`
	SurveyStart   time.Time `db:"survey_start"`
	CreatedAt     time.Time `db:"created_at"`
	ExistTo       time.Time `db:"exist_to"`
}

type Question struct {
	ID           uint64
	QuestionText string `db:"question"`
	Description  string
}
