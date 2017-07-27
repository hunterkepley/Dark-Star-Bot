package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func welcomeMessage(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	s.ChannelMessageSendEmbed(welcomeChannelID, &discordgo.MessageEmbed{
		Title:       "Welcome!",
		Description: fmt.Sprintf("Welcome to the server %s! Please go to the #tag_role_assignment room and set your lane and rank using `$role`. Example; `$role mid gold` or `$role top adc platinum`", makeMentionRegular(e.User.ID))})
}
