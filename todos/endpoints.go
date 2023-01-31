package todos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/mongoconnection"
)

func ReturnTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	reqStatus := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.TODOS)
	json.NewEncoder(w).Encode(reqStatus)
}
