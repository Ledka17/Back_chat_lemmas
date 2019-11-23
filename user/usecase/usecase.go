package usecase

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/user"
	"time"
)

type userUsecase struct {
	repository user.Repository
}

func NewUserUsecase(repository user.Repository) user.Usecase {
	return &userUsecase{repository: repository}
}

func (u *userUsecase) GetUserMessages(userID int) ([]model.Message, error) {
	return u.repository.GetUserMessages(userID)
}

func (u *userUsecase) SendMessage(userFromID, userToID int, text string) error {
	return u.repository.CreateMessage(userFromID, userToID, text, time.Now())
}
