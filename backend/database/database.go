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
		DB_Main.CreateCollection(context.Background(), "filesyncs")
		DB_Main.CreateCollection(context.Background(), "rcons")
		DB_Main.CreateCollection(context.Background(), "stats")

		DB_Main.CreateCollection(context.Background(), "firewall_servers")
		DB_Main.Collection("firewall_servers").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "firewall_bans")
		DB_Main.Collection("firewall_bans").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})

		DB_Main.CreateCollection(context.Background(), "fail2ban_servers")
		DB_Main.Collection("fail2ban_servers").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "fail2ban_bans")
		DB_Main.Collection("fail2ban_bans").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})

		DB_Main.CreateCollection(context.Background(), "apikeys")
		DB_Main.Collection("apikeys").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"apikey": 1},
			Options: options.Index().SetUnique(true),
		})

		DB_Main.CreateCollection(context.Background(), "recent_bans")
		DB_Main.Collection("recent_bans").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "recent_expired")
		DB_Main.Collection("recent_expired").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})

		DB_Main.CreateCollection(context.Background(), "ip_bans")
		DB_Main.Collection("ip_bans").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "ip_expired")
		DB_Main.Collection("ip_expired").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})

		DB_Main.CreateCollection(context.Background(), "minecraft_servers")
		DB_Main.Collection("minecraft_servers").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"server_id": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "minecraft_bans")
		DB_Main.Collection("minecraft_bans").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "minecraft_expired")
		DB_Main.Collection("minecraft_expired").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})

		DB_Main.CreateCollection(context.Background(), "source_servers")
		DB_Main.Collection("source_servers").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"server_id": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "source_bans")
		DB_Main.Collection("source_bans").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "source_expired")
		DB_Main.Collection("source_expired").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})

		DB_Main.CreateCollection(context.Background(), "misc_servers")
		DB_Main.Collection("misc_servers").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"server_id": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "misc_bans")
		DB_Main.Collection("misc_bans").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})
		DB_Main.CreateCollection(context.Background(), "misc_expired")
		DB_Main.Collection("misc_expired").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    map[string]interface{}{"ip": 1},
			Options: options.Index().SetUnique(true),
		})

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
