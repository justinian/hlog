package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type Log struct {
	Message string `json:"log"`
}

func main() {
	configfile := os.Getenv("HOME") + "/.hakkarc"
	config, err := getConfig(configfile)
	if err != nil {
		log.Fatalf("%s: Could not read config file %s: %s",
			os.Args[0], configfile, err.Error())
	}

	url := config.Logs.Url + "&token=" + config.Logs.Token
	data, err := json.Marshal(Log{strings.Join(os.Args[1:], " ")})
	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{}
	if !config.Connection.Verify_Ssl {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	cl := &http.Client{Transport: tr}

	buf := bytes.NewBuffer(data)
	resp, err := cl.Post(url, "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status[:2] != "20" {
		log.Fatal(resp.Body)
	}
}
