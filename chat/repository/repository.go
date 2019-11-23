package repository

import (
	"github.com/Ledka17/Back_chat_lemmas/chat"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	messageTable = "message"
)

type databaseRepository struct {
	db *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) chat.Repository {
	return &databaseRepository{db: db}
}

func (r *databaseRepository) CreateMessage(userFromID, userToID int, text string, time time.Time) error {
	_, err := r.db.Exec(
		`insert into "`+messageTable+`" (user_from_id, user_to_id, text, sent_date) values ($1, $2, $3, $4)`,
		userFromID, userToID, text, time.Unix(),
	)
	return err
}
