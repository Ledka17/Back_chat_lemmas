package repository

import (
	"database/sql"
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/user"
	"github.com/jmoiron/sqlx"
)

const userTable = "user"

type databaseRepository struct {
	db *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) user.Repository {
	return &databaseRepository{
		db,
	}
}

func (r *databaseRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, `select * from "`+userTable+`" order by id`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *databaseRepository) GetByID(id int) (*model.User, error) {
	userByID := model.User{}
	err := r.db.Get(&userByID, `select * from "`+userTable+`" where id=$1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &userByID, nil
}
