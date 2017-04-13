package core

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

// Bot is the definition of the chat bot.
type Bot struct {
	Commands []*Command
}

func (b Bot) Run(token string) error {
	// TODO: Validate commands

	// discordgo requires "Bot " prefix for Bot applications
	if !strings.HasPrefix(token, "Bot ") {
		token = "Bot " + token
	}

	// Initialize discord client
	discord, err := discordgo.New(token)
	if err != nil {
		return err
	}

	// Start bot
	session := &BotSession{
		bot: b,

		discord: discord,
	}
	session.Run()

	return nil
}

// BotSession is an active bot session.
type BotSession struct {
	bot     Bot
	discord *discordgo.Session

	readyState *discordgo.Ready
	user       *discordgo.User
}

func (bs *BotSession) Run() error {
	// Add Handlers
	{
		// Listen for ready state to receive bot information
		bs.discord.AddHandler(func(ds *discordgo.Session, r *discordgo.Ready) {
			bs.readyState = r
			bs.user = r.User

			logrus.WithField("ID", r.User.ID).Infof("I have finally found myself.")
		})

		// Create context for handlers
		ctx := &Context{
			BotSession: bs,
		}

		// Add handler for commands
		bs.discord.AddHandler(func(ds *discordgo.Session, m *discordgo.MessageCreate) {
			// Check if message is a command
			if !strings.HasPrefix(m.Content, CommandPrefix) {
				return
			}

			// Find command, if exists
			// TODO: Optimize, possibly with a map
			for _, cmd := range bs.bot.Commands {
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
		})
	}

	// Open the websocket and begin listening.
	err := bs.discord.Open()
	if err != nil {
		return fmt.Errorf("error opening connection: %s", err)
	}

	logrus.Info("Bot is now running.  Press CTRL-C to exit.")

	// Simple way to keep program running until CTRL-C is pressed.
	// TODO: Add signal handler to exit gracefully.
	<-make(chan struct{})

	return nil
}
