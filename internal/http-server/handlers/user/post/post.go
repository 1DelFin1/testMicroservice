package post

import (
	"log/slog"
	"net/http"

	"github.com/1DelFin1/testMicroservice/internal/http-server/handlers/user"
	resp "github.com/1DelFin1/testMicroservice/internal/lib/api/response"

	"github.com/1DelFin1/testMicroservice/internal/lib/sl"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Creator interface {
	CreateUser(userData *user.Data) (int64, error)
}

type Response struct {
	resp.Response
	ID int64 `json:"user_id"`
}

func New(log *slog.Logger, creator Creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.post.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var userData user.Data
		err := render.DecodeJSON(r.Body, &userData)
		if err != nil {
			log.Error("failed to parse request", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to parse request"))
			return
		}

		// TODO: validation

		userID, err := creator.CreateUser(&userData)
		if err != nil {
			log.Error("failed to create user", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to create user"))
			return
		}

		render.JSON(w, r, Response{
			Response: resp.OK(),
			ID:       userID,
		})
	}
}
