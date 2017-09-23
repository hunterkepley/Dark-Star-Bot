package main

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var (
	insults []string
)

func insultCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Insults someone
	if len(splitMsgLowered) > 1 { // If it just isnt `$insult`

		for i := 0; i < len(splitMsgLowered)-1; i++ {
			insultUser(s, m, splitMsgLowered[i+1])
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please type a name after `$insult` !")
	}
}

func insultUser(s *discordgo.Session, m *discordgo.MessageCreate, usr string) { // Get insulted, boyo
	ri := insults[rand.Intn(len(insults))]
	if len(m.Mentions) > 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@!%s>, %s", m.Mentions[0].ID, ri))
	} else {
		s.ChannelMessageSend(m.ChannelID, "You did not specify some person to slam with a hot insult.")
	}
}

func loadInsults() {
	insults = append(insults,
		"If you were any less intelligent we'd have to water you three times a week.",
		"I'm not saying I hate you, but I would unplug your life support to charge my phone.",
		"I would insult you but nature did a better job.",
		"You're the reason the gene pool needs a lifeguard.",
		"If I ate a bowl of alphabet soup, I could shit out a smarter sentence than any of yours.",
		"You're not pretty enough to be this stupid.",
		"2090 called. You're dead and you wasted your time on earth.",
		"You're as funny as Amy Schumer.",
		"Why dont you slip into something more comfortable? (Like a coma.)",
		"I'd slap you, but that'd be animal abuse.",
		"I want you to know that someone cares. [Not anyone here, but maybe someone.]",
		"You know what's funny? Not you. So shut the hell up.",
		"Someday you'll go far! I hope you will stay there.",
		"You're proof evolution can go in reverse",
		"You're fat.",
		"If you ran like your mouth, you'd be in good shape.",
		"Every kiss begins with 'K', too bad ugly begins with 'U'.",
		"As an outsider, what do you think of the human race?",
		"You're a real life 40 year old virgin",
		"You are a ferger.",
		"I don't like you particularly.")
}
