package cmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bitbucket.org/stackguru/stackguru-go/core"
)

// EchoCommand responds with the message it receives
var EchoCommand = &core.Command{
	Name:        "echo",
	Description: "what you say to me, i say to you.",
	Action:      echoAction,
}

func echoAction(ctx *core.Context, ds *discordgo.Session, m *discordgo.Message, content string) error {
	// Send same message back to user
	response := fmt.Sprintf("<@%s>: %s", m.Author.ID, content)

	_, err := ds.ChannelMessageSend(m.ChannelID, response)

	return err
}
