package core

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func handleMessageCreate(ctx *Context, ds *discordgo.Session, m *discordgo.MessageCreate) {
	request := m.Content

	// Don't communicate with other bots
	if m.Author.Bot {
		return
	}

	// Don't communicate with myself
	if m.Author.ID == ctx.User.ID {
		return
	}

	// Check if message is a command
	if !strings.HasPrefix(request, CommandPrefix) {
		return
	}

	// Update the request
	request = cleanUpRequest(request, CommandPrefix)

	// Find command, if exists
	for title, _ := range ctx.Bot.Commands {
		if !strings.HasPrefix(request, title) {
			continue
		}

		request = cleanUpRequest(request, title)

		// Invoke command
		err := ctx.Bot.Commands[title].Action(ctx, ds, m.Message, request)
		if err != nil {
			logrus.Errorf("Command [%s]: %s", title, err)
		}

		// command was found, stop looping
		break
	}
}

// Removes a substring from the string and cleans up leading & trailing spaces.
func cleanUpRequest(str, remove string) string {
	result := strings.TrimPrefix(str, remove)
	return strings.TrimSpace(result)
}