package auth

import "net/http"

// AuthMiddleware is a middleware that ensures only authenticated users can access the WebSocket endpoint
func (a *AuthService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the session token from the request
		token := r.Header.Get("Authorization")

		// Check if the session token is valid
		user := a.GetSessionUser(token)
		if user == nil {
			// Invalid session token or unauthenticated user
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// User is authenticated, continue to the WebSocket handler
		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:5173" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
