package config

// Config contains all configuration objects.
type Config struct {
	Application *ApplicationConfig
	Discord     *DiscordConfig
}

// ApplicationConfig is the general application configuration.
type ApplicationConfig struct {
}

// DiscordConfig is the discord configuration.
type DiscordConfig struct {
	Token string `toml:"token"`
}
