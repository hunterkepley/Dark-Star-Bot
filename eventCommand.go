package main

import (
	"github.com/bwmarrin/discordgo"
)

var (
	eventActive = false
	currentEvent = ""
)

func eventCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if eventActive == false {
		currentGuild := getGuild(s, m) // Makes it only for mods/admins/overlords, etc
		currentMember := getMember(s, m)
		tempRoleIDM := findRoleID("Dark Mod", currentGuild)
		tempRoleIDA := findRoleID("Dark Admins", currentGuild)
		tempRoleIDD := findRoleID("Dark Overlord", currentGuild)
		hasRoleM := memberHasRole(currentMember, tempRoleIDM)
		hasRoleA := memberHasRole(currentMember, tempRoleIDA)
		hasRoleD := memberHasRole(currentMember, tempRoleIDD)
		if !hasRoleM && !hasRoleA && !hasRoleD {
			s.ChannelMessageSend(m.ChannelID, "You can't start an event.")
		} else {
			startEvent(s, m)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "There is already an event active, please type `$endevent` to end the current event if you wish to make a new one.")
	}
}

func endeventCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	endEvent(s, m)
	s.ChannelMessageSend(m.ChannelID, "Event successfully ended!")
}

func signupCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func startEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	a := ""
	for i := 0; i < len(splitMsgLowered); i++ {
		if i > 0 {
			a += (splitMsgLowered[i] + " ") // Creates a whole string from the event.
		}
	}
	currentEvent = a
	s.ChannelMessageSend(m.ChannelID, currentEvent)
	eventActive = true
}

func endEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	currentEvent = ""
	eventActive = false
}