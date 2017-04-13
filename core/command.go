package core

const CommandPrefix = "!"

// Command is the interface that every bot command must implement.
type Command struct {
	// Name of the command
	Name string
	
	// Aliases for the command
	Aliases []string

	// Description of what the command does
	Description string

	// Check if this command is deactivated
	Deactivated bool

	Action CommandActionFunc
}

