package http

import (
	delivery "github.com/Ledka17/Back_chat_lemmas/delivery/http"
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/support"
	"github.com/labstack/echo"
	"net/http"
)

type SupportHandler struct {
	usecase support.Usecase
}

func NewSupportHandler(e *echo.Echo, usecase support.Usecase) SupportHandler {
	handler := SupportHandler{usecase: usecase}
	e.GET(delivery.ApiV1SupportGetChats, handler.GetChatsHandler)
	return handler
}

type chat struct {
	User        model.User    `json:"user"`
	LastMessage model.Message `json:"last_message"`
}

func (h SupportHandler) GetChatsHandler(c echo.Context) error {
	users, err := h.usecase.GetUsers()
	if err != nil {
		return err
	}
	chats := make([]chat, 0, len(users))
	for _, user := range users {
		lastMessage, err := h.usecase.GetLastMessage(user.ID)
		if err != nil {
			return nil
		}
		chats = append(chats, chat{
			User:        user,
			LastMessage: *lastMessage,
		})
	}
	return c.JSON(http.StatusOK, chats)
}