package repository

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/user"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	messageTable = "message"
)

type databaseRepository struct {
	db *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) user.Repository {
	return &databaseRepository{
		db,
	}
}

func (r *databaseRepository) GetUserMessages(userID int) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Select(
		&messages,
		`select * from "`+messageTable+`" where user_from_id=$1 or user_to_id=$1 order by time`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *databaseRepository) CreateMessage(userFromID, userToID int, text string, time time.Time) error {
	_, err := r.db.Exec(
		`insert into "`+messageTable+`" (user_from_id, user_to_id, text, sent_date) values ($1, $2, $3, $4)`,
		userFromID, userToID, text, time.Unix(),
	)
	return err
}
