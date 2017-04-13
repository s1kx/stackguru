package core

import "github.com/bwmarrin/discordgo"

// CommandActionFunc is the action to execute when a command is called.
type CommandActionFunc func(ctx *Context, ds *discordgo.Session, m *discordgo.Message, content string) error
