# Dark-Star-Bot
Role management and server management discord bot made for the Dark Star Community, all in golang.

This bot was made for *Dark Star Gaming* so some of it has hardcoded roles, but are easily changed to fit any server.

Using the [discordgo library](https://www.github.com/bwmarrin/discordgo)

Use `go get github.com/bwmarrin/discordgo` to get it!

## To compile:

1) Download the Go compiler from http://golang.org

2) Go to the location of the source files

3) Type `go build` [I'd advise against `go run` because you have to use a token to run the bot correctly.]

4) Continue to running the bot.

## To run the bot for yourself:

1) Make a discord app

2) Make the discord app a bot

3) Take the token from the app and put it in a .bat or a .sh like this: Dark-Star-Bot.exe -t TOKENHERE

4) Use an OAuth2 Generator to get an invite link

5) MAKE SURE that you have a 'roles/' directory with the application, and that you have .json files! Check below for how to do.

6) Invite to your server and wammo!

## Authors

* **Hunter Kepley** - *Created bot* - [Github](https://www.github.com/hunterkepley)

# YOU NEED JSON FILES AND A MUTED ROLE FOR CERTAIN FUNCTIONALITY!

## How the $mute command works
You need to add a role called Muted

Also, make sure the rooms you want the people to be muted in are set in the permissions properly!

## How the .json files are located and done
The JSON files are of one config.json file, and of as many other JSON files you need, one for each server.

They MUST be placed in EXEDIRECTORY/roles/

* Check jsonParser.go for the structs, but a basic file structure is below

#### Example location:

~/Documents/Bot/Application

~/Documents/Bot/roles/config.json

~/Documents/Bot/roles/server1.json

#### Example config.json:

```
{
    "files" : [
        {
            "ID" : "123123123123123123",
            "location" : "roles/server1.json"
        },
        {
            "ID" : "321312312312312312",
            "location" : "roles/server2.json"
        }
    ]
}
```

#### Example server.json:

* Precursor notes:
1. Calls are what users input after `$role` to get the role
2. Role is the actual role given, has to be exact
3. Locked is if the user can *get* the role

```
{
    "serverID" : "123123123123123123",
    "roles" : [
        {
            "calls" : [
                "bottom",
                "bot"
            ],
            "role" : "Bottom",
            "locked" : true
        },
        {
            "calls" : [
                "top"
            ],
            "role" : "Top",
            "locked" : false
        }
    ],
    "welcomeMessage" : {
            "ID" : "333333333333333333",
            "message": "Hey %s! Welcome to the server! Go to <#333333333333333444> to set your roles!",
            "type" : "Welcome!"
    },
    "goodbyeMessage" : {
            "ID" : "222222222222222222",
            "message" : "Goodbye, %s!",
            "type" : "Later!"
    },
    "banMessage" : {
            "ID" : "222222222222222222",
            "message" : "We don't want you here, %s!",
            "type" : "Ban Hammer!"
    }
}
```
