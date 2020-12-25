package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Print("Enter your token: ")
	fmt.Scan(&token)
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		panic("Wrong token")
	}
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	s.AddHandler(memberUpdate)
	s.Open()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	s.Close()
}

func memberUpdate(s *discordgo.Session, presence *discordgo.PresenceUpdate) {
	for _, game := range presence.Activities {
		if game.Name == "ROBLOX" {
			guild, err := s.Guild(presence.GuildID)
			if err != nil {
				fmt.Printf("Failed to get the guild %s (GuildID) the error is: %s\n", presence.GuildID, err.Error())
				return
			}
			privChannel, err := s.UserChannelCreate(guild.OwnerID)
			if err != nil {
				fmt.Printf("Failed to create a DM for the owner <@!%s> of %s the error is: %s\n", guild.OwnerID, guild.Name, err.Error())
				return
			}
			s.ChannelMessageSend(privChannel.ID, fmt.Sprintf("Found someone playing ROBLOX <@!%s> in your guild %s", presence.User.ID, guild.Name))
		}
	}
}

var (
	token string
)
