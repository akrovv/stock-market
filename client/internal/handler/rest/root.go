package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type rootHandler struct {
}

func NewRootHandler() *rootHandler {
	return &rootHandler{}
}

func (h *rootHandler) Main(c echo.Context) error {
	return c.String(http.StatusOK, "sent static")
}
