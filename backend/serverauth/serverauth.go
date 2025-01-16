package serverauth

import (
	"context"
	"globalbans/backend/database"
	"globalbans/backend/models"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Routes(e *echo.Echo) {
	e.GET("/generateapikey", GenerateAPIKey)
}

func GenerateAPIKey(c echo.Context) error {
	apiKey := uuid.New().String()
	_, err := database.DB_Main.Collection("apikeys").InsertOne(context.Background(), bson.M{"apikey": apiKey})
	if err != nil {
		return c.JSON(500, "Failed to generate API key")
	}
	return c.JSON(200, apiKey)
}

func ValidateAPIKey(c echo.Context) bool {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return false
	}
	apiKey := strings.TrimPrefix(authorizationHeader, "Bearer ")

	filter := bson.M{"apikey": apiKey}
	var existingAPIKey models.APIKey
	err := database.DB_Main.Collection("apikeys").FindOne(context.TODO(), filter).Decode(&existingAPIKey)
	if err != nil {
		return false
	}
	return true
}
