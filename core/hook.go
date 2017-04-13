package core

// Hook interface for anything that is supposed to react on a event, besides commands.
type Hook struct {
	// Name of the hook
	Name string

	// Description of what the hook does
	Description string
	
	// Events for the hook should react to
	Events []string

	// Check if this hook is deactivated
	Deactivated bool

	// Command behavior
	Action HookActionFunc
}

