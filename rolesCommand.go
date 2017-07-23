package main

import (
	"github.com/bwmarrin/discordgo"
)

func rolesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Roles Available:",
		Description: "Jungle, \nADC, \nTop, \nMid, \nSupport, \nUnranked, \nBronze, \nSilver, \nGold, \nPlatinum, \nDiamond, \nMasters, \nChallenger"})
}
