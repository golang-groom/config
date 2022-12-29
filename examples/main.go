package main

import (
	"fmt"

	"github.com/golang-groom/config"
)

func main() {
	fmt.Println(config.ParseConf())

	config.AddSubcommand("groom-fake", "Fake groom subcommand.")

	fmt.Println(config.ParseConf())

}
