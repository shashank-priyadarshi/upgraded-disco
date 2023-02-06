package middleware

import (
	"fmt"
	"net/http"
	"server/config"
	"strings"
)

func InternalOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if !strings.Contains(origin, config.FetchConfig().SERVERORIGIN) {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ExternalOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println(origin)
		if !(strings.Contains(origin, "http://localhost:4200") || strings.Contains(origin, "https://ssnk.in")) {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
