package database

import (
	"context"
	"os"
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
		logs.LogError("Error loading .env file", 0, "database/database.go")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		logs.LogError("Error creating MongoDB client", 0, "database/database.go")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logs.LogError("Error connecting to MongoDB", 0, "database/database.go")
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logs.LogError("Error pinging MongoDB", 0, "database/database.go")
		return
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
		logs.LogError("Error creating config collection", 0, "database/database.go")
	}
}

func GetCollection(collection string) *mongo.Collection {
	return DB_Main.Collection(collection)
}
