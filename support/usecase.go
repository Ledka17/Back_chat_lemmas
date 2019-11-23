package support

import "github.com/Ledka17/Back_chat_lemmas/model"

type Usecase interface {
	GetUsers() ([]model.User, error)
	GetLastMessage(userID int) (*model.Message, error)
}
