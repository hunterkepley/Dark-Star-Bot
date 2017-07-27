package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func bugCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Reports a bug as a PM to me, replace the ID to your ID or even a channelID if you want
	if len(splitMsgLowered) > 1 {
		bugReport := strings.Join(splitMsgLowered[1:], " ")
		channel := createChannel(s, m, botOwnerID)

		authorChannel, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to message user who sent bug report. Please go on the github and make an `issue` about this. Use `$github` to see the github.")
		}
		s.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed{
			Title: "Bug Report:",
			Description: fmt.Sprintf("<@!%s>: %s",
				m.Author.ID,
				bugReport)})
		s.ChannelMessageSend(authorChannel.ID, "Sent bug report! Thanks for helping :smile:")
	} else {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "You need to write a message after `$bug`!"})
	}
}
