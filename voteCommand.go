package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func voteCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Add vote role to someone
	if len(splitMsgLowered) > 1 { // If it just isnt `$vote`
		assignVote(s, m, splitMsgLowered[1])
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please type `TSM` or `IMT` name after `$vote` !")
	}
}

func assignVote(s *discordgo.Session, m *discordgo.MessageCreate, givenRole string) {
	// Handles the vote roles
	switch givenRole { // The switch statements with things like `tsm` and `solomid
	case "tsm", "solomid", "teamsolomid": //				   make it easier on people to add their roles.
		giveVoteRole(s, m, "Team SoloMid")
	case "imt", "immortals", "immortal", "imortals", "imortal":
		giveVoteRole(s, m, "Immortals")
	default:
		s.ChannelMessageSend(m.ChannelID, "Invalid vote! The teams are `TSM` and `IMT` [`Team SoloMid` and `Immortals`]")
	}
}

func giveVoteRole(s *discordgo.Session, m *discordgo.MessageCreate, roleNeeded string) { // Assigns the role based off of a role needed

	currentGuild := getGuild(s, m)
	currentMember := getMember(s, m)

	if currentGuild != nil && currentMember != nil {

		imtRoleID := "" // The temporary storage for role id.
		tsmRoleID := ""
		tempRoleID := ""

		imtRoleID = findRoleID("Immortals", currentGuild)
		tsmRoleID = findRoleID("Team SoloMid", currentGuild)
		tempRoleID = findRoleID(roleNeeded, currentGuild)

		hasimtRole := memberHasRole(currentMember, imtRoleID)
		hastsmRole := memberHasRole(currentMember, tsmRoleID)

		if !hasimtRole && !hastsmRole { // Give that guy a role
			err := s.GuildMemberRoleAdd(currentGuild.ID, m.Author.ID, tempRoleID) // Give the role
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Unable to assign vote! Message <@!121105861539135490> and tell him there's a problem.")
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Voted for %s!", roleNeeded))
			}
		} else { // Already voted
			s.ChannelMessageSend(m.ChannelID, "You've already voted!")
		}

	}

}
