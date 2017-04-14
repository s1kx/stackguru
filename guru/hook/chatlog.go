package hook

import (
	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"

	"github.com/s1kx/unison"
	"github.com/s1kx/unison/events"
)

var ChatlogHook = &unison.EventHook{
	Name:        "chatlog",
	Description: "Logs messages in given channels to support reviewing them later.",
	Events: []events.EventType{
		events.MessageCreateEvent,
	},
	OnEvent: unison.EventHandlerFunc(chatlogAction),
}

func chatlogAction(ctx *unison.Context, dv *events.DiscordEvent) (bool, error) {
	var m *discordgo.Message

	// Check event type
	switch ev := dv.Event.(type) {
	default:
		return false, nil
	case *discordgo.MessageCreate:
		m = ev.Message
	}

	// Log message
	logrus.Infof("[chatlog] <%s>: %s", m.Author.Username, m.ContentWithMentionsReplaced())

	return true, nil
}
