/*
Package parses the config (global and user) and provides to the application.
It is a support library for the `groom` suite of application.
*/
package config

import (
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
        suffix: "/etc/groom/groom.conf",
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
type GroomConfig struct {
    GroomInstallation string `toml:"groomInstallation"`
    GroomConfigLocation string
}

/*
Function parses the configuration file.
If no config found at default locations, return good-enough values.
*/
func ParseConf() *GroomConfig {

    path , err := getConfigPath()

    if err != nil {
        return generateDefaultConfig()
    }

    config , err := decodeConfig(path)

    if err != nil {
        return generateDefaultConfig()
    }

    return config

}

func generateDefaultConfig() *GroomConfig {
    return &GroomConfig {
        GroomInstallation: "~/.local/groom",
    }
}

