package servers

import (
	"context"
	"globalbans/backend/database"
	"globalbans/backend/logs"
	"globalbans/backend/models"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func globalServerList(c echo.Context) ([]models.Server, error) {
	var servers []models.Server

	minecraftservers := database.DB_Main.Collection("minecraft_servers")
	sourceservers := database.DB_Main.Collection("source_servers")
	miscservers := database.DB_Main.Collection("misc_servers")

	mccur, err := minecraftservers.Find(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	scur, err := sourceservers.Find(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	mcur, err := miscservers.Find(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	var minecraftServers []models.Server
	var sourceServers []models.Server
	var miscServers []models.Server

	if err = mccur.All(context.Background(), &minecraftServers); err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	if err = scur.All(context.Background(), &sourceServers); err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	if err = mcur.All(context.Background(), &miscServers); err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	servers = append(servers, minecraftServers...)
	servers = append(servers, sourceServers...)
	servers = append(servers, miscServers...)

	for i := 0; i < len(servers); i++ {
		for j := 0; j < len(servers); j++ {
			if servers[i].Game < servers[j].Game {
				servers[i], servers[j] = servers[j], servers[i]
			}
		}
	}

	return servers, nil
}

var (
	cacheData      []models.Server
	cacheTimestamp time.Time
	cacheMutex     sync.Mutex
	cacheDuration  = 60 * time.Minute
)

func GetAllServersHandler(c echo.Context) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if time.Since(cacheTimestamp) < cacheDuration {
		return c.JSON(http.StatusOK, cacheData)
	}
	data, err := globalServerList(c)
	if err != nil {
		logs.LogError("Error fetching bans", 0, "bans/bans.go")
		return c.String(http.StatusInternalServerError, "Error fetching bans")
	}
	cacheData = data
	cacheTimestamp = time.Now()

	return c.JSON(http.StatusOK, cacheData)
}

func ResetServerCache() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	cacheTimestamp = time.Time{}
}

func GetMinecraftServers(c echo.Context) error {
	var servers []models.Server
	cursor, err := database.DB_Main.Collection("minecraft_servers").Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/servers/servers.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	if err = cursor.All(context.Background(), &servers); err != nil {
		logs.LogError("Error decoding servers", 0, "integrations/servers/servers.go")
		return c.String(http.StatusInternalServerError, "Error decoding servers")
	}

	return c.JSON(http.StatusOK, servers)
}

func GetSourceServers(c echo.Context) error {
	var servers []models.Server
	cursor, err := database.DB_Main.Collection("source_servers").Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/servers/servers.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	if err = cursor.All(context.Background(), &servers); err != nil {
		logs.LogError("Error decoding servers", 0, "integrations/servers/servers.go")
		return c.String(http.StatusInternalServerError, "Error decoding servers")
	}

	return c.JSON(http.StatusOK, servers)
}

func GetMiscServers(c echo.Context) error {
	var servers []models.Server
	cursor, err := database.DB_Main.Collection("misc_servers").Find(context.TODO(), bson.M{})
	if err != nil {
		logs.LogError("Error querying the database", 0, "integrations/servers/servers.go")
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	if err = cursor.All(context.Background(), &servers); err != nil {
		logs.LogError("Error decoding servers", 0, "integrations/servers/servers.go")
		return c.String(http.StatusInternalServerError, "Error decoding servers")
	}

	return c.JSON(http.StatusOK, servers)
}
