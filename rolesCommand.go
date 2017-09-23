package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func rolesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var roles []string
	var serverID string
	_ = serverID // Added because it thinks it isn't being used.
	dsrFiles := getFilesFromDir("roles/*.dsr")
	for i := 0; i < len(dsrFiles); i++ {
		_, troles, serverID := handledsr(dsrFiles[i])
		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		} else {
			guildID := channel.GuildID
			if serverID == guildID {
				roles = troles
			}
		}
	}
	ps := ""
	if len(roles) > 0 {
		for i := 0; i < len(roles); i++ {
			ps += fmt.Sprintf("%s\n", roles[i])
		}
	} else {
		ps = "No roles available!"
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{ // Straight up prints all the roles you can have.. Nothing else to it
		Title:       "Roles Available:",
		Description: ps})
}
