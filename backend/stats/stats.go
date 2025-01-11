package stats

import (
	"context"
	"globalbans/backend/database"
	"globalbans/backend/logs"
	"runtime"
	"time"

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
