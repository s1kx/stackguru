package hook

import (
	"fmt"
	//"errors"

	"github.com/bwmarrin/discordgo"

	"bitbucket.org/stackguru/stackguru-go/core"
	"bitbucket.org/stackguru/stackguru-go/core/events"
)

var ChatlogHook = &core.Hook{
	Name:        "chatlog",
	Description: "Logs messages in given channels to support reviewing them later.",
	Events:		 []string{
		events.MessageCreate,
	},
	Action:      chatlogAction,
}

func chatlogAction(ctx *core.Context, event string, ds *discordgo.Session, m *discordgo.Message) error {
	// Do something interesting with the m object

	// show in the terminal that this runs correctly
	fmt.Println("chatlog was run")

	return nil
}