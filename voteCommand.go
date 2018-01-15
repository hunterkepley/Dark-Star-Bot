package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/* RARELY IN USE SO IT'S REALLY BADLY MADE FOR NOW */

func voteCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Add vote role to someone
	if len(splitMsgLowered) > 1 { // If it just isnt `$vote`
		assignVote(s, m, splitMsgLowered[1])
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please type `Faker`, `Bjergsen`, `VVVERT`, `FoFo`, `PowerOfEvil`, `Jisu`, `Meiko`, `Frozen`, `Uzi`, `Levi`, `brTT`, `Sneaky`, `PraY`, `SwordArT`, `Rekkles`, or `Zeitnot` after `$vote` !")
	}
}

func assignVote(s *discordgo.Session, m *discordgo.MessageCreate, givenRole string) {
	// Handles the vote roles
	switch givenRole { // The switch statements with things like `tsm` and `solomid
	case "faker", "Faker": //				   make it easier on people to add their roles.
		giveVoteRole(s, m, "Faker")
	case "bjergsen", "Bjergsen":
		giveVoteRole(s, m, "Bjergsen")
	case "vvvert", "VVVERT", "Vvvert":
		giveVoteRole(s, m, "VVVERT")
	case "fofo", "FOFO", "FoFo", "Fofo":
		giveVoteRole(s, m, "FoFo")
	case "powerofevil", "PowerofEvil", "PowerOfEvil", "Powerofevil":
		giveVoteRole(s, m, "PowerOfEvil")
	case "jisu", "Jisu", "JISU":
		giveVoteRole(s, m, "Jisu")
	case "Meiko", "meiko":
		giveVoteRole(s, m, "Meiko")
	case "frozen", "Frozen", "FROZEN":
		giveVoteRole(s, m, "Frozen")
	case "uzi", "Uzi", "UZI":
		giveVoteRole(s, m, "Uzi")
	case "levi", "Levi":
		giveVoteRole(s, m, "Levi")
	case "brTT", "BRTT", "brtt":
		giveVoteRole(s, m, "brTT")
	case "Sneaky", "sneaky", "sneakers":
		giveVoteRole(s, m, "Sneaky")
	case "PraY", "pray", "PRAY", "Pray":
		giveVoteRole(s, m, "PraY")
	case "SwordArT", "swordart", "Swordart", "SwordArt":
		giveVoteRole(s, m, "SwordArT")
	case "Rekkles", "rekkles":
		giveVoteRole(s, m, "Rekkles")
	case "zeitnot", "Zeitnot", "ZeitNot":
		giveVoteRole(s, m, "Zeitnot")
	default:
		s.ChannelMessageSend(m.ChannelID, "Invalid vote! The players are `Faker`, `Bjergsen`, `VVVERT`, `FoFo`, `PowerOfEvil`, `Jisu`, `Meiko`, `Frozen`, `Uzi`, `Levi`, `brTT`, `Sneaky`, `PraY`, `SwordArT`, `Rekkles`, or `Zeitnot`!")
	}
}

func giveVoteRole(s *discordgo.Session, m *discordgo.MessageCreate, roleNeeded string) { // Assigns the role based off of a role needed

	currentGuild := getGuild(s, m)
	currentMember := getMember(s, m)

	if currentGuild != nil && currentMember != nil {

		fakerRoleID := findRoleID("Faker", currentGuild)
		bjergsenRoleID := findRoleID("Bjergsen", currentGuild)
		vvvertRoleID := findRoleID("VVVERT", currentGuild)
		fofoRoleID := findRoleID("FoFo", currentGuild)
		powerofevilRoleID := findRoleID("PowerOFEvil", currentGuild)
		jisuRoleID := findRoleID("Jisu", currentGuild)
		meikoRoleID := findRoleID("Meiko", currentGuild)
		frozenRoleID := findRoleID("Frozen", currentGuild)
		uziRoleID := findRoleID("Uzi", currentGuild)
		leviRoleID := findRoleID("Levi", currentGuild)
		brttRoleID := findRoleID("brTT", currentGuild)
		sneakyRoleID := findRoleID("Sneaky", currentGuild)
		prayRoleID := findRoleID("PraY", currentGuild)
		swordartRoleID := findRoleID("SwordArT", currentGuild)
		rekklesRoleID := findRoleID("Rekkles", currentGuild)
		zeitnotRoleID := findRoleID("Zeitnot", currentGuild)
		tempRoleID := findRoleID(roleNeeded, currentGuild)

		hasfakerRole := memberHasRole(currentMember, fakerRoleID)
		hasbjergsenRole := memberHasRole(currentMember, bjergsenRoleID)
		hasvvvertRole := memberHasRole(currentMember, vvvertRoleID)
		hasfofoRole := memberHasRole(currentMember, fofoRoleID)
		haspowerofevilRole := memberHasRole(currentMember, powerofevilRoleID)
		hasjisuRole := memberHasRole(currentMember, jisuRoleID)
		hasmeikoRole := memberHasRole(currentMember, meikoRoleID)
		hasfrozenRole := memberHasRole(currentMember, frozenRoleID)
		hasuziRole := memberHasRole(currentMember, uziRoleID)
		hasleviRole := memberHasRole(currentMember, leviRoleID)
		hasbrttRole := memberHasRole(currentMember, brttRoleID)
		hassneakyRole := memberHasRole(currentMember, sneakyRoleID)
		hasprayRole := memberHasRole(currentMember, prayRoleID)
		hasswordartRole := memberHasRole(currentMember, swordartRoleID)
		hasrekklesRole := memberHasRole(currentMember, rekklesRoleID)
		haszeitnotRole := memberHasRole(currentMember, zeitnotRoleID)

		if !hasfakerRole && !hasbjergsenRole && !hasvvvertRole &&
			!hasfofoRole && !haspowerofevilRole && !hasjisuRole &&
			!hasmeikoRole && !hasfrozenRole &&
			!hasuziRole && !hasleviRole && !hasbrttRole &&
			!hassneakyRole && !hasprayRole && !hasswordartRole &&
			!hasrekklesRole && !haszeitnotRole { // Give that guy a role
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
