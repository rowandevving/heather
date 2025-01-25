# Heather
A small discord bot I'm making with discordgo. If you have any cool suggestions be sure to make an issue so that I don't run out of ideas!

## Prerequisites
- A working Go environment
- A discord bot with the Message Content intent

## Setup
Clone this repository, and then create a configuration file. All you need to get started with a basic configuration is a command prefix and the location of the database the bot will use:
```json
{
  "databaseDir": "[PATH_TO_DATABASE_FOLDER]",
  "prefix": "[Prefix for commands e.g. !, ?, - etc]"
}
```
> [!NOTE]  
> A database will be created if one isn't found in the specified directory

Set the bot token of the discord bot you wish to use in the HEATHER_TOKEN environment variable (which can also be defined in a `.env` file), then execute `go run . -settings "[PATH_TO_YOUR_CONFIG_FILE]"` to start the bot and connect to discord.

## Tags
Tags are shortcodes that can be sent in a message for the bot to reply with a quick message. Useful for things like quicklinks or FAQs.

Tags can be called in a message with `--[tag name]-[sub-tag name]`

Tags can be configured using the `tags` array in the configuration file as so:

```json
"tags": [
  {
    "name": "hello",
    "message": "Bonjour", //message to send
    "subtags": [

      {
        "name": "world",
        "message": "Bonjour le monde"
      },
      {
        "name": "universe",
        "message": "Bonjour l'univers"
      }

    ]
  }
]
```
## Colour Roles
Role colours allow you to define custom colour roles your members can give themselves using the `color [color preset]` command.

Define colour presets in the `colors` array like this:

```json
"color": {

  "enabled": true,

  "colors": [
    {
      "name": "red",
      "hex": "ff0000"
    }
  ]
}
```
## Moderation

### Trusted Role

A Trusted Role is a role which is given to a member once they have sent a certain number of messages, making them "trusted" by the server.

This threshold and said role can be configured as so:

```json
"moderation": {

  ...
  "trustedRole": "[discord role name]",
  "trustedThreshold": 60,
  ...
}
```
