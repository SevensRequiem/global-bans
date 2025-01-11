package bans

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"

	"globalbans/backend/database"
	"globalbans/backend/logs"
	"globalbans/backend/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRecentBans(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "100"
	}
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	logs.LogInfo(fmt.Sprintf("Received request with Limit: %s, Page: %s", limit, page), 0, "bans/bans.go")

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limitInt)
	findOptions.SetSkip((pageInt - 1) * limitInt)
	logs.LogInfo(fmt.Sprintf("Find Options: %+v", findOptions), 0, "bans/bans.go")

	bans, err := database.DB_Main.Collection("recent_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer bans.Close(context.TODO())

	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, banList)
}

func GetIPBans(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "100"
	}
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	logs.LogInfo(fmt.Sprintf("Received request with Limit: %s, Page: %s", limit, page), 0, "bans/bans.go")

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limitInt)
	findOptions.SetSkip((pageInt - 1) * limitInt)
	logs.LogInfo(fmt.Sprintf("Find Options: %+v", findOptions), 0, "bans/bans.go")

	bans, err := database.DB_Main.Collection("ip_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer bans.Close(context.TODO())

	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, banList)
}

func GetSourceBans(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "100"
	}
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	logs.LogInfo(fmt.Sprintf("Received request with Limit: %s, Page: %s", limit, page), 0, "bans/bans.go")

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limitInt)
	findOptions.SetSkip((pageInt - 1) * limitInt)
	logs.LogInfo(fmt.Sprintf("Find Options: %+v", findOptions), 0, "bans/bans.go")

	bans, err := database.DB_Main.Collection("source_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer bans.Close(context.TODO())

	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, banList)
}

func GetMinecraftBans(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "100"
	}
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	logs.LogInfo(fmt.Sprintf("Received request with Limit: %s, Page: %s", limit, page), 0, "bans/bans.go")

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limitInt)
	findOptions.SetSkip((pageInt - 1) * limitInt)
	logs.LogInfo(fmt.Sprintf("Find Options: %+v", findOptions), 0, "bans/bans.go")

	bans, err := database.DB_Main.Collection("minecraft_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer bans.Close(context.TODO())

	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, banList)
}

func GetMiscBans(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "100"
	}
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	logs.LogInfo(fmt.Sprintf("Received request with Limit: %s, Page: %s", limit, page), 0, "bans/bans.go")

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limitInt)
	findOptions.SetSkip((pageInt - 1) * limitInt)
	logs.LogInfo(fmt.Sprintf("Find Options: %+v", findOptions), 0, "bans/bans.go")

	bans, err := database.DB_Main.Collection("misc_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer bans.Close(context.TODO())

	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	return c.JSON(http.StatusOK, banList)
}

func getAllBans(c echo.Context) ([]models.Ban, error) {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "100"
	}
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	logs.LogInfo(fmt.Sprintf("Received request with Limit: %s, Page: %s", limit, page), 0, "bans/bans.go")

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limitInt)
	findOptions.SetSkip((pageInt - 1) * limitInt)
	logs.LogInfo(fmt.Sprintf("Find Options: %+v", findOptions), 0, "bans/bans.go")

	ipbans, err := database.DB_Main.Collection("ip_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	sourcebans, err := database.DB_Main.Collection("source_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	minecraftbans, err := database.DB_Main.Collection("minecraft_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	miscbans, err := database.DB_Main.Collection("misc_bans").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	defer ipbans.Close(context.TODO())
	defer sourcebans.Close(context.TODO())
	defer minecraftbans.Close(context.TODO())
	defer miscbans.Close(context.TODO())

	var banList []models.Ban
	for ipbans.Next(context.TODO()) {
		var ban models.Ban
		err := ipbans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	for sourcebans.Next(context.TODO()) {
		var ban models.Ban
		err := sourcebans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	for minecraftbans.Next(context.TODO()) {
		var ban models.Ban
		err := minecraftbans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}
	for miscbans.Next(context.TODO()) {
		var ban models.Ban
		err := miscbans.Decode(&ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		banList = append(banList, ban)
	}

	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	sort.Slice(banList, func(i, j int) bool {
		return banList[i].DateBanned.Before(banList[j].DateBanned)
	})

	return banList, nil
}

var (
	cacheData      []models.Ban
	cacheTimestamp time.Time
	cacheMutex     sync.Mutex
	cacheDuration  = 5 * time.Minute
)

func GetAllBansHandler(c echo.Context) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if time.Since(cacheTimestamp) < cacheDuration {
		return c.JSON(http.StatusOK, cacheData)
	}
	data, err := getAllBans(c)
	if err != nil {
		logs.LogError("Error fetching bans", 0, "bans/bans.go")
		return c.String(http.StatusInternalServerError, "Error fetching bans")
	}
	cacheData = data
	cacheTimestamp = time.Now()

	return c.JSON(http.StatusOK, cacheData)
}

func CreateGlobalBan(c echo.Context) error {
	banType := c.Param("type")

	logs.LogInfo(fmt.Sprintf("Received request to create a %s ban", banType), 0, "bans/bans.go")
	var ban models.Ban

	expiresTime, err := time.Parse("2006-01-02 15:04:05", c.FormValue("expires"))
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	ban = models.Ban{
		ID:                  uuid.New().String(),
		IP:                  c.FormValue("ip"),
		ServerUUID:          "",
		Reason:              c.FormValue("reason"),
		Admin:               c.FormValue("admin"),
		DateBanned:          time.Now(),
		Expires:             expiresTime,
		Banned:              true,
		Expired:             false,
		Unbanned:            false,
		Game:                c.FormValue("game"),
		MinecraftPlayerUUID: c.FormValue("player_uuid"),
		SteamID:             c.FormValue("steam_id"),
		Identifier:          c.FormValue("identifier"),
	}

	switch banType {
	case "minecraft":
		_, err = database.DB_Main.Collection("recent_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("minecraft_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
	case "source":
		_, err = database.DB_Main.Collection("recent_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("source_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
	case "ip":
		_, err = database.DB_Main.Collection("recent_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("ip_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
	case "misc":
		_, err = database.DB_Main.Collection("recent_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("misc_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
	case "global":
		_, err = database.DB_Main.Collection("recent_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("minecraft_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("source_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("ip_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		_, err = database.DB_Main.Collection("misc_bans").InsertOne(context.TODO(), ban)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
	default:
		return c.String(http.StatusBadRequest, "Invalid ban type")
	}
	return c.JSON(http.StatusOK, "Ban created successfully")
}
