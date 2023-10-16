package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/baptistemehat/go-leadsheet/core/common/logger"
	"github.com/rs/zerolog"
)

func main() {

	// Input flags
	var configPath string
	flag.StringVar(&configPath, "c", "", "Path to config.yaml file")

	flag.Usage = func() {
		fmt.Printf("Usage of go-leadsheet/core: \n")
		fmt.Printf("  ./core -c CONFIG_PATH\n")
		flag.PrintDefaults() // prints default usage
	}
	flag.Parse()

	// TODO create a Flag struct to store all input params and handle usage
	if configPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	logger.Init(zerolog.DebugLevel)

	app, err := NewApp(configPath)
	if err != nil {
		logger.Logger.Fatal().Msgf("%s", err)
	}

	app.Run()
}
