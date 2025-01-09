package discord

import (
	"fmt"

	"globalbans/backend/config"
	"globalbans/backend/database"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var s *discordgo.Session

var db = database.DB_Main

func Start() {
	godotenv.Load()
	fmt.Println("Starting Discord bot", config.Token)
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
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
