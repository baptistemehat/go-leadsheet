package main

import (
	"log"

	"github.com/baptistemehat/go-leadsheet/core/api"
	"github.com/baptistemehat/go-leadsheet/core/pdfGenerator"
	songFormatter "github.com/baptistemehat/go-leadsheet/core/song/formatter"
	"github.com/baptistemehat/go-leadsheet/core/song/parser"
)

type App struct {
	restApi *api.RestApi
}

func NewApp() (*App, error) {

	// a builder defines the parser and formatter to use
	builder := pdfGenerator.Builder{
		Parser:    parser.InlineChordParser{},
		Formatter: &songFormatter.LatexSongFormatter{},
	}

	// create pdf generator
	pdfGenerator, err := pdfGenerator.NewPdfGenerator(builder)
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
	log.Printf("App started, listening on " + port)
	app.restApi.ListenAndServe(port)
}
