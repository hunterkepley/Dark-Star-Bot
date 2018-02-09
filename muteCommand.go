package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	mutedUsers      []string // UserIDs
	mutedTimes      []int    //Minutes
	mutedStartTimes []time.Time
)

func muteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	currentMember, err := getMember(s, m)
	if err != nil {
		fmt.Println("Unabled to grab member, ")
		fmt.Println(err)
		return
	}
	currentGuild, err := getGuild(s, m)
	if err != nil {
		fmt.Println("Unabled to grab guild, ")
		fmt.Println(err)
		return
	}
	tempRoleIDM := findRoleID("Dark Mod", currentGuild) // Change these 3 to roles you want to be allowed to say @here/@everyone!
	tempRoleIDA := findRoleID("Dark Admins", currentGuild)
	tempRoleIDD := findRoleID("Dark Overlord", currentGuild)
	hasRoleM := memberHasRole(currentMember, tempRoleIDM)
	hasRoleA := memberHasRole(currentMember, tempRoleIDA)
	hasRoleD := memberHasRole(currentMember, tempRoleIDD)
	if hasRoleM || hasRoleA || hasRoleD {
		tm, err := strconv.Atoi(splitMsgLowered[2]) // 0 is mute, 1 is the user @'d, 2 is the minutes [should be]
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Please type minutes as the 3rd value! [Ex: `$mute @JohnDoe 30`]")
		} else {
			if len(splitMsgLowered) == 3 {
				minutesMuted := tm
				mutedID := m.Mentions[0].ID
				tempRoleID := findRoleID("Muted", currentGuild)
				err = s.GuildMemberRoleAdd(currentGuild.ID, mutedID, tempRoleID) // Give the muted role
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, "Unable to mute! Please make sure you have a Muted role, then make sure you have permissions in rooms you want them to be muted in set properly!")
					log.Fatal(err)
				}
				if err == nil { // Didnt want this popping up if there was an error
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Muted <@%s> for %s minutes.", mutedID, splitMsgLowered[2]))
					mutedUsers = append(mutedUsers, mutedID)
					mutedTimes = append(mutedTimes, minutesMuted)
					mutedStartTimes = append(mutedStartTimes, time.Now())
				}
			} else {
				s.ChannelMessageSend(m.ChannelID, "Please type the @'d user, then the minutes! [Ex: `$mute @JohnDoe 30`]")
			}
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "You aren't an admin or up so you can't mute!")
	}
}

func checkMutes(s *discordgo.Session, m *discordgo.MessageCreate) {
	for i := 0; i < len(mutedUsers); i++ {
		t := time.Now()
		elapsed := t.Sub(mutedStartTimes[i])
		if int(elapsed.Minutes()) >= mutedTimes[i] {
			currentGuild, err := getGuild(s, m)
			if err != nil {
				fmt.Println("Unabled to grab guild, ")
				fmt.Println(err)
				return
			}
			tempRoleID := findRoleID("Muted", currentGuild)
			err = s.GuildMemberRoleRemove(currentGuild.ID, mutedUsers[i], tempRoleID) // Remove the role
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unable to remove mute from <@%s>! Most likely the user has left or they don't have the role anymore!", mutedUsers[i]))
				log.Fatal(err)
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> You are now unmuted!", mutedUsers[i]))
				mutedUsers = append(mutedUsers[:i], mutedUsers[i+1:]...)
				mutedTimes = append(mutedTimes[:i], mutedTimes[i+1:]...)
				mutedStartTimes = append(mutedStartTimes[:i], mutedStartTimes[i+1:]...)
			}
		}
	}
}
