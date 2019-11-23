package http

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
}

func (h Handler) Ok(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func (h Handler) OkWithBody(c echo.Context, body interface{}) error {
	return c.JSON(http.StatusOK, body)
}

func (h Handler) Error(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, err.Error())
}
