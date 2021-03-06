package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/s1kx/stackguru/config"
	"github.com/s1kx/stackguru/guru"
	"github.com/s1kx/stackguru/version"
)

const (
	defaultConfigPath = "~/.config/stackguru.toml"

	// EnvVarPrefix is the prefix for environment variables
	EnvVarPrefix = "STACKGURU"
)

var (
	// Version is the current package version
	Version string
)

// conf is a local package variable for access to the config from all cli commands
var conf config.Config

func main() {
	// Set package version for it to be accessible by subpackages
	version.PackageVersion = Version

	// Initialize command-line application
	app := &cli.App{
		Name:    "stackguru-go",
		Usage:   "discord nootropics bot",
		Version: version.PackageVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				EnvVars: envVarNames("CONFIG"),
				Value:   defaultConfigPath,
				Usage:   "path to configuration (toml format)",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				EnvVars: envVarNames("DEBUG"),
				Usage:   "debug mode",
			},
		},
		Before: initApplication,
		// Commands: []*cli.Command{
		// 	launchCmd,
		// },
		Action: runApplication,
	}
	app.Run(os.Args)
}

func envVarName(name string) string {
	return fmt.Sprintf("%s_%s", EnvVarPrefix, strings.ToUpper(name))
}

func envVarNames(names ...string) []string {
	res := make([]string, len(names))
	for i, name := range names {
		res[i] = envVarName(name)
	}
	return res
}

var logFormatter = logrus.TextFormatter{
	FullTimestamp:   true,
	TimestampFormat: "2006-01-02 15:04:05",
}

func initApplication(c *cli.Context) error {
	configPath := c.String("config")
	debug := c.Bool("debug")

	// Configure logger.
	logrus.SetFormatter(&logFormatter)
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Load configuration in to package variable conf.
	err := config.Load(configPath, &conf)
	if err != nil {
		logrus.Fatalf("Error loading configuration: %s", err)
		return err
	}

	return nil
}

func runApplication(c *cli.Context) error {
	guru.RunBot(&conf)

	return nil
}
