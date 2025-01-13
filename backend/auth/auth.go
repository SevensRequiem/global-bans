package auth

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"

	"globalbans/backend/config"
	"globalbans/backend/database"
	"globalbans/backend/logs"
	"globalbans/backend/models"
)

var oauthConf *oauth2.Config
var ClientID string
var ClientSecret string

func init() {
	gob.Register(models.User{})

	DiscordClientID := config.ClientID
	DiscordClientSecret := config.ClientSecret
	DiscordRedirectURI := os.Getenv("DISCORD_REDIRECT_URI")
	oauthConf = &oauth2.Config{
		ClientID:     DiscordClientID,
		ClientSecret: DiscordClientSecret,
		RedirectURL:  DiscordRedirectURI,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
		Scopes: []string{"identify", "email"},
	}
}

func CallbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	client := oauthConf.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer resp.Body.Close()

	user := models.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	fmt.Println(user)

	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the user already exists in the database
	filter := bson.M{"_id": user.ID}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		user.DateCreated = time.Now()
		user.DoesExist = true
		_, err = collection.InsertOne(ctx, user)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
	} else if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	// Store the user in the session
	sess, err := session.Get("session", c)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user"] = user

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func LoginHandler(c echo.Context) error {
	url := oauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func LogoutHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{MaxAge: -1}
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func AdminCheck(c echo.Context) bool {
	collection := database.GetCollection("users")

	sess, err := session.Get("session", c)
	if err != nil {
		return false
	}

	userSessionValue, ok := sess.Values["user"]
	if !ok {
		return false
	}

	user, ok := userSessionValue.(models.User)
	if !ok {
		return false
	}

	var userFromDB models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, bson.M{"_id": user.ID}).Decode(&userFromDB)
	if err != nil {
		return false
	}

	if userFromDB.Groups.Admin {
		return true
	}

	return false
}

func GetUserByID(userID string) (*models.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := models.User{}
	err := collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return &user, nil
}

func DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return nil
}

func GetUsers() ([]models.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUsersByGroup(group string) ([]models.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := collection.Find(ctx, bson.M{"groups": group})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetTotalUsers() int {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0
	}
	return int(count)
}

func GetCurrentUser(c echo.Context) (*models.LoggedInUser, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
		return nil, err
	}

	userSessionValue, ok := sess.Values["user"]
	if !ok {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError("user not found in session", line, file)
		}
		return nil, errors.New("user not found in session")
	}

	user, ok := userSessionValue.(models.LoggedInUser)
	if !ok {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return &user, nil
}

func CheckLoggedIn(c echo.Context) (*models.LoggedInUser, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return nil, err
	}

	userSessionValue, ok := sess.Values["user"]
	if !ok {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	user, ok := userSessionValue.(models.User)
	if !ok {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return &models.LoggedInUser{ID: user.ID, Username: user.Username}, nil
}

func UpdateGroups(c echo.Context, userID string, groups string) error {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"groups": groups}})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return nil
}

func RefreshSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	userSessionValue, ok := sess.Values["user"]
	if !ok {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	user, ok := userSessionValue.(models.User)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	userFromDB, err := GetUserByID(user.ID)
	if err != nil {
		return err
	}

	sess.Values["user"] = *userFromDB

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return nil
}

func GetUserFromContext(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return nil, fmt.Errorf("failed to get user from context")
	}

	return &user, nil
}
