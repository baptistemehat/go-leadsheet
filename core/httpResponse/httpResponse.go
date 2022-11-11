package httpResponse

import (
	"net/http"
)

func SendResponse(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func ServeFile(w http.ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, name)
}

func Accepted(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("accepted"))
}

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request"))
}

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
}
