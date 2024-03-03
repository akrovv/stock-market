package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/akrovv/client/internal/domain"
	"github.com/akrovv/client/internal/service"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService    UserService
	sessionService SessionService
}

func NewUserHandler(userService UserService, sessionService SessionService) *userHandler {
	return &userHandler{userService: userService, sessionService: sessionService}
}

func (h *userHandler) Register(c echo.Context) error {
	value := c.Request().Header.Get("Content-Type")

	if value == "" {
		return c.JSON(http.StatusBadRequest, "missed content-type")
	}

	data, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "bad body")
	}

	user := domain.User{}
	err = json.Unmarshal(data, &user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unmarshall returned error")
	}

	saveUserDTO := service.SaveUser{Email: user.Email, Nickname: user.Nickname, Password: user.Password, Country: user.Country}
	profile := domain.Profile{Email: user.Email, Nickname: user.Nickname, Country: user.Country}

	err = h.userService.Save(&saveUserDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "server can't save user")
	}

	id, err := h.sessionService.Create(&profile)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "method create session returned error")
	}

	c.SetCookie(&http.Cookie{Name: "session-id", Value: id.ID, Expires: time.Now().Add(time.Hour * 12), HttpOnly: true})
	return c.JSON(http.StatusCreated, &user)
}

func (h *userHandler) Login(c echo.Context) error {
	value := c.Request().Header.Get("Content-Type")

	if value == "" {
		return c.JSON(http.StatusBadRequest, "missed content-type")
	}

	data, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "bad body")
	}

	user := domain.User{}
	err = json.Unmarshal(data, &user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unmarshall returned error")
	}

	geUserDTO := service.GetUser{Email: user.Email, Password: user.Password}
	profile, err := h.userService.Get(&geUserDTO)

	if err != nil {
		if err.Error() == "empty result" {
			return c.JSON(http.StatusUnauthorized, "user doesn't exist")
		}

		return c.JSON(http.StatusInternalServerError, "server can't save user")
	}

	id, err := h.sessionService.Create(profile)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "method create session returned error")
	}

	c.SetCookie(&http.Cookie{Name: "session-id", Value: id.ID, Expires: time.Now().Add(time.Hour * 12), HttpOnly: true})
	return c.JSON(http.StatusOK, &user)
}

func (h *userHandler) Profile(c echo.Context) error {
	contextProfile := c.Get("profile")

	if contextProfile == nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	profile, ok := contextProfile.(*domain.Profile)

	if !ok {
		return c.JSON(http.StatusInternalServerError, "bad context with profile")
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *userHandler) Logout(c echo.Context) error {
	contextProfile := c.Get("profile")

	if contextProfile == nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	c.SetCookie(&http.Cookie{Name: "session-id", Value: "", Expires: time.Now().Add(-1), HttpOnly: true})
	return c.JSON(http.StatusOK, "exited")
}
