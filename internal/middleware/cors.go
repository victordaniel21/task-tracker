package middleware

import "net/http"

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Allow any origin (for development). In production, change "*" to your domain.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 2. Allow specific methods (GET, POST, PUT, DELETE)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 3. Allow headers (Content-Type is needed for JSON)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 4. Handle "Preflight" requests
		// Browsers send an OPTIONS request first to check permissions.
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
