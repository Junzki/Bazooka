package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"

	"bazooka/internal/bazooka/core"
	"bazooka/internal/pkg/assets"
)

var (
	configFile string
	dir        string
)

func loadFlags() {
	cwd, err := os.Getwd()
	if nil != err {
		cwd = ""
	}

	flag.StringVar(&configFile, "c", "", "Path to config file.")
	flag.StringVar(&dir, "chdir", cwd, "Working directory, current directory by default.")
	flag.Parse()
}

func main() {
	loadFlags()
	if "" == configFile {
		log.Fatal("Config file not specified.")
	}

	var err error = nil

	cfg := core.GetConfig()
	err = cfg.FromFile(assets.ExpandUserDir(configFile))
	if nil != err {
		log.Fatal(err.Error())
	}

	_, err = core.InitApp(cfg, dir)
	if nil != err {
		log.Fatal(err)
	}
}
