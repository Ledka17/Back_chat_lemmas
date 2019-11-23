package ws

import (
	"errors"
	"github.com/Ledka17/Back_chat_lemmas/chat"
	delivery "github.com/Ledka17/Back_chat_lemmas/delivery/http"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	delivery.Handler
	usecase chat.Usecase
}

func NewChatHandler(e *echo.Echo, usecase chat.Usecase) ChatHandler {
	handler := ChatHandler{usecase: usecase}
	e.GET(delivery.ApiV1ChatStream, handler.ChatStreamHandler)
	return handler
}

func (h ChatHandler) ChatStreamHandler(c echo.Context) error {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	defer conn.Close()

	userID, err := h.getUserID(c)
	if err != nil {
		conn.WriteJSON(err)
		return nil
	}
	anotherUserID, _ := strconv.Atoi(c.Param("anotherUserID"))

	messages := h.usecase.ListenMessages(userID)
	stop := make(chan bool)
	go func() {
		for {
			select {
			case message := <-messages:
				err := conn.WriteJSON(message)
				if err != nil {
					h.usecase.StopListenMessages(userID)
					return
				}
			case <-stop:
				return
			}
		}
	}()

	for {
		err := h.processRequest(userID, anotherUserID, conn)
		if err != nil {
			stop <- true
			return nil
		}
	}
}

func (h ChatHandler) processRequest(userID, anotherUserID int, c *websocket.Conn) error {
	_, requestReader, _ := c.NextReader()
	if requestReader == nil {
		return errors.New("no content")
	}
	requestBytes, err := ioutil.ReadAll(requestReader)
	if err != nil {
		return err
	}
	message := string(requestBytes)
	err = h.usecase.SendMessage(userID, anotherUserID, message)
	if err != nil {
		c.WriteJSON(err)
	}
	return nil
}

func (h ChatHandler) getUserID(c echo.Context) (int, error) {
	sessionIDCookie, err := c.Cookie(delivery.SessionIDCookieName)
	if err != nil {
		return 0, errors.New("no session cookie")
	}
	userID, _ := strconv.Atoi(sessionIDCookie.Value)
	return userID, nil
}
