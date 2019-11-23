package user

import (
	delivery "github.com/Ledka17/Back_chat_lemmas/delivery/http"
	"github.com/Ledka17/Back_chat_lemmas/user"
	"github.com/labstack/echo"
	"strconv"
)

type UserHandler struct {
	delivery.Handler
	usecase user.Usecase
}

func NewUserHandler(usecase user.Usecase, e *echo.Echo) UserHandler {
	handler := UserHandler{usecase: usecase}
	e.GET(delivery.ApiV1UserGetMessages, handler.GetUserMessagesHandler)
	return handler
}

func (h *UserHandler) GetUserMessagesHandler(c echo.Context) error {
	userIdString, err := h.getUser(c)
	if err != nil {
		return h.Error(c, err)
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return h.Error(c, err)
	}
	resp, err := h.usecase.GetUserMessages(userId)
	return h.OkWithBody(c, resp)
}

func (h *UserHandler) getUser(c echo.Context) (string, error) {
	cookie, err := c.Cookie(delivery.SessionIDCookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
