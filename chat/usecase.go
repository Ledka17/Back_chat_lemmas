package chat

import "github.com/Ledka17/Back_chat_lemmas/model"

type Usecase interface {
	SendMessage(userFromID, userToID int, text string) error
	ListenMessages(userID int) chan model.Message
	StopListenMessages(userID int)
}
