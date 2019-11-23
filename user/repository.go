package user

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
)

type Repository interface {
	GetMessagesFromOrToUser(userID int) ([]model.Message, error)
}
