/*
Package parses the config (global and user) and provides to the application.
It is a support library for the `groom` suite of application.
*/
package config

import (
	"os"

	"github.com/pspiagicw/colorlog"
)

// Helper struct contains the prefix (Environment Variable) and suffix (constant)
type configLocation struct {
	// Environment Variable in the path
	prefix string
	// Constant in the path
	suffix string
}

/*
Global variable defines the location of the config file to be present.
The order of the array determines the order of the config location checking.

Locations for parsing (In order of preference)
- /etc/groom/config.toml
- $XDG_CONFIG_HOME/.config/config.toml
- $HOME/.config/config.toml
*/
var GROOM_CONFIG_LOCATIONS []configLocation = []configLocation{
	configLocation{
		prefix: "",
		suffix: "/etc/groom/groom.toml",
	},
	configLocation{
		prefix: "XDG_CONFIG_HOME",
		suffix: "/groom/config.toml",
	},
	configLocation{
		prefix: "HOME",
		suffix: "/.config/groom/config.toml",
	},
}

/* Structure stores all configuration options.
These include
- GROOM_INSTALLATION (For `groom-install`)
-
*/

type Plugin struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
}
type GroomConfig struct {
	GroomBin   string `toml:"groomBin"`
	ConfigPath string

	Plugins []Plugin `toml:"plugin"`
}

/*
Function parses the configuration file.
If no config found at default locations, return good-enough values.
*/
func ParseConf() *GroomConfig {

	path, err := getConfigPath()

	if err != nil {
		return generateDefaultConfig()
	}

	config, err := decodeConfig(path)

	if err != nil {
		return generateDefaultConfig()
	}

	return config
}
func saveConfig(config *GroomConfig) {

	path := config.ConfigPath

	if path == "" {
		colorlog.LogError("There was a error reading groom config. Thus cannot save config.")

	}

	contents := encodeConfig(config)

	os.WriteFile(path, contents, 0644)

}
func AddSubcommand(name string, description string) {
	config := ParseConf()

	for _, plugin := range config.Plugins {
		if plugin.Name == name {
			return
		}
	}

	config.Plugins = append(config.Plugins, Plugin{Name: name, Description: description})

	saveConfig(config)
}

func generateDefaultConfig() *GroomConfig {
	return &GroomConfig{
		GroomBin: "~/.local/bin",
	}
}
