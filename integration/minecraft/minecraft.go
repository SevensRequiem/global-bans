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
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(e *echo.Echo) {
	e.GET("/api/minecraft/ping", func(c echo.Context) error {
		return Ping(c)
	})
	e.POST("/api/minecraft/server", func(c echo.Context) error {
		return Server(c)
	})
	e.GET("/api/minecraft/banlist", func(c echo.Context) error {
		return Banlist(c)
	})
	e.GET("/api/minecraft/selfbanlist", func(c echo.Context) error {
		return SelfBanlist(c)
	})
}

func Ping(c echo.Context) error {
	ip := c.QueryParam("ip")
	port := c.QueryParam("port")
	logs.LogHTTP(fmt.Sprintf("Minecraft Ping %s:%s", ip, port), 0, "integrations/minecraft.go")
	return c.String(http.StatusOK, "Pong")
}

func Server(c echo.Context) error {
	ip := c.QueryParam("ip")
	port := c.QueryParam("port")

	// Validate parameters
	if ip == "" || port == "" {
		return c.String(http.StatusBadRequest, "IP and port must be provided")
	}

	// Check if server with given IP and port exists
	filter := bson.M{"ip": ip, "port": port}
	var existingServer models.Server
	err := database.DB_Main.Collection("servers").FindOne(context.TODO(), filter).Decode(&existingServer)

	if err == mongo.ErrNoDocuments {
		// Server doesn't exist, create new one
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
			logs.LogError(fmt.Sprintf("Error inserting new server with IP %s and Port %s: %v", ip, port, err), 0, "integrations/minecraft.go")
			return c.String(http.StatusInternalServerError, "Error inserting new server")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"uuid":    uuid,
			"message": "Server created successfully",
		})
	} else if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	// If the server exists, return its UUID
	return c.String(http.StatusOK, existingServer.UUID)
}

func Banlist(c echo.Context) error {
	uuid := c.QueryParam("uuid")

	// Validate parameters
	if uuid == "" {
		return c.String(http.StatusBadRequest, "Server ID must be provided")
	}

	// Check if server with given UUID exists
	filter := bson.M{"server_uuid": uuid}
	var existingServer models.Server
	err := database.DB_Main.Collection("servers").FindOne(context.TODO(), filter).Decode(&existingServer)

	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "Server not found")
	} else if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	// Get all bans for the server
	var bans []models.Ban
	cursor, err := database.DB_Main.Collection("bans").Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	if err = cursor.All(context.Background(), &bans); err != nil {
		logs.LogError("Error decoding bans", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error decoding bans")
	}

	return c.JSON(http.StatusOK, bans)
}

func SelfBanlist(c echo.Context) error {
	uuid := c.QueryParam("uuid")

	// Validate parameters
	if uuid == "" {
		return c.String(http.StatusBadRequest, "Server ID must be provided")
	}

	// Check if server with given UUID exists
	filter := bson.M{"server_uuid": uuid}
	var existingServer models.Server
	err := database.DB_Main.Collection("servers").FindOne(context.TODO(), filter).Decode(&existingServer)

	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "Server not found")
	} else if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	// Get all self bans for the server
	var selfBans []models.Ban
	cursor, err := database.DB_Main.Collection("bans").Find(context.TODO(), filter)
	if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	if err = cursor.All(context.Background(), &selfBans); err != nil {
		logs.LogError("Error decoding self bans", 0, "integrations/minecraft.go")
		return c.String(http.StatusInternalServerError, "Error decoding self bans")
	}

	return c.JSON(http.StatusOK, selfBans)
}
