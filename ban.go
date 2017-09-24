package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func banMessage(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	config, err := getConfigForGuildID(e.GuildID)
	if err != nil {
		log.Fatal(err)
	}

	s.ChannelMessageSendEmbed(config.banChannelID, &discordgo.MessageEmbed{
		Title:       "Banned!",
		Description: fmt.Sprintf(config.banMessage, e.User.Username)})
}
