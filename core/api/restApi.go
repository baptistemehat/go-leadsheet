package api

import (
	"core/httpResponse"
	"core/pdfGenerator"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RestApi struct {
	pdfGenerator *pdfGenerator.PdfGenerator
	endpoints    map[string]func(http.ResponseWriter, *http.Request)
}

// NewRestApi returns a new RestApi instance
func NewRestApi(p *pdfGenerator.PdfGenerator) (*RestApi, error) {

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

	for route, handler := range restApi.endpoints {
		r.HandleFunc(route, handler)
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// ENDPOINTS
// health
func (ri *RestApi) health(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received API request: health")
	switch r.Method {
	case http.MethodGet:
		httpResponse.SendResponse(w, "ok")
	default:
		httpResponse.MethodNotAllowed(w)
	}
}

// song
func (restApi *RestApi) song(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received API request: song")

	type Msg struct {
		Source string `json:"source"`
	}

	msg := &Msg{}

	switch r.Method {
	case http.MethodGet:
		switch restApi.pdfGenerator.Status() {
		case pdfGenerator.StatusDone:
			httpResponse.ServeFile(w, r, restApi.pdfGenerator.Output())
		default:
			httpResponse.BadRequest(w)
		}
		return
	case http.MethodPost:

		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			httpResponse.BadRequest(w)
			return
		}

		// TODO : make GeneratePdf start goroutines so that we won't need to use defer statements
		defer restApi.pdfGenerator.GeneratePdf(msg.Source)
		httpResponse.Accepted(w)

	default:
		httpResponse.MethodNotAllowed(w)
	}
}

// status
func (restApi *RestApi) status(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received API request: status")

	switch r.Method {
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
