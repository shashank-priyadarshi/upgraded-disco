package middleware

import (
	"fmt"
	"net/http"
	"server/config"
)

func OriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Printf("%v/trigger\n", config.FetchConfig().SERVERORIGIN)
		if origin != fmt.Sprintf("%v/trigger", config.FetchConfig().SERVERORIGIN) {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
