package fetch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Joke struct {
	Value string `json:"value"`
}

func GetChuckNorrisJoke(s *discordgo.Session, m *discordgo.MessageCreate) {

	url := "https://api.chucknorris.io/jokes/random"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "I had trouble fetching a joke for you. Try again later.")
	}
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	chuckJoke := Joke{}
	jsonErr := json.Unmarshal(body, &chuckJoke)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	resEmbed := &discordgo.MessageEmbed{
		//Title:       "Chuck Norris Joke",
		Description: chuckJoke.Value,
		Color:       0x00ff00, // Green
	}

	s.ChannelMessageSendEmbed(m.ChannelID, resEmbed)
}
