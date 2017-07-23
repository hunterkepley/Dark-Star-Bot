package main

import (
	"github.com/bwmarrin/discordgo"
)

func rolesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{ // Straight up prints all the roles you can have.. Nothing else to it
		Title:       "Roles Available:",
		Description: "Jungle, \nADC, \nTop, \nMid, \nSupport, \nUnranked, \nBronze, \nSilver, \nGold, \nPlatinum, \nDiamond, \nMasters, \nChallenger"})
}
