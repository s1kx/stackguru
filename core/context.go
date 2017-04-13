package core

// Context is a type that is passed to every handler
// in a bot application.
// It can be used to refer back to main components.
type Context struct {
	BotSession *BotSession
}
