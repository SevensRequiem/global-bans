package source

import (
	"fmt"
	"globalbans/backend/logs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/api/source/ping", func(c echo.Context) error {
		return Ping(c)
	})
}

func Ping(c echo.Context) error {
	ip := c.Param("ip")
	port := c.Param("port")
	logs.LogHTTP(fmt.Sprintf("SourceGame Ping %s:%s", ip, port), 0, "integrations/source.go")
	return c.String(http.StatusOK, "Pong")
}
