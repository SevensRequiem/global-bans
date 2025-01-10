package minecraft

import (
	"fmt"
	"globalbans/backend/logs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/api/minecraft/ping", func(c echo.Context) error {
		return Ping(c)
	})
}

func Ping(c echo.Context) error {
	ip := c.Param("ip")
	port := c.Param("port")
	logs.LogHTTP(fmt.Sprintf("Minecraft Ping %s:%s", ip, port), 0, "integrations/minecraft.go")
	return c.String(http.StatusOK, "Pong")
}
