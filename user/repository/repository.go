package repository

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/user"
	"github.com/jmoiron/sqlx"
)

const (
	messageTable = "message"
	userTable    = "user"
)

type databaseRepository struct {
	db *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) user.Repository {
	return &databaseRepository{
		db,
	}
}

func (r *databaseRepository) GetMessagesFromOrToUser(userID int) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Select(
		&messages,
		`select * from "`+messageTable+`" where user_from_id=$1 or user_to_id=$1 order by sent_date`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *databaseRepository) CreateUser(name string) (userID int, err error) {
	_, err = r.db.Exec(
		`insert into "`+userTable+`" (name) values ($1)`,
		name,
	)
	if err != nil {
		return 0, err
	}
	var u model.User
	err = r.db.Get(
		&u,
		`select * from "`+userTable+`" order by id desc`,
	)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (r *databaseRepository) UpdateUser(userID int, name, email string) error {
	_, err := r.db.Exec(
		`update "`+userTable+`" set name=$1, email=$2 where id=$3`,
		name, email, userID,
	)
	return err
}
