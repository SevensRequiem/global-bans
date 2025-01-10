package discord

import (
	"fmt"
	"runtime"

	"globalbans/backend/config"
	"globalbans/backend/database"
	"globalbans/backend/logs"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var s *discordgo.Session

var db = database.DB_Main

func Start() {
	godotenv.Load()
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Println("Error creating Discord session,", err)
			logs.LogError(err.Error(), line, file)
		}
	}
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err = dg.Open()
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Println("Error opening connection,", err)
			logs.LogError(err.Error(), line, file)
		}
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	Status(dg)

	s = dg

}

func Status(s *discordgo.Session) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "Aurora",
				Type: discordgo.ActivityTypeWatching,
			},
		},
		Status: "online",
	})
}
