package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func muteCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // This needs to be redone at some point, so damn messy
	currentMember, err := getMember(s, m)
	if err != nil {
		fmt.Println("Unabled to grab member, ")
		log.Fatal(err)
	}
	currentGuild, err := getGuild(s, m)
	if err != nil {
		fmt.Println("Unabled to grab guild")
		log.Fatal(err)
	}
	rolesNeeded := []string{"Dark Mod", "Dark Admins", "Dark Overlord", "Community Leader"}
	canMute := false
	for i := 0; i < len(rolesNeeded); i++ { // Checks if the muter has roles above
		tempRoleID := findRoleID(rolesNeeded[i], currentGuild)
		hasRoleID := memberHasRole(currentMember, tempRoleID)
		if hasRoleID {
			canMute = true
		}
	}
	if canMute {
		if len(splitMsgLowered) == 2 {
			if len(m.Mentions) > 0 {
				mutedID := m.Mentions[0].ID
				tempRoleID := findRoleID("Muted", currentGuild)
				mutedMember, _ := s.GuildMember(currentGuild.ID, mutedID)
				hasRole := memberHasRole(mutedMember, tempRoleID)
				if !hasRole {
					err = s.GuildMemberRoleAdd(currentGuild.ID, mutedID, tempRoleID) // Give the muted role
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "Unable to mute! Please make sure you have a Muted role, then make sure you have permissions in rooms you want them to be muted in set properly!")
						log.Fatal(err)
					}
					if err == nil { // Didnt want this popping up if there was an error
						s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Muted <@%s>.", mutedID))
					}
				} else {
					err := s.GuildMemberRoleRemove(currentGuild.ID, mutedID, tempRoleID) // Remove the muted role
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "Unable to unmute! Please make sure you have a Muted role, then make sure you have permissions in rooms you want them to be muted in set properly!")
						log.Fatal(err)
					}
					if err == nil { // Didnt want this popping up if there was an error
						s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unmuted <@%s>!", mutedID))
					}
				}
			} else {
				s.ChannelMessageSend(m.ChannelID, "Please type the @'d user [Ex: `$mute @JohnDoe#1234`]")
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "The format is, `$mute @JohnDoe#1234`")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "You cant mute anyone!")
	}
}
