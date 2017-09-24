package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func rolesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var config DSGConfig
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Fatal(err)
	} else {
		config, err := getConfigForGuildId(channel.GuildID)
		_ = config
		if err != nil {
			log.Fatal(err)
		}
	}

	ps := ""
	if len(config.roles) > 0 {
		for i := 0; i < len(config.roles); i++ {
			ps += fmt.Sprintf("%s\n", config.roles[i])
		}
	} else {
		ps = "No roles available!"
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{ // Straight up prints all the roles you can have.. Nothing else to it
		Title:       "Roles Available:",
		Description: ps})
}
