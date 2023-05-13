package httpResponse

import (
	"net/http"

	"github.com/baptistemehat/go-leadsheet/core/common/logger"
)

func SendResponse(w http.ResponseWriter, response string) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")

	status := http.StatusOK
	w.WriteHeader(status)

	w.Write([]byte(response))

	logger.Logger.Info().Msgf("response sent: %d %s", status, response)
}

func ServeFile(w http.ResponseWriter, r *http.Request, name string) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")

	http.ServeFile(w, r, name)

	logger.Logger.Info().Msgf("file served: %s", name)
}

func Accepted(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")

	status := http.StatusAccepted
	w.WriteHeader(http.StatusAccepted)

	response := "accepted"
	w.Write([]byte(response))

	logger.Logger.Info().Msgf("response sent: %d %s", status, response)
}

func BadRequest(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")

	status := http.StatusBadRequest
	w.WriteHeader(status)

	response := "bad request"
	w.Write([]byte("bad request"))

	logger.Logger.Info().Msgf("response sent: %d %s", status, response)

}

func InternalServerError(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")

	status := http.StatusInternalServerError
	w.WriteHeader(status)

	response := "internal sever error"
	w.Write([]byte(response))

	logger.Logger.Info().Msgf("response sent: %d %s", status, response)

}

func MethodNotAllowed(w http.ResponseWriter) {
	// need to set header before using WriteHeader, before using Write
	w.Header().Set("Access-Control-Allow-Origin", "*")

	status := http.StatusMethodNotAllowed
	w.WriteHeader(status)

	response := "method not allowed"
	w.Write([]byte(response))

	logger.Logger.Info().Msgf("response sent: %d %s", status, response)
}
