package support

import "github.com/Ledka17/Back_chat_lemmas/model"

type Repository interface {
	GetAllUsers() ([]model.User, error)
	GetLastMessage(userID int) (*model.Message, error)
	GetAllMessages(userID int) ([]model.Message, error)
}
