package hook

import (
	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"

	"github.com/s1kx/unison"
)

var ChatlogHook = &unison.EventHook{
	Name:        "chatlog",
	Description: "Logs messages in given channels to support reviewing them later.",
	// Events: []string{
	// 	events.MessageCreate,
	// },
	OnEvent: unison.EventHandlerFunc(chatlogAction),
}

func chatlogAction(ctx *unison.Context, ds *discordgo.Session, event interface{}) (bool, error) {
	var m *discordgo.Message

	// Check event type
	switch ev := event.(type) {
	default:
		return false, nil
	case *discordgo.MessageCreate:
		m = ev.Message
	}

	// Log message
	logrus.Infof("[chatlog] <%s>: %s", m.Author.Username, m.ContentWithMentionsReplaced())

	return true, nil
}
