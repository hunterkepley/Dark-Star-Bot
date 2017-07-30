package main

import (
	"fmt"

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
	// Handles the ingame roles
	switch givenRole { // The switch statements with things like `ad` and `adc` is to
	case "top": //				   make it easier on people to add their roles.
		giveRole(s, m, "Top")
	case "adc", "ad":
		giveRole(s, m, "ADC")
	case "supp", "support":
		giveRole(s, m, "Support")
	case "mid":
		giveRole(s, m, "Mid")
	case "jungle", "jg":
		giveRole(s, m, "Jungle")
	case "fill":
		giveRole(s, m, "Fill")
		// Handles the ingame ranks
	case "unranked":
		giveRole(s, m, "Unranked")
	case "bronze":
		giveRole(s, m, "Bronze")
	case "silver":
		giveRole(s, m, "Silver")
	case "gold":
		giveRole(s, m, "Gold")
	case "platinum", "plat":
		giveRole(s, m, "Platinum")
	case "diamond":
		giveRole(s, m, "Diamond")
	case "master", "masters":
		giveRole(s, m, "Masters")
	case "challenger":
		giveRole(s, m, "Challenger")
	default:
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
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role added!", roleNeeded))
			}
		} else { // Bye bye role
			err := s.GuildMemberRoleRemove(currentGuild.ID, m.Author.ID, tempRoleID) // Remove the role
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Unable to remove role! Message <@!121105861539135490> and tell him there's a problem.")
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role removed!", roleNeeded))
			}
		}

	}

}
