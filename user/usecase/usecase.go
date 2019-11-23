package usecase

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/user"
)

type userUsecase struct {
	repository user.Repository
}

func NewUserUsecase(repository user.Repository) user.Usecase {
	return &userUsecase{repository: repository}
}

func (u *userUsecase) GetUserMessages(userID int) ([]model.Message, error) {
	return u.repository.GetMessagesFromOrToUser(userID)
}
