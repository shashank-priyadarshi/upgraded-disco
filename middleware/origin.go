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
		if !(strings.Contains(origin, config.FetchConfig().ALLOWEDORIGIN)) {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AddResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding,X-CSRF-Token, Authorization, Referrer-Policy")

		next.ServeHTTP(w, r)
	})
}
