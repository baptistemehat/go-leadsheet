package main

import (
	"github.com/baptistemehat/go-leadsheet/core/api"
	"github.com/baptistemehat/go-leadsheet/core/common/logger"
	"github.com/baptistemehat/go-leadsheet/core/config"
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/formatting/latexformatting"
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing"
	"github.com/baptistemehat/go-leadsheet/core/pdfgenerator"
)

type App struct {
	restApi *api.RestApi
}

// NewApp creates a new App
func NewApp(configPath string) (*App, error) {

	builder := pdfgenerator.Builder{
		Parser:        inlinechordparsing.InlineChordParser{},
		SongFormatter: &latexformatting.LatexSongFormatter{},
	}

	config, err := config.LoadConfiguration(configPath)
	if err != nil {
		return nil, err
	}

	pdfGenerator, err := pdfgenerator.NewPdfGenerator(builder, *config)
	if err != nil {
		return nil, err
	}

	restApi, err := api.NewRestApi(pdfGenerator)
	if err != nil {
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
