package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func banMessage(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	s.ChannelMessageSendEmbed(welcomeChannelID, &discordgo.MessageEmbed{
		Title:       "BANNED",
		Description: fmt.Sprintf("Girl, bye %s!", makeMentionRegular(e.User.ID))})
}
