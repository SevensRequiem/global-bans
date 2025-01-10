package database

import (
	"context"
	"os"
	"runtime"
	"time"

	"globalbans/backend/logs"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB_Main *mongo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError(err.Error(), line, file)
		}
	}
	DB_Main = client.Database(os.Getenv("MONGO_DB"))

	if DB_Main != nil {
		DB_Main.CreateCollection(context.Background(), "config")
		DB_Main.CreateCollection(context.Background(), "banned")
		DB_Main.CreateCollection(context.Background(), "expired")
		DB_Main.CreateCollection(context.Background(), "deleted")
		DB_Main.CreateCollection(context.Background(), "firewalls")
		DB_Main.CreateCollection(context.Background(), "filesyncs")
		DB_Main.CreateCollection(context.Background(), "rcons")
	} else {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			logs.LogError("Database is nil", line, file)
		}
	}
}

func GetCollection(collection string) *mongo.Collection {
	return DB_Main.Collection(collection)
}
