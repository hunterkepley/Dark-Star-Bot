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
	helpMsg = fmt.Sprintf("Prefix: `$`\nHelp\nRole\nRoles")

	splitMsgLowered = []string{}
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

	if len(splitMsgLowered) > 0 { // Prevented a really rare and weird bug about going out of index.
		parseCommand(s, m, splitMsgLowered[0]) // Really shouldnt happen since `MessageCreate` is about
	} // 										messages made on create...
}
