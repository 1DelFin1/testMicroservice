package delete

import (
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/1DelFin1/testMicroservice/internal/lib/api/response"
	"github.com/1DelFin1/testMicroservice/internal/lib/sl"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Deleter interface {
	DeleteUser(userID int64) error
}

func New(log *slog.Logger, deleter Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		userID := chi.URLParam(r, "userID")

		userIDInt, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			log.Info("invalid userID", slog.String("userID", userID))
			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		err = deleter.DeleteUser(userIDInt)
		if err != nil {
			log.Error("failed to delete user", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to delete user"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
