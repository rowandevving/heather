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
  "databaseDir": "[PATH_TO_DATABASE_FOLDER]" //Heather will create a database if one doesn't already exist
}
```
Then, execute `go run . -settings "[PATH_TO_YOUR_CONFIG_FILE]"` to connect the bot up and create a websocket connection.
