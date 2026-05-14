package httpapi

import (
	"net/http"

	"github.com/1DelFin1/testMicroservice/internal/domains"
)

func RequestIDMiddleware(genID domains.IDGenerator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := genID.NewID()
			w.Header().Set("Request-ID", requestID)
			next.ServeHTTP(w, r)
		})
	}
}
