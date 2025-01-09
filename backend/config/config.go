package config

import (
	"context"
	"fmt"

	"globalbans/backend/database"
	"globalbans/backend/logs"
	"globalbans/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db = database.DB_Main
var Token string
var ClientID string
var ClientSecret string
var GuildID string
var ChannelID string

func init() {
	EnsureConfig()
	cfg := db.Collection("config").FindOne(context.Background(), bson.M{})
	if cfg.Err() != nil {
		logs.LogFatal("Error fetching config from MongoDB", 0, "discord/discord.go")
	}

	var config models.Config

	err := cfg.Decode(&config)
	if err != nil {
		logs.LogFatal("Error decoding config from MongoDB", 0, "discord/discord.go")
		fmt.Println(err)
	}

	Token = config.Token
	ClientID = config.ClientID
	ClientSecret = config.ClientSecret
	GuildID = config.GuildID

}

func EnsureConfig() {
	cfg := db.Collection("config").FindOne(context.Background(), bson.M{})
	if cfg.Err() != nil {
		if cfg.Err() == mongo.ErrNoDocuments {
			config := models.Config{
				Token:        "",
				ClientID:     "",
				ClientSecret: "",
				GuildID:      "",
			}

			_, err := db.Collection("config").InsertOne(context.Background(), config)
			if err != nil {
				logs.LogFatal("Error inserting config into MongoDB", 0, "discord/discord.go")
			}
		} else {
			logs.LogFatal("Error fetching config from MongoDB", 0, "discord/discord.go")
		}
	}
}

func ReloadConfig() {
	cfg := db.Collection("config").FindOne(context.Background(), bson.M{})
	if cfg.Err() != nil {
		logs.LogFatal("Error fetching config from MongoDB", 0, "discord/discord.go")
	}

	var config models.Config

	err := cfg.Decode(&config)
	if err != nil {
		logs.LogFatal("Error decoding config from MongoDB", 0, "discord/discord.go")
		fmt.Println(err)
	}

	Token = config.Token
	ClientID = config.ClientID
	ClientSecret = config.ClientSecret
	GuildID = config.GuildID
}
