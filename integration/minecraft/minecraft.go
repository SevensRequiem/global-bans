package minecraft

import (
	"context"
	"encoding/json"
	"fmt"
	"globalbans/backend/database"
	"globalbans/backend/global"
	"globalbans/backend/logs"
	"globalbans/backend/models"
	"globalbans/backend/serverauth"
	"net/http"
	"os"
	"runtime"
	"time"

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

	e.POST("/api/minecraft/ban", func(c echo.Context) error {
		return MinecraftBan(c)
	})
}

func Ping(c echo.Context) error {
	if !serverauth.ValidateAPIKey(c) {
		return c.String(http.StatusUnauthorized, "Invalid API key")
	}
	ip := c.QueryParam("ip")
	port := c.QueryParam("port")
	logs.LogHTTP(fmt.Sprintf("Minecraft Ping %s:%s", ip, port), 0, "integrations/minecraft.go")
	return c.String(http.StatusOK, "Pong")
}

func MinecraftBan(c echo.Context) error {
	if !serverauth.ValidateAPIKey(c) {
		return c.String(http.StatusUnauthorized, "Invalid API key")
	}
	playerIP := c.QueryParam("playerip")
	player := c.QueryParam("player")
	reason := c.QueryParam("reason")
	expires := c.QueryParam("expires")
	admin := c.QueryParam("admin")
	serverUUID := c.QueryParam("server")

	if player == "" || reason == "" || expires == "" || admin == "" || serverUUID == "" {
		return c.String(http.StatusBadRequest, "Invalid parameters")
	}

	filter := bson.M{"server_id": serverUUID}
	var existingServer models.Server
	err := database.DB_Main.Collection("minecraft_servers").FindOne(context.TODO(), filter).Decode(&existingServer)
	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "Server not found")
	} else if err != nil {
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	filter = bson.M{"player": player, "server_uuid": serverUUID}
	var existingBan models.Ban
	err = database.DB_Main.Collection("minecraft_bans").FindOne(context.TODO(), filter).Decode(&existingBan)
	if err != mongo.ErrNoDocuments {
		return c.String(http.StatusConflict, "Player is already banned")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://playerdb.co/api/player/minecraft/%s", player), nil)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	req.Header.Set("User-Agent", "GlobalBans "+global.GetVersion()+" - "+os.Getenv("BASE_URL"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer resp.Body.Close()

	type playerDataPlayer struct {
		Username string `json:"username"`
		ID       string `json:"id"`
	}

	type data struct {
		Player playerDataPlayer `json:"player"`
	}

	type response struct {
		Data data `json:"data"`
	}

	var playerData response

	err = json.NewDecoder(resp.Body).Decode(&playerData)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	if playerData.Data.Player.Username == "" {
		return c.String(http.StatusNotFound, "Player not found")
	}
	parsedExpires, err := time.Parse("2006-01-02", expires)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	uuid := uuid.New().String()
	ban := models.Ban{
		ID:                  uuid,
		IP:                  playerIP,
		Identifier:          player,
		Reason:              reason,
		Expires:             parsedExpires,
		Admin:               admin,
		ServerUUID:          serverUUID,
		MinecraftPlayerUUID: playerData.Data.Player.ID,
		Game:                "minecraft",
		Banned:              true,
	}

	_, err = database.DB_Main.Collection("minecraft_bans").InsertOne(context.TODO(), ban)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	_, err = database.DB_Main.Collection("recent_bans").InsertOne(context.TODO(), ban)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"uuid":    uuid,
		"message": "Ban created successfully",
	})
}

func Server(c echo.Context) error {
	if !serverauth.ValidateAPIKey(c) {
		return c.String(http.StatusUnauthorized, "Invalid API key")
	}
	ip := c.QueryParam("ip")
	port := c.QueryParam("port")

	if ip == "" || port == "" {
		return c.String(http.StatusBadRequest, "IP and port must be provided")
	}

	filter := bson.M{"ip": ip, "port": port}
	var existingServer models.Server
	err := database.DB_Main.Collection("minecraft_servers").FindOne(context.TODO(), filter).Decode(&existingServer)

	if err == mongo.ErrNoDocuments {
		uuid := uuid.New().String()

		server := models.Server{
			ID:       uuid,
			IP:       ip,
			Port:     port,
			ServerID: uuid,
			Game:     "minecraft",
		}

		_, err = database.DB_Main.Collection("minecraft_servers").InsertOne(context.TODO(), server)
		if err != nil {
			logs.LogError(fmt.Sprintf("Error inserting new server with IP %s and Port %s: %v", ip, port, err), 0, "integrations/minecraft.go")
			return c.String(http.StatusInternalServerError, "Error inserting new server")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"uuid":    uuid,
			"message": "Server created successfully",
		})
	} else if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	// If the server exists, return its UUID
	return c.String(http.StatusOK, existingServer.ServerID)
}

func Banlist(c echo.Context) error {
	if !serverauth.ValidateAPIKey(c) {
		return c.String(http.StatusUnauthorized, "Invalid API key")
	}
	uuid := c.QueryParam("uuid")

	// Validate parameters
	if uuid == "" {
		return c.String(http.StatusBadRequest, "Server ID must be provided")
	}

	// Check if server with given UUID exists
	filter := bson.M{"server_uuid": uuid}
	var existingServer models.Server
	err := database.DB_Main.Collection("minecraft_servers").FindOne(context.TODO(), filter).Decode(&existingServer)

	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "Server not found")
	} else if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	// Get all bans for the server
	var bans []models.Ban
	cursor, err := database.DB_Main.Collection("bans").Find(context.TODO(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	if err = cursor.All(context.Background(), &bans); err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, bans)
}

func SelfBanlist(c echo.Context) error {
	if !serverauth.ValidateAPIKey(c) {
		return c.String(http.StatusUnauthorized, "Invalid API key")
	}
	uuid := c.QueryParam("uuid")

	// Validate parameters
	if uuid == "" {
		return c.String(http.StatusBadRequest, "Server ID must be provided")
	}

	// Check if server with given UUID exists
	filter := bson.M{"server_uuid": uuid}
	var existingServer models.Server
	err := database.DB_Main.Collection("minecraft_bans").FindOne(context.TODO(), filter).Decode(&existingServer)

	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusNotFound, "Server not found")
	} else if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	// Get all self bans for the server
	var selfBans []models.Ban
	cursor, err := database.DB_Main.Collection("minecraft_bans").Find(context.TODO(), filter)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	if err = cursor.All(context.Background(), &selfBans); err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, selfBans)
}
