package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

type Log struct {
	Message string `json:"log"`
	Privacy string `json:"state"`
}

func main() {
	configfile := os.Getenv("HOME") + "/.hakkarc"
	config, err := getConfig(configfile)
	if err != nil {
		log.Fatalf("%s: Could not read config file %s: %s",
			os.Args[0], configfile, err.Error())
	}

	privacy := flag.String("priv", config.Logs.Default_Privacy, "Privacy level to log with (public, anonymous, private)")
	flag.Parse()

	logItem := Log{
		Message: strings.Join(flag.Args(), " "),
		Privacy: *privacy,
	}

	url := config.Logs.Url + "&token=" + config.Logs.Token
	data, err := json.Marshal(logItem)
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
