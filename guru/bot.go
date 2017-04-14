package guru

import (
	"github.com/Sirupsen/logrus"
	"github.com/s1kx/stackguru/config"
	"github.com/s1kx/stackguru/guru/cmd"
	"github.com/s1kx/stackguru/guru/hook"
	"github.com/s1kx/unison"
)

func RunBot(conf *config.Config) {
	// Create bot structure
	settings := &unison.BotSettings{
		Token: conf.Discord.Token,

		Commands: []*unison.Command{
			cmd.EchoCommand,
		},
		EventHooks: []*unison.EventHook{
			hook.ChatlogHook,
		},
	}

	// Start the bot
	err := unison.RunBot(settings)
	if err != nil {
		logrus.Error(err)
	}
}
