package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"

	m "api.beermenu.com/models"
)

var Data m.TomlConfig

func init() {

	log.Println("Config Package")

	if _, err := toml.DecodeFile("config/config.dev.toml", &Data); err != nil {

		fmt.Println(err)
		return

	}

}
