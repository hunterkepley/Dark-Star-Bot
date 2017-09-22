# Dark-Star-Bot
Role management and server management discord bot made for the Dark Star Community, all in golang.

Using the discordgo library [https://www.github.com/bwmarrin/discordgo]

# To compile:

1) Download the Go compiler from http://golang.org
2) Go to the location of the source files
3) Type `go build` [I'd advise against `go run` because you have to use a token to run the bot correctly.]
4) Continue to running the bot.

# To run the bot for yourself:

1) Make a discord app
2) Make the discord app a bot
3) Take the token from the app and put it in a .bat or a .sh like this: Dark-Star-Bot.exe -t TOKENHERE
4) Use an OAuth2 Generator to get an invite link
4.5) MAKE SURE that you have a 'roles/' directory with the application, and that you have a .dsr file! That is how the bot knows which roles you can call from the bot. See below to see how to make a .dsr file properly.
5) Invite to your server and wammo!

# How to make a .dsr file:

## It's structured like so:

GuildID;
call1,call2=Role;

## The calls MUST NOT have spaces, the Roles CAN have spaces

exampleone=Example One;

## A call is what the user will type to request a role,

Example:
### IN DSR FILE:

GuildID;
role1,roleone=Role1;

### USER CAN TYPE THESE TO CALL THE ROLE:

'$role role1'
'$role roleone'

### Make SURE that the .dsr files are in a directory called 'roles' and that directory is in the same directory as the application.

Example;

Folder1/Application
Folder1/roles/server1roles.dsr
