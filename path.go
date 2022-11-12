package config

import (
	"fmt"
	"os"
)

/*
Function checks paths in GROOM_CONFIG_LOCATIONS and gives the most non-privileged location.
Gives error if none if it encounters error
*/
func getConfigPath() (string , error) {
    var configPath string

    for _, location := range GROOM_CONFIG_LOCATIONS {

        prefix := location.prefix

        suffix := location.suffix

        if prefix == "" {

            exists := locationExists(suffix)

            if exists {
                configPath = suffix
            }

            continue
        }

        calculatedPath := checkFullPath(prefix, suffix)

        if calculatedPath != "" {

            configPath = calculatedPath
        }
    }

    if configPath == "" {
        return "" , fmt.Errorf("No config location present!")
    }

    return configPath , nil
}

/*
Helper function to expand prefix and check full path location for existence
*/
func checkFullPath(prefix, suffix string) string {
    expandedPrefix := expandPrefix(prefix)

    if expandedPrefix == "" {
        return ""
    }

    path := expandedPrefix + suffix

    exists := locationExists(path)

    if exists {
        return path
    }

    return ""
}

// Helper function to check file existence
func locationExists(path string) bool {
    _ , err := os.Stat(path)
    if err != nil {
        return false
    }
    return true
}

// Helper function to expand given environment variable
func expandPrefix( prefix string) string {
    return os.Getenv(prefix)
}

