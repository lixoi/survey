package storage

import "time"

type Event struct {
	ID             int64
	Title          string
	CreatedAt      time.Time `db:"created_at"`
	ExistTo        time.Time `db:"exist_to"`
	Description    string
	UserID         string    `db:"user_id"`
	TimeSendReport time.Time `db:"time_send_report"`
	// TODO
}
