package main

import (
	"log"
	"os"
	"time"

	"globalbans/backend/discord"
	"globalbans/backend/home"
	"globalbans/backend/routes"

	schedule "globalbans/backend/scheduler"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	renderer := home.NewTemplateRenderer("frontend/views/*.html")
	renderer.LoadTemplates()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("renderer", renderer)
			return next(c)
		}
	})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET is not set")
	}
	// Middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${id} ${time_rfc3339} ${remote_ip} > ${method} > ${uri} > ${status} ${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:_csrf,header:X-CSRF-Token",
		CookieMaxAge:   86400,
		CookieSecure:   true,
		CookieHTTPOnly: false,
		CookieName:     "_csrf",
	}))

	s := schedule.NewScheduler()
	s.ScheduleTask(schedule.Task{
		Action: func() {
			// some function
		},
		Duration: 24 * time.Hour,
	})
	go s.Run()
	routes.Routes(e)
	go discord.Start()
	e.Logger.Fatal(e.StartTLS(":8888", "certificates/cert.crt", "certificates/key.pem"))
}
