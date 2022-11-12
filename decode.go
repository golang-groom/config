package config

import (
	"github.com/BurntSushi/toml"
)

// Function decodes a given file and returns `GroomConfig`
func decodeConfig(path string) (*GroomConfig , error){

    var config GroomConfig

    _ , err := toml.DecodeFile(path, &config)

    if err != nil {
        return nil , err
    }

    config.GroomConfigLocation = path

    return &config , nil

}
