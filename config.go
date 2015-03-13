package main

import (
	"code.google.com/p/gcfg"
)

type Config struct {
	Logs struct {
		Token string
		Url   string
	}
}

var defaultConfig = `
	[logs]
	url = https://www.hakkalabs.co/api/webhooks?service=custom
`

func getConfig(filename string) (*Config, error) {
	var config Config
	if err := gcfg.ReadStringInto(&config, defaultConfig); err != nil {
		panic("Failed to parse default config: " + err.Error())
	}

	err := gcfg.ReadFileInto(&config, filename)
	return &config, err
}
