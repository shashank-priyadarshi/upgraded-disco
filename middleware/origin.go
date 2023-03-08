package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"server/auth"
	"server/config"
	"strings"
)

type RouteHandler struct{}

func (*RouteHandler) InternalOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if !strings.Contains(origin, config.FetchConfig().SERVERORIGIN) {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (*RouteHandler) ExternalOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if !(strings.Contains(origin, config.FetchConfig().ALLOWEDORIGIN)) {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (*RouteHandler) AddResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding,X-CSRF-Token, Authorization, Referrer-Policy")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (*RouteHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := auth.User{}
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing authorization header"))
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization header"))
			return
		}

		var err, credentials = error(nil), []byte{}

		switch authHeaderParts[0] {
		case "Basic":
			user.Action = 1
			credentials, err = base64.StdEncoding.DecodeString(authHeaderParts[1])
		case "Bearer":
			user.Action = 2
			credentials = []byte(authHeaderParts[1])
		default:
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization method"))
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization credentials"))
			return
		}

		usernamePassword := strings.Split(string(credentials), ":")
		if len(usernamePassword) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization credentials"))
			return
		}

		user.Username = usernamePassword[0]
		user.Password = usernamePassword[1]

		rawCreds, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please try again"))
			return
		}
		token, err := user.ParseCredentials(rawCreds)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization credentials"))
			return
		}

		w.WriteHeader(http.StatusOK)
		if user.Action == 1 {
			w.Write([]byte(token))
		}
		next.ServeHTTP(w, r)

	})
}
