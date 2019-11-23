package repository

import (
	"database/sql"
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/support"
	"github.com/jmoiron/sqlx"
)

const (
	userTable    = "user"
	messageTable = "message"
)

type databaseRepository struct {
	db *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) support.Repository {
	return &databaseRepository{
		db,
	}
}
func (r *databaseRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, `select * from "`+userTable+`" order by id`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *databaseRepository) GetLastMessage(userID int) (*model.Message, error) {
	lastMessage := model.Message{}
	err := r.db.Get(
		&lastMessage,
		`select * from "`+messageTable+`" where user_from_id=$1 or user_to_id=$1 order by sent_date desc`,
		userID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &lastMessage, nil
}

func (r *databaseRepository) GetAllMessages(userID int) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Select(
		&messages,
		`select * from "`+messageTable+`" where user_from_id=$1 or user_to_id=$1 order by sent_date desc`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
