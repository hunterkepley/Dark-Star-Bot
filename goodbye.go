package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func goodbyeMessage(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	config, err := getConfigForGuildId(e.GuildID)
	if err != nil {
		log.Fatal(err)
	}

	s.ChannelMessageSendEmbed(config.goodbyeChannelId, &discordgo.MessageEmbed{
		Title:       "Later!",
		Description: fmt.Sprintf(config.goodbyeMessage, e.User.Username)})
}
