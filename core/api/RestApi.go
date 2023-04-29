package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/baptistemehat/go-leadsheet/core/httpResponse"
	"github.com/baptistemehat/go-leadsheet/core/pdfGenerator"

	"github.com/gorilla/mux"
)

type RestApi struct {
	pdfGenerator *pdfGenerator.PdfGenerator
	endpoints    map[string]func(http.ResponseWriter, *http.Request)
}

// NewRestApi creates a new RestApi
func NewRestApi(p *pdfGenerator.PdfGenerator) (*RestApi, error) {

	restApi := &RestApi{
		pdfGenerator: p,
		endpoints:    make(map[string]func(http.ResponseWriter, *http.Request)),
	}

	// define endpoints and handler functions
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
	log.Fatal(http.ListenAndServe(addr, nil))
}

// **********************
//      ENDPOINTS
// **********************

// health handles /health endpoint
func (ri *RestApi) health(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received API request: %s %s %s", r.Method, r.URL.Path, r.URL.Query())

	switch r.Method {

	// GET
	case http.MethodGet:
		httpResponse.SendResponse(w, "ok")

	default:
		httpResponse.MethodNotAllowed(w)
	}
}

// song handles /song endpoint
func (restApi *RestApi) song(w http.ResponseWriter, r *http.Request) {

	// TODO : find a good way to log API calls, with string reduction for query parameters
	log.Printf("Received API request: %s %s %s", r.Method, r.URL.Path, r.URL.Query())

	switch r.Method {

	// GET
	case http.MethodGet:

		switch restApi.pdfGenerator.Status() {

		case pdfGenerator.StatusDone:
			httpResponse.ServeFile(w, r, restApi.pdfGenerator.Output())

		default:
			httpResponse.BadRequest(w)
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
				httpResponse.BadRequest(w)
				return
			}

			// TODO : ? add channels to transmit error
			go restApi.pdfGenerator.GeneratePdfFromBuffer(msg.Leadsheet)
			httpResponse.Accepted(w)

		case "file":
			// TODO : handle file upload

		default:
			httpResponse.BadRequest(w)
			return
		}

	default:
		httpResponse.MethodNotAllowed(w)
	}
}

// status handles /status endpoint
func (restApi *RestApi) status(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received API request: %s %s %s", r.Method, r.URL.Path, r.URL.Query())

	switch r.Method {

	// GET
	case http.MethodGet:

		status := restApi.pdfGenerator.Status()
		if status.String() == "" {
			httpResponse.InternalServerError(w)
		}
		// TODO : send JSON response
		// TODO : add error messages if error
		httpResponse.SendResponse(w, status.String())

	default:
		httpResponse.MethodNotAllowed(w)
	}
}
