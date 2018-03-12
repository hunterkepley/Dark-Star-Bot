package main

import (
	"github.com/bwmarrin/discordgo"

	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Variables used for command line parameters.
var (
	Token string
)

// Custom variables
var (
	helpMsg = "Prefix: `$`\nHelp\nRole\nRoles\nBug\nGithub\nMute\nMuted\nMembers"

	splitMsgLowered = []string{}

	botOwnerID = "" // Change to your id on discord
)

func makeSplitMessage(s *discordgo.Session, m *discordgo.MessageCreate) []string {
	// The message, split up
	splitMessage := strings.Fields(strings.ToLower(m.Content))

	return splitMessage
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord sessions using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(messageCreate)
	// Register the guildMemberAddHandler func as a callback for GuildMemberAdd events
	dg.AddHandler(guildMemberAddHandler)
	// Register the guildMemberRemoveHandler func as a callback for GuildMemberRemove events
	dg.AddHandler(guildMemberRemoveHandler)
	// Register the guildMemberBannedHandler func as a callback for GuildBanAdd events
	dg.AddHandler(guildMemberBannedHandler)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	loadCommands()

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("The bot is now running. Press CTRL-C to stop")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session
	defer dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) { // Message handling
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	splitMsgLowered = makeSplitMessage(s, m)

	checkHereEveryone(s, m)

	if len(splitMsgLowered) > 0 { // Prevented a really rare and weird bug about going out of index.
		parseCommand(s, m, splitMsgLowered[0]) // Really shouldnt happen since `MessageCreate` is about
	} // 										messages made on create...
}

func checkHereEveryone(s *discordgo.Session, m *discordgo.MessageCreate) { // Doesn't let members who aren't mods, admins, or owners @here or @everyone. Please change to your needs.
	if strings.Contains(m.Content, "@here") || strings.Contains(m.Content, "@everyone") { // [Yes, discord has functionality for this but I was instructed to add this for whatever stupid reason]
		currentGuild, err := getGuild(s, m)
		if err != nil {
			fmt.Println("Unabled to grab guild, ")
			fmt.Println(err)
			return
		}
		currentMember, err := getMember(s, m)
		if err != nil {
			fmt.Println("Unabled to grab member, ")
			fmt.Println(err)
			return
		}
		rolesNeeded := []string{"Dark Mod", "Dark Admins", "Dark Overlord", "Staff"}
		canHere := false
		for i := 0; i < len(rolesNeeded); i++ { // Checks if the muter has roles above
			tempRoleID := findRoleID(rolesNeeded[i], currentGuild)
			hasRoleID := memberHasRole(currentMember, tempRoleID)
			if hasRoleID {
				canHere = true
			}
		}
		if !canHere {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s>, please don't use `@here` or `@everyone`, if the message is important, talk to an admin or mod", m.Author.ID))
			s.ChannelMessageDelete(m.ChannelID, m.ID)
		}
	}
}

func guildMemberAddHandler(s *discordgo.Session, e *discordgo.GuildMemberAdd) { // Handles GuildMemberAdd'ing
	if e.User.Bot {
		return
	}

	welcomeMessage(s, e)
}

func guildMemberRemoveHandler(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	if e.User.Bot {
		return
	}

	goodbyeMessage(s, e)
}

func guildMemberBannedHandler(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	if e.User.Bot {
		return
	}

	banMessage(s, e)
}

func getGuild(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Guild, error) { // Returns guild
	currentChannel, err := getChannel(s, m)
	currentGuild, err := s.Guild(currentChannel.GuildID) // Create the current guild object

	return currentGuild, err
}

func getChannel(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Channel, error) { // Returns channel
	currentChannel, err := s.Channel(m.ChannelID) // Create the current channel object

	return currentChannel, err
}

func getMember(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Member, error) { // Returns member
	currentGuild, err := getGuild(s, m)
	member, err := s.State.Member(currentGuild.ID, m.Author.ID)

	return member, err
}

func getState(s *discordgo.Session) *discordgo.State { // Returns state
	state := s.State

	return state
}

func findRoleID(roleNeeded string, currentGuild *discordgo.Guild) string { // Returns a role ID from a list of roles based off of a string
	rID := ""
	for i := 0; i < len(currentGuild.Roles); i++ {
		if currentGuild.Roles[i].Name == roleNeeded {
			rID = currentGuild.Roles[i].ID
		}
	}

	return rID
}

func memberHasRole(currentMember *discordgo.Member, tempRoleID string) bool { // Returns true if the member has a role, otherwise, false
	for i := 0; i < len(currentMember.Roles); i++ { // If the user doesn't have that role, add it, otherwise, remove it
		if currentMember.Roles[i] == tempRoleID {
			return true
		}
	}

	return false
}

func createChannel(s *discordgo.Session, m *discordgo.MessageCreate, ID string) *discordgo.Channel {
	channel, err := s.UserChannelCreate(ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Unable to create new channel.")
	}

	return channel
}
