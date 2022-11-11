package main

import (
	"core/api"
	"core/pdfGenerator"
	"log"
)

type App struct {
	restApi *api.RestApi
}

func NewApp() (*App, error) {

	pdfGenerator, err := pdfGenerator.NewPdfGenerator()
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
