package get

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/1DelFin1/testMicroservice/internal/http-server/handlers/user"
	resp "github.com/1DelFin1/testMicroservice/internal/lib/api/response"
	"github.com/1DelFin1/testMicroservice/internal/lib/sl"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Getter interface {
	GetUser(userID int64) (*user.Data, error)
}

type Response struct {
	resp.Response
	User *user.Data `json:"user,omitempty"`
}

func New(log *slog.Logger, getter Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.get.New"

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

		userData, err := getter.GetUser(userIDInt)
		if err != nil {
			log.Error("failed to get user", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to get user"))
			return
		}

		render.JSON(w, r, Response{
			Response: resp.OK(),
			User:     userData,
		})
	}
}
