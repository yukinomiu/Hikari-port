package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var cfg *config

type config struct {
	ListenAddress string `json:"listenAddress"`
	ListenPort    uint16 `json:"listenPort"`
	TargetAddress string `json:"targetAddress"`
	TargetPort    uint16 `json:"targetPort"`
}

func loadConfig() {
	// read config from args
	var configFilePath string
	args := os.Args
	if len(args) >= 2 {
		configFilePath = args[1]
	} else {
		cd := args[0]
		configFilePath = path.Join(cd, "../config.json")
	}

	log.Printf("loading config file '%v'\n", configFilePath)

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("read config file '%v' err, %v\n", configFilePath, err)
	}

	cfg = &config{}
	json.Unmarshal(data, cfg)
	log.Printf("using config file '%v'\n", configFilePath)
	log.Println("config:\n", string(data))

	// init status
	initStatus()
}
