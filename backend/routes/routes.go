package routes

import (
	//echo
	"globalbans/backend/auth"
	"globalbans/backend/bans"
	"globalbans/backend/home"
	"globalbans/integration/minecraft"
	"globalbans/integration/ping"
	"globalbans/integration/source"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// Home
	e.GET("/", func(c echo.Context) error {
		return home.HomeHandler(c)
	})
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
	e.GET("/api/bans", func(c echo.Context) error {
		return bans.GetAllBans(c)
	})
	e.GET("/api/bans/search", func(c echo.Context) error {
		return bans.SearchBans(c)
	})
	e.GET("/api/ban/:id", func(c echo.Context) error {
		return bans.GetBan(c)
	})
	e.POST("/api/ban", func(c echo.Context) error {
		return bans.CreateBan(c)
	})
	e.DELETE("/api/ban/:id", func(c echo.Context) error {
		return bans.DeleteBan(c)
	})
	e.GET("/api/ping", func(c echo.Context) error {
		return ping.Ping(c)
	})
	e.GET("/api/server/ingest/bans", func(c echo.Context) error {
		return bans.IngestBans(c)
	})

	minecraft.Routes(e)
	source.Routes(e)
}
