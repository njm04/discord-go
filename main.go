// created by: Neil John Medalla
package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strings"
	"disord-flames/flames"
	)

var (
	commandPrefix string
	botID string
	test []string
)

func main () {
	discord, err := discordgo.New("Bot NDcxOTM1OTU0ODM2NTIwOTkw.DjsR9w.IMJeOG3VaAltWxkpMZrntfiQa5U")
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)
	botID = user.ID

	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "A friendly helpful bot")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("Envoy has started on %d servers\n", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})
}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if len(message.Message.Content) == 0 {
		return
	}
	argsArray := strings.Split(message.Message.Content, " ")
	if len(argsArray) == 4 && argsArray[0][1:] == "match" {
		names := argsArray[1:]
		letters := flames.FindUniqueLettersOfThreeNames(names)
		fmt.Println(letters)
		matchedAs := flames.Flames(letters, names)
		embed := &discordgo.MessageEmbed{
			Description: matchedAs,
		}

		user := message.Author
		if user.ID == discord.State.User.ID || user.Bot {
			return
		}
		if strings.HasPrefix(argsArray[0], commandPrefix) {
			command := argsArray[0][1:]
			if command == "match" {
				discord.ChannelMessageSendEmbed(message.ChannelID, embed)
			} else {
				discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
					Description: "I don't know that command!",
				})
			}
		}
	}

	if len(argsArray) == 3 && argsArray[0][1:] == "match"{
		names := argsArray[1:]
		letters := flames.FindUniqueLetters(names)
		fmt.Println(letters)
		matchedAs := flames.Flames(letters, names)
		embed := &discordgo.MessageEmbed{
			Description: matchedAs,
		}

		user := message.Author
		if user.ID == discord.State.User.ID || user.Bot {
			return
		}
		if strings.HasPrefix(argsArray[0], commandPrefix) {
			command := argsArray[0][1:]
			if command == "match" {
				discord.ChannelMessageSendEmbed(message.ChannelID, embed)
			} else {
				discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
					Description: "I don't know that command!",
				})
			}
		}
	}
}