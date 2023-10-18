package api

import (
	"encoding/json"
	"net/http"

	"github.com/baptistemehat/go-leadsheet/core/api/httpresponse"
	"github.com/baptistemehat/go-leadsheet/core/common/logger"
	"github.com/baptistemehat/go-leadsheet/core/pdfgenerator"

	"github.com/gorilla/mux"
)

type RestApi struct {
	pdfGenerator *pdfgenerator.PdfGenerator
	endpoints    map[string]func(http.ResponseWriter, *http.Request)
}

// NewRestApi creates a new RestApi
func NewRestApi(p *pdfgenerator.PdfGenerator) (*RestApi, error) {

	restApi := &RestApi{
		pdfGenerator: p,
		endpoints:    make(map[string]func(http.ResponseWriter, *http.Request)),
	}

	restApi.endpoints["/api/health"] = restApi.health
	restApi.endpoints["/api/song"] = restApi.song
	restApi.endpoints["/api/status"] = restApi.status

	return restApi, nil
}

// ListenAndServe listens and serves clients
func (restApi *RestApi) ListenAndServe(addr string) {
	r := mux.NewRouter()

	// associate endpoint with handler function
	for route, handler := range restApi.endpoints {
		r.HandleFunc(route, handler)
	}

	// start server
	http.Handle("/", r)
	logger.Logger.Fatal().Msgf("%s", http.ListenAndServe(addr, nil))
}

// **********************
//      ENDPOINTS
// **********************

// health handles /health endpoint
func (ri *RestApi) health(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Info().Msgf("received API request: %s %s %s", r.Method, r.URL.Path, r.URL.RawQuery)

	switch r.Method {

	// GET
	case http.MethodGet:
		httpresponse.SendResponse(w, "ok")

	default:
		httpresponse.MethodNotAllowed(w)
	}
}

// song handles /song endpoint
func (restApi *RestApi) song(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Info().Msgf("received API request: %s %s %s", r.Method, r.URL.Path, r.URL.RawQuery)

	switch r.Method {

	// GET
	case http.MethodGet:

		switch restApi.pdfGenerator.Status() {

		case pdfgenerator.StatusDone:
			httpresponse.ServeFile(w, r, restApi.pdfGenerator.Output())

		default:
			httpresponse.BadRequest(w)
		}
		return

	// POST
	case http.MethodPost:

		intputType := r.URL.Query().Get("type")

		switch intputType {
		case "text":

			// TODO : remove this feature since metadata are passed in leadsheet
			// TODO : create a Schema in a json file ? Shared file with UI
			type Msg struct {
				Title     string `json:"title"`
				Composer  string `json:"composer"`
				Leadsheet string `json:"leadsheet"`
			}

			msg := &Msg{}

			err := json.NewDecoder(r.Body).Decode(&msg)
			if err != nil {
				httpresponse.BadRequest(w)
				return
			}

			// TODO : ? add channels to transmit error
			go restApi.pdfGenerator.GeneratePdfFromBuffer(msg.Leadsheet)
			httpresponse.Accepted(w)

		case "file":
			// TODO : handle file upload

		default:
			httpresponse.BadRequest(w)
			return
		}

	default:
		httpresponse.MethodNotAllowed(w)
	}
}

// status handles /status endpoint
func (restApi *RestApi) status(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Info().Msgf("received API request: %s %s %s", r.Method, r.URL.Path, r.URL.RawQuery)

	switch r.Method {

	// GET
	case http.MethodGet:

		status := restApi.pdfGenerator.Status()
		if status.String() == "" {
			httpresponse.InternalServerError(w)
		}
		// TODO : send JSON response
		// TODO : add error messages if error
		httpresponse.SendResponse(w, status.String())

	default:
		httpresponse.MethodNotAllowed(w)
	}
}
