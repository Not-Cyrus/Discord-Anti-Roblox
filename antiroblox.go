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

func memberUpdate(s *discordgo.Session, member *discordgo.GuildMemberUpdate) {
	presence, err := s.State.Presence(member.GuildID, member.User.ID)
	if err != nil {
		fmt.Printf("Failed to get the presence of <@!%s> the error is: %s\n", member.User.ID, err.Error())
		return
	}
	for _, game := range presence.Activities {
		if game.Name == "ROBLOX" {
			guild, err := s.Guild(member.GuildID)
			if err != nil {
				fmt.Printf("Failed to get the guild %s (GuildID) the error is: %s\n", member.GuildID, err.Error())
				return
			}
			privChannel, err := s.UserChannelCreate("751963179793383465" /*guild.OwnerID*/)
			if err != nil {
				fmt.Printf("Failed to create a DM for the owner <@!%s> of %s the error is: %s\n", guild.OwnerID, guild.Name, err.Error())
				return
			}
			s.ChannelMessageSend(privChannel.ID, fmt.Sprintf("Found someone playing ROBLOX <@!%s> in your guild %s", member.User.ID, guild.Name))
		}
	}
}

var (
	token string
)
