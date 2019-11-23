package repository

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/user"
	"github.com/jmoiron/sqlx"
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
