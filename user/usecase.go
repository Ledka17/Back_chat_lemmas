package user

import "github.com/Ledka17/Back_chat_lemmas/model"

type Usecase interface {
	GetUserMessages(userID int) ([]model.Message, error)
	RegisterUser() (userID int, err error)
	UpdateUser(userID int, name, email string) error
}
