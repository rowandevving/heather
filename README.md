# Heather
A small discord bot I'm making with discordgo. If you have any cool suggestions be sure to make an issue so that I don't run out of ideas!

## Prerequisites
- A working Go environment
- A discord bot with the Message Content intent

## Setup
Clone this repository, and then create a configuration file. All you need to get started with a basic configuration is your bot's token:
```json
{
  "token": "[YOUR_PRIVATE_BOT_TOKEN_HERE]"
  "databaseDir": "[PATH_TO_DATABASE_FOLDER]"
  "prefix": "[Prefix for commands e.g. !, ?, - etc]"
}
```
> [!NOTE]  
> A database will be created if one isn't found in the specified directory

Then, execute `go run . -settings "[PATH_TO_YOUR_CONFIG_FILE]"` to connect the bot up and create a websocket connection.
