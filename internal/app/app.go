package app

import (
	"net/http"

	"github.com/1DelFin1/testMicroservice/config"
	"github.com/1DelFin1/testMicroservice/internal/http-server/handlers/user/delete"
	"github.com/1DelFin1/testMicroservice/internal/http-server/handlers/user/get"
	"github.com/1DelFin1/testMicroservice/internal/http-server/handlers/user/post"
	"github.com/1DelFin1/testMicroservice/internal/lib/logger"
	"github.com/1DelFin1/testMicroservice/internal/lib/sl"
	"github.com/1DelFin1/testMicroservice/internal/storage/pg"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run(cfg *config.Config) {
	log := logger.SetupLogger(cfg.App.Env)
	log.Info("Logger initialized")

	storage, err := pg.New(cfg)
	if err != nil {
		log.Error("unable to connect to database", sl.Err(err))
		return
	}
	defer storage.Close(log)
	log.Info("PG pool initialized")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", post.New(log, storage))
		r.Get("/{userID}", get.New(log, storage))
		r.Delete("/{userID}", delete.New(log, storage))
	})

	srv := &http.Server{
		Addr:    cfg.App.Address,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("Failed to start server", sl.Err(err))
	}

	log.Error("server shutdown")
}
