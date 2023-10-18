package main

import (
	"github.com/baptistemehat/go-leadsheet/core/api"
	"github.com/baptistemehat/go-leadsheet/core/common/logger"
	"github.com/baptistemehat/go-leadsheet/core/config"
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/formatting/latexformatting"
	parsing "github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing"
	"github.com/baptistemehat/go-leadsheet/core/pdfgenerator"
)

// TODO : make this path relative to the app
const path = "/home/baptiste/Programing/projects/github.com/baptistemehat/go-leadsheet/config.yaml"

type App struct {
	restApi *api.RestApi
}

// NewApp creates a new App
func NewApp() (*App, error) {

	builder := pdfgenerator.Builder{
		Parser:        parsing.InlineChordParser{},
		SongFormatter: &latexformatting.LatexSongFormatter{},
	}

	config, err := config.LoadConfiguration(path)
	if err != nil {
		return nil, err
	}

	pdfGenerator, err := pdfgenerator.NewPdfGenerator(builder, *config)
	if err != nil {
		logger.Logger.Fatal().Msgf("%s", err)
		return nil, err
	}

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
