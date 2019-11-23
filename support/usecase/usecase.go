package usecase

import (
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/support"
)

type SupportUsecase struct {
	repository support.Repository
}

func NewSupportUsecase(repository support.Repository) SupportUsecase {
	return SupportUsecase{repository: repository}
}

func (u SupportUsecase) GetUsers() ([]model.User, error) {
	return u.repository.GetAllUsers()
}

func (u SupportUsecase) GetLastMessage(userID int) (*model.Message, error) {
	return u.repository.GetLastMessage(userID)
}

func (u SupportUsecase) GetAllMessages(userID int) (*model.Message, error) {
	return nil, nil
}
