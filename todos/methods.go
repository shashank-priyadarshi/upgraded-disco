package todos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/mongoconnection"
)

func returnTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.TODOS)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("error while encoding request data for endpoint %v: %v\n", r.URL.Path, err)
	}
}
