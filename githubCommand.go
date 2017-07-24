package main

import "github.com/bwmarrin/discordgo"

func githubCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Github",
		Description: "https://www.github.com/hunterkepley/dark-star-bot"})
}
