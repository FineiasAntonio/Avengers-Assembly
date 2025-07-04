package middleware

import "net/http"

type CORSMiddleware struct {
}

func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

func (c *CORSMiddleware) LiberarCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, PUT, DELETE")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
