package core

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func handleMessageCreate(ctx *Context, ds *discordgo.Session, m *discordgo.MessageCreate) {
	// Avoid talking to myself and other bots
	if m.Author.Bot || m.Author.ID == ctx.User.ID {
		return
	}
	// Check if message is a command
	if !strings.HasPrefix(m.Content, CommandPrefix) {
		return
	}

	// Find command, if exists
	// TODO: Optimize, possibly with a map
	for _, cmd := range ctx.Bot.Commands {
		prefix := CommandPrefix + cmd.Name
		if !strings.HasPrefix(m.Content, prefix) {
			continue
		}

		// Trim command from content
		content := strings.TrimPrefix(m.Content, prefix)

		// Invoke command
		err := cmd.Action(ctx, ds, m.Message, content)
		if err != nil {
			logrus.Errorf("Command [%s]: %s", cmd.Name, err)
		}
	}
}
