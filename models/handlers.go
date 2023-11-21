package models

import "net/http"

type HandlerInterface interface {
	GetEmotions(w http.ResponseWriter, r *http.Request, videoUrl string)
	EmotionsMiddleware(func(w http.ResponseWriter, r *http.Request, videoUrl string)) http.Handler
}
