package auth

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/scrypt"

	"globalbans/backend/database"
	"globalbans/backend/models"
)

var secret string
var keyLen = 32

func init() {
	gob.Register(models.User{})
	secret = os.Getenv("SECRET")
}
func AuthCheck(c echo.Context) bool {
	sess, _ := session.Get("session", c)
	if sess.Values["username"] != nil {
		return true
	}
	return false
}

func IsAdmin(c echo.Context) bool {
	sess, _ := session.Get("session", c)
	if sess.Values["group"] == "admin" {
		return true
	}
	return false
}

func IsMod(c echo.Context) bool {
	sess, _ := session.Get("session", c)
	if sess.Values["group"] == "mod" {
		return true
	}
	return false
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var login LoginData
	if err := c.Bind(&login); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	var result models.User
	err := database.DB_Main.Collection("users").FindOne(c.Request().Context(), bson.M{"username": login.Username}).Decode(&result)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}
	if !checkPassword(login.Password, result.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	session, err := session.Get("session", c)
	if err != nil {
		log.Println("Error getting session:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Domain:   os.Getenv("ADMIN_URL"),
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	}
	session.Values["username"] = login.Username
	if result.Groups.Admin {
		session.Values["group"] = "admin"
	} else if result.Groups.Mod {
		session.Values["group"] = "mod"
	}
	if err := session.Save(c.Request(), c.Response()); err != nil {
		log.Println("Error saving session:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
	return c.JSON(http.StatusOK, map[string]string{"success": "success"})
}

func Logout(c echo.Context) error {
	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		Domain:   os.Getenv("ADMIN_URL"),
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	}
	session.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/")
}

func hashPassword(password string) (string, error) {

	hash, err := scrypt.Key([]byte(password), []byte(secret), 1<<14, 8, 1, keyLen)
	if err != nil {
		return "", err
	}

	secrethash := base64.StdEncoding.EncodeToString([]byte(secret)) + "." + base64.StdEncoding.EncodeToString(hash)
	return secrethash, nil
}

func checkPassword(password, hash string) bool {
	saltHash := strings.Split(hash, ".")
	if len(saltHash) != 2 {
		return false // Invalid hash format
	}

	saltBytes, err := base64.StdEncoding.DecodeString(saltHash[0])
	if err != nil {
		return false // Failed to decode salt
	}

	hashBytes, err := base64.StdEncoding.DecodeString(saltHash[1])
	if err != nil {
		return false // Failed to decode hash
	}

	newHash, err := scrypt.Key([]byte(password), saltBytes, 1<<14, 8, 1, keyLen)
	if err != nil {
		return false
	}

	return bytes.Equal(hashBytes, newHash)
}

func defaultAdmin() {
	password, err := hashPassword("admin")
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}
	user := models.User{
		Username: "admin",
		Password: password,
		Groups: models.Group{
			Admin: true,
		},
	}
	_, err = database.DB_Main.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal("Error inserting default admin:", err)
	}
}
