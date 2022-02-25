package main

import (
	"fmt"
	"goDiscordBots/goBotTemplate/fetch"
	"time"

	"github.com/bwmarrin/discordgo"
)

var i int

func main() {

	bot, err := discordgo.New("Bot " + "ODc0NDI1NjQ0MzM3OTkxNzMw.YRGyVw.zXnrGo8Y-ZQX51xbv8UFg_obLdE") //figure out how to get the token from .env
	if err != nil {
		panic(err)
	}

	// register events
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)

	err = bot.Open() // creates web socket to Discord
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	for {
		if i == 1 {
			time.Sleep(time.Second * 3)
			return
		}
	} // this for loop keeps the bot running until i == 1
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("WE ARE LIVE!")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!enditall" {
		i = 1
		s.ChannelMessageSend(m.ChannelID, "Goodbye!")
		return
	}

	if m.Content == "!chuck" {
		fetch.GetChuckNorrisJoke(s, m)
	}

}
