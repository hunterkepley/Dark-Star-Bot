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
		"Why dont you slip into something more comfortable? (Like a coma.)",
		"If I were a bird, you'd be the first person I'd shit on.",
		"I'd slap you, but that'd be animal abuse.",
		"I want you to know that someone cares. [Not anyone here, but maybe someone.]",
		"You know what's funny? Not you. So shut the hell up.",
		"Someday you'll go far! I hope you will stay there.",
		"You're proof evolution can go in reverse",
		"You're fat.",
		"Fuck you lol",
		"I'd think of something better, but like your dick, life's too short.",
		"If you ran like your mouth, you'd be in good shape.",
		"Every kiss begins with 'K', too bad ugly begins with 'U'.",
		"kys.",
		"Is your ass jealous of all that shit coming out of your mouth?",
		"What the fuck did you just fucking say about me, you little bitch? I’ll have you know I graduated top of my class in the Navy Seals, and I’ve been involved in numerous secret raids on Al-Quaeda, and I have over 300 confirmed kills. I am trained in gorilla warfare and I’m the top sniper in the entire US armed forces. You are nothing to me but just another target. I will wipe you the fuck out with precision the likes of which has never been seen before on this Earth, mark my fucking words. You think you can get away with saying that shit to me over the Internet? Think again, fucker. As we speak I am contacting my secret network of spies across the USA and your IP is being traced right now so you better prepare for the storm, maggot. The storm that wipes out the pathetic little thing you call your life. You’re fucking dead, kid. I can be anywhere, anytime, and I can kill you in over seven hundred ways, and that’s just with my bare hands. Not only am I extensively trained in unarmed combat, but I have access to the entire arsenal of the United States Marine Corps and I will use it to its full extent to wipe your miserable ass off the face of the continent, you little shit. If only you could have known what unholy retribution your little “clever” comment was about to bring down upon you, maybe you would have held your fucking tongue. But you couldn’t, you didn’t, and now you’re paying the price, you goddamn idiot. I will shit fury all over you and you will drown in it. You’re fucking dead, kiddo.",
		"Don't take it personally, fam; You just strike me as a braindead conspircy theorist autist that has nothing intriguing going on in his life and has to speculate on hypotheticals to get a half-chub in the morning like a slackjawed drooling brainlet. I don't want to come off as abrasive, friend, So you have my apologies for rubbing you the wrong way, Just please-- In future, Try your best not to be a thoughtless mouthpiece for clickbait twitter accounts? It's very unbecomming of a woke ass nigga... I just don't have tollerence for mini-minds, tbh. This is the marketplace of ideas. If your product is unpaletable, I'm the soccer-mom nigga that always leaves a mean reveiw on yelp.",
		"As an outsider, what do you think of the human race?",
		"You're a real life 40 year old virgin",
		"You are a ferger.",
		"You're a deadshit.",
		"Bill O'Reilly is more bearable than your douchebaguette lookin ass.",
		"Dickass.",
		"Suck butter from my ass",
		"May you be struck by a dick",
		"May God give you to search for your children with a Geiger counter",
		"I don't like you particularly.")
}
