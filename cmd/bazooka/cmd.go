package main

import (
	"flag"
	"log"

	_ "github.com/sirupsen/logrus"

	"bazooka/internal/bazooka/core"
)


var (
	configFile string
)

func loadFlags() {
	flag.StringVar(&configFile, "c", "", "Path to config file.")
	flag.Parse()
}

func main() {
	loadFlags()
	if "" == configFile {
		log.Fatal("Config file not specified.")
	}

	var err error = nil

	cfg := core.GetConfig()
	err = cfg.FromFile(configFile)
	if nil != err {
		log.Fatal(err.Error())
	}

	_, err = core.GetDbConn()
	if nil != err {
		log.Fatal(err.Error())
	}
}
