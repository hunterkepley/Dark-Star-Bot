package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func roleCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Add role to someone
	if len(splitMsgLowered) > 1 { // If it just isnt `$role`

		// Handles the ingame roles
		switch splitMsgLowered[1] { // The switch statements with things like `ad` and `adc` is to
		case "top": //				   make it easier on people to add their roles.
			assignRole(s, m, "Top")
		case "adc":
		case "ad":
			assignRole(s, m, "ADC")
		case "support":
			assignRole(s, m, "Support")
		case "mid":
			assignRole(s, m, "Mid")
		case "jungle":
		case "jg":
			assignRole(s, m, "Jungle")
		case "fill":
			assignRole(s, m, "Fill")
			// Handles the ingame ranks
		case "unranked":
			assignRole(s, m, "Unranked")
		case "bronze":
			assignRole(s, m, "Bronze")
		case "gold":
			assignRole(s, m, "Gold")
		case "silver":
			assignRole(s, m, "Silver")
		case "platinum":
		case "plat":
			assignRole(s, m, "Platinum")
		case "diamond":
			assignRole(s, m, "Diamond")
		case "master":
		case "masters":
			assignRole(s, m, "Masters")
		case "challenger":
			assignRole(s, m, "Challenger")
		default:
			s.ChannelMessageSend(m.ChannelID, "You cannot add that role! Use `$Roles` to see all available roles")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please type a role name after `$role` !")
	}
}

func assignRole(s *discordgo.Session, m *discordgo.MessageCreate, roleNeeded string) { // Assigns the role based off of a role needed
	currentChannel, err := s.Channel(m.ChannelID) // Create the current channel object
	if err != nil {
		fmt.Println("Error getting channel", err)
	}
	currentGuild, err := s.Guild(currentChannel.GuildID) // Create the current guild object
	if err != nil {
		fmt.Println("Error getting guild", err)
	}

	state, err := s.State.Member(currentGuild.ID, m.Author.ID)
	if err != nil {
		fmt.Println("Error making state", err)
	}

	tempRoleID := "" // The temporary storage for role id.
	for i := 0; i < len(currentGuild.Roles); i++ {
		if currentGuild.Roles[i].Name == roleNeeded {
			tempRoleID = currentGuild.Roles[i].ID
		}
	}
	continueBool := true
	for i := 0; i < len(state.Roles); i++ {
		if state.Roles[i] == tempRoleID {
			continueBool = false
			break
		}
	}
	if continueBool {
		err := s.GuildMemberRoleAdd(currentGuild.ID, m.Author.ID, tempRoleID) // GuildID, userID, roleID .. returns an error
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to add role! Message <@!121105861539135490> and tell him there's a problem.")
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role added!", roleNeeded))
	} else {
		err := s.GuildMemberRoleRemove(currentGuild.ID, m.Author.ID, tempRoleID) // GuildMemberRoleRemove(guildID, userID, roleID) .. returns an error
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Unable to remove role! Message <@!121105861539135490> and tell him there's a problem.")
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role removed!", roleNeeded))
	}

}
