package config

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/pspiagicw/colorlog"
)

// Function decodes a given file and returns `GroomConfig`
func decodeConfig(path string) (*GroomConfig, error) {

	var config GroomConfig

	_, err := toml.DecodeFile(path, &config)

	if err != nil {
		return nil, err
	}

	config.ConfigPath = path

	return &config, nil

}

func encodeConfig(config *GroomConfig) []byte {

	var buffer bytes.Buffer

	encoder := toml.NewEncoder(&buffer)

	err := encoder.Encode(config)

	if err != nil {
		colorlog.LogError("Error converting toml file.")
	}

	return buffer.Bytes()

}
