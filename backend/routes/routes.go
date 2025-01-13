package routes

import (
	//echo
	"globalbans/backend/auth"
	"globalbans/backend/bans"
	"globalbans/backend/home"
	"globalbans/backend/stats"
	"globalbans/backend/user"
	"globalbans/integration/minecraft"
	"globalbans/integration/ping"
	"globalbans/integration/source"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// Home
	e.Static("/", "frontend/static")
	e.Static("/static", "frontend/static")
	e.Static("/assets", "frontend/assets")
	e.GET("/home", func(c echo.Context) error {
		return home.HomeHandler(c)
	})
	e.GET("/admin", func(c echo.Context) error {
		return home.AdminHandler(c)
	})
	// Auth
	e.GET("/login", func(c echo.Context) error {
		return auth.LoginHandler(c)
	})
	e.GET("/logout", func(c echo.Context) error {
		return auth.LogoutHandler(c)
	})
	e.GET("/auth/callback", func(c echo.Context) error {
		return auth.CallbackHandler(c)
	})

	// API

	e.GET("/api/ping", func(c echo.Context) error {
		return ping.Ping(c)
	})

	e.GET("/api/stats/weekly", func(c echo.Context) error {
		statsData := stats.GetWeeklyStats()
		return c.JSON(http.StatusOK, statsData)
	})
	e.POST("/api/stats/weekly", func(c echo.Context) error {
		bansStr := c.QueryParam("bans")
		bans, err := strconv.Atoi(bansStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid bans value")
		}
		day := time.Now()
		stats.PlusWeeklyBan(bans, day)
		return c.JSON(http.StatusOK, "Success")
	})

	minecraft.Routes(e)
	source.Routes(e)
	bans.Routes(e)
	user.Routes(e)
}
