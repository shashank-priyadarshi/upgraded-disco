package todos

import (
	"fmt"
	"net/http"
	"server/config"
	mongoconnection "server/db/mongo"

	logger "github.com/rs/zerolog/log"
)

func returnTodos(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.TODOS)
	w.Write(response)
}
