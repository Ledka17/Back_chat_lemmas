package user

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"time"
)

type Repository interface {
	GetUserMessages(userID int) ([]model.Message, error)
	CreateMessage(userFromID, userToID int, text string, time time.Time) error
}
