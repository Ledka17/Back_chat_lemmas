package user

import "github.com/Ledka17/Back_chat_lemmas/model"

type Usecase interface {
	GetUserMessages(userID int) ([]model.Message, error)
	SendMessage(userFromID, userToID int, text string) error
}
