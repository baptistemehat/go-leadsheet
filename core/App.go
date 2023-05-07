package main

import (
	"github.com/baptistemehat/go-leadsheet/core/api"
	"github.com/baptistemehat/go-leadsheet/core/common/logger"
	"github.com/baptistemehat/go-leadsheet/core/config"
	"github.com/baptistemehat/go-leadsheet/core/pdfGenerator"
	"github.com/baptistemehat/go-leadsheet/core/song/formatter"
	"github.com/baptistemehat/go-leadsheet/core/song/parser"
)

type App struct {
	restApi *api.RestApi
}

// NewApp creates a new App
func NewApp() (*App, error) {

	// a builder defines the parser and formatter to use
	builder := pdfGenerator.Builder{
		Parser:    parser.InlineChordParser{},
		Formatter: &formatter.LatexSongFormatter{},
	}

	path := "/home/baptiste/Programing/projects/github.com/baptistemehat/go-leadsheet/core/config/testResources/config.yaml"
	config, err := config.LoadConfiguration(path)

	if err != nil {
		return nil, err
	}

	// create pdf generator
	pdfGenerator, err := pdfGenerator.NewPdfGenerator(builder, *config)
	if err != nil {
		logger.Logger.Fatal().Msgf("%s", err)
		return nil, err
	}

	// create rest api
	restApi, err := api.NewRestApi(pdfGenerator)
	if err != nil {
		logger.Logger.Fatal().Msgf("%s", err)
		return nil, err
	}

	app := &App{
		restApi: restApi,
	}

	return app, nil
}

func (app *App) Run() {
	port := ":8000"
	logger.Logger.Info().Msgf("app started, listening on %s", port)
	app.restApi.ListenAndServe(port)
}
