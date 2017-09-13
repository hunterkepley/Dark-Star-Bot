package main

import (
	"github.com/bwmarrin/discordgo"

	"strings"
)

var (
	commMap = make(map[string]Command)

	help   = Command{"help", "Displays all commands. Also can display specific information using `$help` and a command after, for example, `$help role`.", helpCommand}
	role   = Command{"role", "Lets you add either a lane role or a rank role to yourself. `$role jungle` for example.", roleCommand}
	roles  = Command{"roles", "Displays list of all roles available to add to yourself.", rolesCommand}
	bug    = Command{"bug", "Sends a bug report to the creator of Dark Star Bot.", bugCommand}
	github = Command{"github", "Displays a link to the github of the bot", githubCommand}
	vote   = Command{"vote", "Vote for either TSM or IMT, you can only pick one!", voteCommand}
	insult = Command{"insult", "Insult your (least) favorite person! Easily done by typing `$insult @NAMEHERE`", insultCommand}
)

// Command : Every command is made into a struct to make it simpler to work with and eliminate if statements
type Command struct {
	name        string
	description string
	exec        func(*discordgo.Session, *discordgo.MessageCreate)
}

func loadCommands() {
	commMap[help.name] = help
	commMap[role.name] = role
	commMap[roles.name] = roles
	commMap[bug.name] = bug
	commMap[github.name] = github
	commMap[vote.name] = vote
	commMap[insult.name] = insult
}

func parseCommand(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	if strings.Contains(string(command[0]), "$") { // If the first word of the message starts with `$`
		command = string(command[1:]) // Remove the `$` from the command
		if command == strings.ToLower(commMap[command].name) {
			commMap[command].exec(s, m)
		}
	}
	return
}
