package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func welcomeMessage(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	config, err := getConfigForGuildID(e.GuildID)
	if err != nil {
		log.Fatal(err)
	}

	s.ChannelMessageSendEmbed(config.welcomeChannelID, &discordgo.MessageEmbed{
		Title:       "Welcome!",
		Description: fmt.Sprintf(config.welcomeMessage, e.User.Mention())})
}
