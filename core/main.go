package main

import (
	"flag"
	"fmt"

	"github.com/baptistemehat/go-leadsheet/core/common/logger"
)

const DEFAULT_CONFIG_FILE string = ""

func main() {

	var configPath string
	flag.StringVar(&configPath, "c", "", "Path to config.yaml file")

	flag.Usage = func() {
		fmt.Printf("Usage of go-leadsheet/core: \n")
		fmt.Printf("  ./core -c CONFIG_PATH\n")
		flag.PrintDefaults() // prints default usage
	}
	flag.Parse()

	if configPath == "" {
		flag.Usage()
		return
	}

	logger.Init()

	app, err := NewApp(configPath)
	if err != nil {
		logger.Logger.Fatal().Msgf("%s", err)
	}

	app.Run()
}
