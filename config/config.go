package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server serverInfo
	Sql    sqlInfo
}

type serverInfo struct {
	Port string
}

type sqlInfo struct {
	Database string `toml:"database"`
	URI      string `toml:"uri"`
}

var MainConfig Config

func init() {
	if _, err := toml.DecodeFile("/Users/davidtamayodomenech/go/src/ardidas/configuration.toml", &MainConfig); err != nil {
		fmt.Println(err)
		return
	}
}
