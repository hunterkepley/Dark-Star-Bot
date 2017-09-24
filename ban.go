package main

import (
	"fmt"
	"log"
	
	"github.com/bwmarrin/discordgo"
)

func banMessage(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	config, err := getConfigForGuildId(e.GuildID)
	if err != nil {
		log.Fatal(err)
	}

	s.ChannelMessageSendEmbed(config.banChannelId, &discordgo.MessageEmbed{
		Title:       "Banned!",
		Description: fmt.Sprintf(config.banMessage, e.User.Username)})
}
