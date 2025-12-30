package show

import (
	"net/http"

	"github.com/Pxe2k/halyk-task/pkg"
	"github.com/gorilla/mux"
)

type Handler struct {
	*mux.Router
	uc UseCase
}

func New(uc UseCase) *Handler {
	h := Handler{
		mux.NewRouter(),
		uc,
	}

	api := h.PathPrefix("/api").Subrouter()
	api.Use(pkg.SetMiddlewareJSON)

	show := api.PathPrefix("/show").Subrouter()

	show.HandleFunc("/reserve-seats", h.researveSeats).Methods(http.MethodPost)

	return &h
}
