package main

import "github.com/baptistemehat/go-leadsheet/core/common/logger"

func main() {

	logger.Init()

	app, err := NewApp()
	if err != nil {
		logger.Logger.Fatal().Msgf("%s", err)
	}

	app.Run()
}
