package stats

import (
	"context"
	"globalbans/backend/database"
	"globalbans/backend/logs"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Weekday string

const (
	Monday    Weekday = "Monday"
	Tuesday   Weekday = "Tuesday"
	Wednesday Weekday = "Wednesday"
	Thursday  Weekday = "Thursday"
	Friday    Weekday = "Friday"
	Saturday  Weekday = "Saturday"
	Sunday    Weekday = "Sunday"
)

type Day struct {
	Weekday Weekday `bson:"weekday"`
	Bans    int     `bson:"bans"`
}

func PlusWeeklyBan(bans int, day time.Time) {
	weekday := Weekday(day.Weekday().String())
	if weekday == Sunday {
		ClearWeekly()
	}
	_, err := database.DB_Main.Collection("stats").UpdateOne(context.Background(), bson.M{"_id": weekday}, bson.M{"$inc": bson.M{"bans": bans}})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
}

func ClearWeekly() {
	_, err := database.DB_Main.Collection("stats").UpdateOne(context.Background(), bson.M{}, bson.M{"$set": bson.M{"bans": 0}})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
}

func GetWeeklyStats() []Day {
	var days []Day
	cur, err := database.DB_Main.Collection("stats").Find(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var day Day
		err := cur.Decode(&day)
		if err != nil {
			_, file, line, ok := runtime.Caller(1)
			if ok {
				logs.LogError(err.Error(), line, file)
			}
		}
		days = append(days, day)
	}
	return days
}

func TotalServers() int {
	mcservers, err := database.DB_Main.Collection("minecraft_servers").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	srcservers, err := database.DB_Main.Collection("source_servers").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	miscservers, err := database.DB_Main.Collection("misc_servers").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	count := mcservers + srcservers + miscservers
	return int(count)
}

func TotalBans() int {
	mcbans, err := database.DB_Main.Collection("minecraft_bans").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	srcbans, err := database.DB_Main.Collection("source_bans").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	miscbans, err := database.DB_Main.Collection("misc_bans").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	count := mcbans + srcbans + miscbans
	return int(count)
}

type Stat struct {
	TotalServers int `json:"total_servers"`
	TotalBans    int `json:"total_bans"`
}

var (
	cacheData      []Stat
	cacheTimestamp time.Time
	cacheMutex     sync.Mutex
	cacheDuration  = 5 * time.Minute
)

func StatsHandler(c echo.Context) error {
	totalServers := TotalServers()
	totalBans := TotalBans()
	stats := Stat{
		TotalServers: totalServers,
		TotalBans:    totalBans,
	}

	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if time.Since(cacheTimestamp) > cacheDuration {
		cacheData = []Stat{stats}
		cacheTimestamp = time.Now()
	}

	return c.JSON(http.StatusOK, stats)
}
