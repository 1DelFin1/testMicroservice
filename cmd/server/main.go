package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: buildHTTPHandler(),
	}

	log.Printf("server %s is running", srv.Addr)
	_ = srv.ListenAndServe()
}
