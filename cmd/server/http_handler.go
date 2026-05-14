package main

import (
	"net/http"

	"github.com/1DelFin1/testMicroservice/internal/infra"
	"github.com/1DelFin1/testMicroservice/internal/protocol/httpapi"
)

func buildHTTPHandler() http.Handler {
	idGen := infra.UUIDGenerator{}

	mux := http.NewServeMux()

	var handler http.Handler = mux
	handler = httpapi.RequestIDMiddleware(idGen)(handler)

	return handler
}
