package hook

import (
	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"

	"bitbucket.org/stackguru/stackguru-go/core"
)

var ChatlogHook = &core.EventHook{
	Name:        "chatlog",
	Description: "Logs messages in given channels to support reviewing them later.",
	// Events: []string{
	// 	events.MessageCreate,
	// },
	OnEvent: chatlogAction,
}

func chatlogAction(ctx *core.Context, ds *discordgo.Session, event interface{}) (bool, error) {
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
