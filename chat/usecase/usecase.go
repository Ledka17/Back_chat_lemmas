package usecase

import (
	"github.com/Ledka17/Back_chat_lemmas/chat"
	"github.com/Ledka17/Back_chat_lemmas/model"
	"time"
)

type chatUsecase struct {
	repository chat.Repository
	listeners  map[int]chan model.Message
}

func NewChatUsecase(repository chat.Repository) chat.Usecase {
	return &chatUsecase{
		repository: repository,
		listeners:  map[int]chan model.Message{},
	}
}

func (u *chatUsecase) SendMessage(userFromID, userToID int, text string) error {
	now := time.Now()
	err := u.repository.CreateMessage(userFromID, userToID, text, now)
	if err != nil {
		return err
	}
	if listener, ok := u.listeners[userToID]; ok {
		listener <- model.Message{
			UserFromID: userFromID,
			UserToID:   userToID,
			Text:       text,
			SentDate:   now,
		}
	}
	return nil
}

func (u *chatUsecase) ListenMessages(userID int) chan model.Message {
	if listener, ok := u.listeners[userID]; ok {
		return listener
	}
	listener := make(chan model.Message, 5)
	u.listeners[userID] = listener
	return listener
}

func (u *chatUsecase) StopListenMessages(userID int) {
	delete(u.listeners, userID)
}
