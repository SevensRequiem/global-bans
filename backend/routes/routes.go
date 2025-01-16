package routes

import (
	//echo
	"globalbans/backend/auth"
	"globalbans/backend/bans"
	"globalbans/backend/home"
	"globalbans/backend/serverauth"
	"globalbans/backend/stats"
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
	e.GET("/", func(c echo.Context) error {
		return home.HomeHandler(c)
	})
	e.GET("/bans", func(c echo.Context) error {
		return home.BansHandler(c)
	})
	e.GET("/servers", func(c echo.Context) error {
		return home.ServersHandler(c)
	})
	e.GET("/appeals", func(c echo.Context) error {
		return home.AppealsHandler(c)
	})
	e.GET("/docs", func(c echo.Context) error {
		return home.DocsHandler(c)
	})

	// Auth
	e.GET("/login", func(c echo.Context) error {
		return home.LoginHandler(c)
	})
	e.POST("/auth/login", func(c echo.Context) error {
		return auth.Login(c)
	})

	e.Static("/assets", "frontend/assets")
	// admin routes
	e.GET("/admin", func(c echo.Context) error {
		if !auth.IsAdmin(c) {
			return home.HomeHandler(c)
		}
		return home.AdminHandler(c)
	})
	e.GET("/admin/bans", func(c echo.Context) error {
		if !auth.IsAdmin(c) {
			return home.HomeHandler(c)
		}
		return home.AdminBansHandler(c)
	})
	e.GET("/admin/dashboard", func(c echo.Context) error {
		if !auth.IsAdmin(c) {
			return home.HomeHandler(c)
		}
		return home.AdminDashboardHandler(c)
	})
	e.GET("/admin/servers", func(c echo.Context) error {
		if !auth.IsAdmin(c) {
			return home.HomeHandler(c)
		}
		return home.AdminServersHandler(c)
	})
	e.GET("/admin/settings", func(c echo.Context) error {
		if !auth.IsAdmin(c) {
			return home.HomeHandler(c)
		}
		return home.AdminSettingsHandler(c)
	})

	e.GET("/api/logout", func(c echo.Context) error {
		return auth.Logout(c)
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
	serverauth.Routes(e)
	e.GET("/api/heath", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
	e.GET("/api/stats", func(c echo.Context) error {
		return stats.StatsHandler(c)
	})
}
