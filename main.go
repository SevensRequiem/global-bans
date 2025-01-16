package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"globalbans/backend/bans"
	"globalbans/backend/discord"
	"globalbans/backend/home"
	"globalbans/backend/routes"
	schedule "globalbans/backend/scheduler"
	"globalbans/integration/firewall"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/exp/rand"
)

func main() {
	//if file exists
	if _, err := os.Stat(".live"); err != nil {
		if os.IsNotExist(err) {
			GenerateSecret()
		} else {
			log.Fatalf("Error checking .live file: %v", err)
		}
	}

	e := echo.New()
	e.Renderer = home.NewTemplateRenderer("frontend/views/*.html")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET is not set")
	}
	// Middleware
	store := sessions.NewCookieStore([]byte(secret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	e.Use(session.Middleware(store))
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${id} ${time_rfc3339} ${remote_ip} > ${method} > ${uri} > ${status} ${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' https://kit.fontawesome.com 'unsafe-inline' https://cdn.jsdelivr.net; style-src 'self' https://kit.fontawesome.com 'unsafe-inline'; font-src 'self' https://kit.fontawesome.com; connect-src 'self' https://ka-f.fontawesome.com; img-src 'self' https://cdn.jsdelivr.net;",
		HSTSExcludeSubdomains: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:  "header:X-CSRF-Token",
		TokenLength:  32,
		ContextKey:   "csrf",
		CookieName:   "csrf",
		CookieDomain: os.Getenv("BASE_URL"),
		CookiePath:   "/",
	}))
	e.HTTPErrorHandler = home.ErrorHandler
	s30 := schedule.NewScheduler()
	s30.ScheduleTask(schedule.Task{
		Action: func() {
			bans.ExpireCheck("minecraft")
			bans.ExpireCheck("source")
			bans.ExpireCheck("misc")
			bans.ExpireCheck("ip")
			firewall.ExpireCheck()
		},
		Duration: 30 * time.Minute,
	})
	go s30.Run()

	routes.Routes(e)
	go discord.Start()
	e.Logger.Fatal(e.StartTLS(":8888", "certificates/cert.crt", "certificates/key.pem"))
}

func GenerateSecret() {
	secret := GenerateRandomString(64)
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	env["SECRET"] = secret
	err = godotenv.Write(env, ".env")
	if err != nil {
		log.Fatal("Error writing .env file")
	}
	_, err = os.Create(".live")
	if err != nil {
		log.Fatal("Error creating .live file")
	}
}

func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
