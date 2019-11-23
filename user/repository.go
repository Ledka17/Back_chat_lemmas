package user

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"time"
)

type Repository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)

	GetUserMessages(userID int) ([]model.Message, error)
	CreateMessage(userFromID, userToID int, text string, time time.Time) error
}
