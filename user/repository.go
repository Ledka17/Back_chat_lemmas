package user

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
)

type Repository interface {
	GetMessagesFromOrToUser(userID int) ([]model.Message, error)
	CreateUser(name string) (userID int, err error)
	UpdateUser(userID int, name, email string) error
}
