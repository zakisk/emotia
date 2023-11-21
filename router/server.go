package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zakisk/emotia/models"
)

type Router struct {
	R *mux.Router
}

func NewRouter(h models.HandlerInterface) *Router {
	router := mux.NewRouter()

	router.Handle("/get-comments-emotions", h.EmotionsMiddleware(h.GetEmotions)).Methods(http.MethodGet).Queries("videoUrl", "{[a-zA-Z]}")

	return &Router{R: router}
}
