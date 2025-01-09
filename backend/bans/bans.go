package bans

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"globalbans/backend/database"
	"globalbans/backend/logs"
	"globalbans/backend/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBan(c echo.Context) error {
	id := c.Param("id")
	var ban models.Ban
	err := database.DB_Main.Collection("banned").FindOne(context.TODO(), models.Ban{ID: id}).Decode(&ban)
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ban)
}
func GetAllBans(c echo.Context) error {
	bans, err := database.DB_Main.Collection("banned").Find(context.Background(), bson.M{})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var banList = []models.Ban{}
	var ban models.Ban

	// Iterate over the cursor
	for bans.Next(context.Background()) {
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		banList = append(banList, ban)
	}

	defer bans.Close(context.Background())
	if len(banList) == 0 {
		logs.LogError("No bans found", 0, "bans/bans.go")
		return c.JSON(http.StatusOK, []string{})
	}
	return c.JSON(http.StatusOK, banList)
}

func GetBans(c echo.Context) error {
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
		logs.LogError(fmt.Sprintf("Error parsing limit: %s", err.Error()), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, "Invalid limit parameter")
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logs.LogError(fmt.Sprintf("Error parsing page: %s", err.Error()), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, "Invalid page parameter")
	}

	findOptions := options.Find()
	findOptions.SetLimit(limitInt)
	findOptions.SetSkip((pageInt - 1) * limitInt)
	logs.LogInfo(fmt.Sprintf("Find Options: %+v", findOptions), 0, "bans/bans.go")

	bans, err := database.DB_Main.Collection("banned").Find(context.TODO(), models.Ban{}, findOptions)
	if err != nil {
		logs.LogError(fmt.Sprintf("Error finding bans: %s", err.Error()), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, "Error retrieving bans")
	}
	defer bans.Close(context.TODO())

	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(fmt.Sprintf("Error decoding ban: %s", err.Error()), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, "Error decoding ban data")
		}
		banList = append(banList, ban)
	}
	if err := bans.Err(); err != nil {
		logs.LogError(fmt.Sprintf("Cursor error: %s", err.Error()), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, "Error iterating over bans")
	}
	logs.LogInfo(fmt.Sprintf("Ban List: %+v", banList), 0, "bans/bans.go")
	return c.JSON(http.StatusOK, banList)
}

func GetBanByID(c echo.Context) error {
	id := c.Param("id")
	var ban models.Ban
	err := database.DB_Main.Collection("banned").FindOne(context.TODO(), models.Ban{ID: id}).Decode(&ban)
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ban)
}

func GetBansByIP(c echo.Context) error {
	ip := c.Param("ip")
	bans, err := database.DB_Main.Collection("banned").Find(context.TODO(), models.Ban{IP: ip})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer bans.Close(context.TODO())
	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		banList = append(banList, ban)
	}
	return c.JSON(http.StatusOK, banList)
}

func GetBansBySteamID(c echo.Context) error {
	steamID := c.Param("steamid")
	bans, err := database.DB_Main.Collection("banned").Find(context.TODO(), models.Ban{SteamID: steamID})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer bans.Close(context.TODO())
	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		banList = append(banList, ban)
	}
	return c.JSON(http.StatusOK, banList)
}

func GetBansByDiscordID(c echo.Context) error {
	discordID := c.Param("discordid")
	bans, err := database.DB_Main.Collection("banned").Find(context.TODO(), models.Ban{DiscordID: discordID})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer bans.Close(context.TODO())
	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		banList = append(banList, ban)
	}
	return c.JSON(http.StatusOK, banList)
}

func GetBansByReason(c echo.Context) error {
	reason := c.Param("reason")
	bans, err := database.DB_Main.Collection("banned").Find(context.TODO(), models.Ban{Reason: reason})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer bans.Close(context.TODO())
	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		banList = append(banList, ban)
	}
	return c.JSON(http.StatusOK, banList)
}

func GetBansByAdmin(c echo.Context) error {
	admin := c.Param("admin")
	bans, err := database.DB_Main.Collection("banned").Find(context.TODO(), models.Ban{Admin: admin})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer bans.Close(context.TODO())
	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		banList = append(banList, ban)
	}
	return c.JSON(http.StatusOK, banList)
}

func GetBansByServer(c echo.Context) error {
	serverip := c.Param("ip")
	serverport := c.Param("port")
	bans, err := database.DB_Main.Collection("banned").Find(context.TODO(), models.Ban{ServerIP: serverip, ServerPort: serverport})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer bans.Close(context.TODO())
	var banList []models.Ban
	for bans.Next(context.TODO()) {
		var ban models.Ban
		err := bans.Decode(&ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		banList = append(banList, ban)
	}
	return c.JSON(http.StatusOK, banList)
}

func SearchBans(c echo.Context) error {
	switch c.Param("type") {
	case "ip":
		return GetBansByIP(c)
	case "steamid":
		return GetBansBySteamID(c)
	case "discordid":
		return GetBansByDiscordID(c)
	case "reason":
		return GetBansByReason(c)
	case "admin":
		return GetBansByAdmin(c)
	case "server":
		return GetBansByServer(c)
	default:
		return GetBans(c)
	}
}

func CreateBan(c echo.Context) error {
	var ban models.Ban
	if err := c.Bind(&ban); err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err := database.DB_Main.Collection("bans").InsertOne(context.TODO(), ban)
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ban)
}

func DeleteBan(c echo.Context) error {
	id := c.Param("id")
	_, err := database.DB_Main.Collection("banned").DeleteOne(context.TODO(), models.Ban{ID: id})
	if err != nil {
		logs.LogError(err.Error(), 0, "bans/bans.go")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

func IngestBans(c echo.Context) error {
	return nil
}

func DummyData() error {
	const amount = 30

	for i := 0; i < amount; i++ {
		ban := models.Ban{
			ID:          strconv.Itoa(i),
			IP:          fmt.Sprintf("192.168.1.%d", i),
			SteamID:     strconv.Itoa(i),
			DiscordID:   strconv.Itoa(i),
			MinecraftID: strconv.Itoa(i),
			MiscID:      strconv.Itoa(i),
			Username:    fmt.Sprintf("user%d", i),
			Reason:      fmt.Sprintf("reason%d", i),
			Admin:       fmt.Sprintf("admin%d", i),
			Game:        fmt.Sprintf("game%d", i),
			DateBanned:  time.Now().Format(time.RFC3339),
			Expires:     time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		}
		_, err := database.DB_Main.Collection("banned").InsertOne(context.TODO(), ban)
		if err != nil {
			logs.LogError(err.Error(), 0, "bans/bans.go")
			continue
		}
	}
	return nil
}
