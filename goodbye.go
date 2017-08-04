package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func goodbyeMessage(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	s.ChannelMessageSendEmbed(goodbyeChannelID, &discordgo.MessageEmbed{
		Title:       "Later!",
		Description: fmt.Sprintf("Goodbye, %s!", makeMentionRegular(e.User.ID))})
}
