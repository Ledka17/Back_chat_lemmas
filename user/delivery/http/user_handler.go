package user

import (
	"errors"
	delivery "github.com/Ledka17/Back_chat_lemmas/delivery/http"
	"github.com/Ledka17/Back_chat_lemmas/user"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	delivery.Handler
	usecase user.Usecase
}

func NewUserHandler(e *echo.Echo, usecase user.Usecase) UserHandler {
	handler := UserHandler{usecase: usecase}
	e.POST(delivery.ApiV1UserRegister, handler.RegisterHandler)
	e.POST(delivery.ApiV1UserUpdate, handler.UpdateHandler)
	e.GET(delivery.ApiV1UserGetMessages, handler.GetUserMessagesHandler)
	return handler
}

func (h *UserHandler) RegisterHandler(c echo.Context) error {
	if _, err := h.getUserID(c); err == nil {
		return h.Error(c, errors.New("you already have session cookie"))
	}
	userID, err := h.usecase.RegisterUser()
	if err != nil {
		h.Error(c, err)
	}
	c.SetCookie(&http.Cookie{
		Name:    delivery.SessionIDCookieName,
		Value:   strconv.Itoa(userID),
		Expires: time.Now().Add(delivery.SessionIDCookieExpire),
		Path:    "/",
	})
	return h.Ok(c)
}

type userToUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) UpdateHandler(c echo.Context) error {
	userID, err := h.getUserID(c)
	if err != nil {
		return h.Error(c, err)
	}
	request := &userToUpdate{}
	if err := c.Bind(request); err != nil {
		return h.Error(c, err)
	}
	err = h.usecase.UpdateUser(userID, request.Name, request.Email)
	if err != nil {
		return h.Error(c, err)
	}
	return h.Ok(c)
}

func (h *UserHandler) GetUserMessagesHandler(c echo.Context) error {
	userID, err := h.getUserID(c)
	if err != nil {
		return h.Error(c, err)
	}
	resp, err := h.usecase.GetUserMessages(userID)
	if err != nil {
		return h.Error(c, err)
	}
	return h.OkWithBody(c, resp)
}

func (h *UserHandler) getUserID(c echo.Context) (int, error) {
	cookie, err := c.Cookie(delivery.SessionIDCookieName)
	if err != nil {
		return 0, err
	}
	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return 0, errors.New("invalid session cookie")
	}
	return userID, nil
}
