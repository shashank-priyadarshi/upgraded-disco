package common

import (
	"fmt"
	"net/http"

	logger "github.com/rs/zerolog/log"
)

func InvalidEndpoint(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	http.Error(w, "Endpoint does not exist", http.StatusNotFound)
}
