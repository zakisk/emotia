package models

import (
	"net/http"
)

type HandlerInterface interface {
	GetEmotions(w http.ResponseWriter, r *http.Request, options GetCommentsOptions)
	EmotionsMiddleware(func(w http.ResponseWriter, r *http.Request, options GetCommentsOptions)) http.Handler
}
