package minecraft

import (
	"context"
	"fmt"
	"globalbans/backend/database"
	"globalbans/backend/logs"
	"globalbans/backend/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Routes(e *echo.Echo) {
	e.GET("/api/minecraft/ping", func(c echo.Context) error {
		return Ping(c)
	})
	e.POST("/api/minecraft/server", func(c echo.Context) error {
		return Server(c)
	})
}

func Ping(c echo.Context) error {
	ip := c.Param("ip")
	port := c.Param("port")
	logs.LogHTTP(fmt.Sprintf("Minecraft Ping %s:%s", ip, port), 0, "integrations/minecraft.go")
	return c.String(http.StatusOK, "Pong")
}

func Server(c echo.Context) error {
	ip := c.Param("ip")
	port := c.Param("port")

	// Check if server with given IP and port exists
	filter := bson.M{"ip": ip, "port": port}
	var existingServer models.Server
	err := database.DB_Main.Collection("servers").FindOne(context.TODO(), filter).Decode(&existingServer)
	if err == nil {
		return c.String(http.StatusOK, existingServer.UUID)
	}

	uuid := uuid.New().String()

	server := models.Server{
		ID:   uuid,
		IP:   ip,
		Port: port,
		UUID: uuid,
		Game: "minecraft",
	}

	_, err = database.DB_Main.Collection("servers").InsertOne(context.TODO(), server)
	if err != nil {
		logs.LogError("Error inserting new server", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error inserting new server")
	}

	return c.String(http.StatusOK, uuid)
}
