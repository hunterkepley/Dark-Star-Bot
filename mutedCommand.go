package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func mutedCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Displays amount of currently Muted users
	currentGuild, err := getGuild(s, m)
	if err != nil {
		fmt.Println("Unabled to grab guild")
	} else {
		amount := 0
		mutedNames := []string{}
		tempRoleID := findRoleID("Muted", currentGuild)
		for i := 0; i < len(currentGuild.Members); i++ {
			for j := 0; j < len(currentGuild.Members[i].Roles); j++ {
				if currentGuild.Members[i].Roles[j] == tempRoleID {
					amount++
					mutedNames = append(mutedNames, currentGuild.Members[i].User.Username)
				}
			}
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%d muted users", amount),
			Description: strings.Join(mutedNames, "\n")})
	}
}
