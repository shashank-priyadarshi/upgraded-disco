package todos

import (
	"fmt"
	"net/http"
	"server/config"
	"server/mongoconnection"
)

func returnTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.TODOS)
	w.Write(response)
}
