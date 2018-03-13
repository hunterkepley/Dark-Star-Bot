package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func membersCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Displays # of members in a pretty embed
	currentGuild, err := getGuild(s, m)
	if err != nil {
		fmt.Println("Unable to get guild")
		log.Fatal(err)
	} else {
		amount := 0
		bots := 0
		for i := 0; i < len(currentGuild.Members); i++ {
			if !currentGuild.Members[i].User.Bot {
				amount++
			} else {
				bots++
			}
		}
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s", currentGuild.Name),
			Thumbnail:   &discordgo.MessageEmbedThumbnail{discordgo.EndpointGuildIcon(currentGuild.ID, currentGuild.Icon), "", 30, 30},
			Description: fmt.Sprintf("There are `%d` member[s] and `%d` bot[s]", amount, bots)})
	}
}
