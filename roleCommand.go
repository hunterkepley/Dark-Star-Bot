package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func roleCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Add role to someone
	if len(splitMsgLowered) > 1 { // If it just isnt `$role`

		for i := 0; i < len(splitMsgLowered)-1; i++ {
			assignRole(s, m, splitMsgLowered[i+1])
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please type a role name after `$role` !")
	}
}

func assignRole(s *discordgo.Session, m *discordgo.MessageCreate, givenRole string) {
	var config dsgConfig
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Fatal(err)
	}
	config, err = getConfigForGuildID(channel.GuildID)
	_ = config
	if err != nil {
		log.Fatal(err)
	}

	roleUsed := false

	// Handles the ingame roles
	for i := 0; i < len(config.calls); i++ {
		for j := 0; j < len(config.calls[i]); j++ {
			switch givenRole {

			case config.calls[i][j]:
				roleUsed = true
				giveRole(s, m, config.roles[i])
			}
		}
	}

	if !roleUsed {
		s.ChannelMessageSend(m.ChannelID, "You cannot add that role! Use `$Roles` to see all available roles")
	}
}

func giveRole(s *discordgo.Session, m *discordgo.MessageCreate, roleNeeded string) { // Assigns the role based off of a role needed

	currentGuild := getGuild(s, m)
	currentMember := getMember(s, m)

	if currentGuild != nil && currentMember != nil {

		tempRoleID := "" // The temporary storage for role id.

		tempRoleID = findRoleID(roleNeeded, currentGuild)

		hasRole := memberHasRole(currentMember, tempRoleID)

		if !hasRole { // Give that guy a role
			err := s.GuildMemberRoleAdd(currentGuild.ID, m.Author.ID, tempRoleID) // Give the role
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Unable to add role! Message <@!121105861539135490> and tell him there's a problem.")
				log.Fatal(err)
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role added!", roleNeeded))
			}
		} else { // Bye bye role
			err := s.GuildMemberRoleRemove(currentGuild.ID, m.Author.ID, tempRoleID) // Remove the role
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Unable to remove role! Message <@!121105861539135490> and tell him there's a problem.")
				log.Fatal(err)
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role removed!", roleNeeded))
			}
		}

	}

}

func getFilesFromDir(s string) []string { // Gets all files from a directory
	r, err := filepath.Glob(s)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
