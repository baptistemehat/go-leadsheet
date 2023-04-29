package httpResponse

import (
	"net/http"
)

func SendResponse(w http.ResponseWriter, response string) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func ServeFile(w http.ResponseWriter, r *http.Request, name string) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, name)
}

func Accepted(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("accepted"))
}

func BadRequest(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request"))
}

func InternalServerError(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func MethodNotAllowed(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
}
