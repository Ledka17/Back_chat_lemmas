package http

import (
	"fmt"
	delivery "github.com/Ledka17/Back_chat_lemmas/delivery/http"
	"github.com/Ledka17/Back_chat_lemmas/model"
	"github.com/Ledka17/Back_chat_lemmas/support"
	"github.com/labstack/echo"
	"strconv"
)

type SupportHandler struct {
	delivery.Handler
	usecase support.Usecase
}

func NewSupportHandler(e *echo.Echo, usecase support.Usecase) SupportHandler {
	handler := SupportHandler{usecase: usecase}
	e.GET(delivery.ApiV1SupportGetChats, handler.GetChatsHandler)
	e.GET(delivery.ApiV1SupportGetChat, handler.GetChatHandler)
	return handler
}

type chat struct {
	User        model.User     `json:"user"`
	LastMessage *model.Message `json:"last_message"`
}

func (h SupportHandler) GetChatHandler(c echo.Context) error {
	userIdString := c.Param("userId")
	fmt.Printf(userIdString)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return h.Error(c, err)
	}
	messages, err := h.usecase.GetAllMessages(userId)
	if err != nil {
		return h.Error(c, err)
	}
	return h.OkWithBody(c, messages)
}

func (h SupportHandler) GetChatsHandler(c echo.Context) error {
	users, err := h.usecase.GetUsers()
	if err != nil {
		return h.Error(c, err)
	}
	chats := make([]chat, 0, len(users))
	for _, user := range users {
		if user.IsSupport {
			continue
		}
		lastMessage, err := h.usecase.GetLastMessage(user.ID)
		if err != nil {
			return nil
		}
		userChat := chat{
			User: user,
		}
		if lastMessage != nil {
			userChat.LastMessage = lastMessage
		}
		chats = append(chats, userChat)
	}
	return h.OkWithBody(c, map[string]interface{}{
		"chats": chats,
	})
}
