package restmiddleware

import (
	"github.com/akrovv/client/internal/handler/rest"
	"github.com/akrovv/client/internal/service"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(session rest.SessionService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("session-id")

			if cookie == nil {
				return next(c)
			}

			if err != nil {
				return err
			}

			getSessionDTO := service.GetSession{ID: cookie.Value}
			profile, err := session.Get(&getSessionDTO)

			if err != nil {
				return err
			}

			c.Set("profile", profile)
			return next(c)
		}
	}
}
