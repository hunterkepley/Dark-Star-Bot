package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Help command... Yeah
	if len(splitMsgLowered) == 1 { // If it's just `[Help`, then do this
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "Commands:",
			Description: fmt.Sprintf("%s", helpMsg)})
	} else { // Otherwise, do this (like `[Help roleAddCommand`)
		if splitMsgLowered[1] == strings.ToLower(commMap[splitMsgLowered[1]].name) { // If the second word (looking for the command, like `[Help >roleAddCommand<`)
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       fmt.Sprintf("%s Help:", splitMsgLowered[1]),
				Description: commMap[splitMsgLowered[1]].description})
		} else { // If it doesn't match a command
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{Title: "Command does not exist!"})
		}
	}
}
