package cmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/s1kx/unison"
)

// EchoCommand responds with the message it receives
var EchoCommand = &unison.Command{
	Name:        "echo",
	Description: "what you say to me, i say to you.",
	Action:      echoAction,
}

func echoAction(ctx *unison.Context, m *discordgo.Message, content string) error {
	// Send same message back to user
	response := fmt.Sprintf("<@%s>: %s", m.Author.ID, content)

	_, err := ctx.Discord.ChannelMessageSend(m.ChannelID, response)

	return err
}
