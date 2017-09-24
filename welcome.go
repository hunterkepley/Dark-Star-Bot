package main

import (
	"fmt"
	"log"
	
	"github.com/bwmarrin/discordgo"
)

func welcomeMessage(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	config, err := getConfigForGuildId(e.GuildID)
	if err != nil {
		log.Fatal(err)
	}

	s.ChannelMessageSendEmbed(config.welcomeChannelId, &discordgo.MessageEmbed{
		Title:       "Welcome!",
		Description: fmt.Sprintf(config.welcomeMessage, makeMentionRegular(e.User.ID))})
}
