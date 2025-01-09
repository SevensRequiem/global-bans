package ping

import (
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(200, "Pong!")
}
