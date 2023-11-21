package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/zakisk/emotia/pkg/errors"
)

func (h *Handler) EmotionsMiddleware(next func(w http.ResponseWriter, r *http.Request, videoUrl string)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		videoUrl := r.URL.Query().Get("videoUrl")
		if err := isYouTubeURL(videoUrl); err != nil {
			h.logger.Error(errors.ErrInvalidVideoUrl(err))
			http.Error(w, errors.ErrInvalidVideoUrl(err).Error(), http.StatusUnprocessableEntity)
			return
		}
		next(w, r, videoUrl)
	})
}

func isYouTubeURL(url string) error {
	youtubePattern1 := regexp.MustCompile(`^https?:\/\/(www\.)?youtube\.com\/watch\?v=[a-zA-Z0-9_-]{11}$`)
	youtubePattern2 := regexp.MustCompile(`^https?:\/\/(www\.)?youtu\.be\/[a-zA-Z0-9_-]`)
	if !youtubePattern1.MatchString(url) && !youtubePattern2.MatchString(url) {
		return fmt.Errorf("Given URL is not a youtube valid URL: `%s`", url)
	}
	return nil
}
