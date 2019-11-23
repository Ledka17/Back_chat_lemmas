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

func (u *userUsecase) RegisterUser() (userID int, err error) {
	return u.repository.CreateUser("Новый клиент")
}

func (u *userUsecase) UpdateUser(userID int, name, email string) error {
	return u.repository.UpdateUser(userID, name, email)
}
