package user

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
)

type Repository interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
}
