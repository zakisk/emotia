package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/zakisk/emotia/models"
	"github.com/zakisk/emotia/pkg/errors"
)

func (h *Handler) EmotionsMiddleware(next func(w http.ResponseWriter, r *http.Request, options models.GetCommentsOptions)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		videoUrl := r.URL.Query().Get("videoUrl")
		if err := isYouTubeURL(videoUrl); err != nil {
			h.logger.Error(errors.ErrInvalidVideoUrl(err))
			http.Error(w, errors.ErrInvalidVideoUrl(err).Error(), http.StatusUnprocessableEntity)
			return
		}

		videoID, err := extractVideoID(videoUrl)
		if err != nil {
			h.logger.Error(errors.ErrInvalidVideoUrl(err))
			http.Error(w, errors.ErrInvalidVideoUrl(err).Error(), http.StatusUnprocessableEntity)
			return
		}

		options := models.GetCommentsOptions{
			Parts:      []string{models.Snippet, models.Replies},
			MaxResults: "100",
			VideoID:    videoID,
		}

		pageToken := r.URL.Query().Get("pageToken")
		if len(pageToken) > 0 {
			options.PageToken = pageToken
		}
		next(w, r, options)
	})
}

func isYouTubeURL(url string) error {
	youtubePattern1 := regexp.MustCompile(`^https?:\/\/(www\.)?youtube\.com\/watch\?v=[a-zA-Z0-9_-]{11}$`)
	youtubePattern2 := regexp.MustCompile(`^https?:\/\/(www\.)?youtu\.be\/[a-zA-Z0-9_-]`)
	youtubePattern3 := regexp.MustCompile(`^https?:\/\/(www\.)?youtube\.com\/shorts\/[a-zA-Z0-9_-]{11}$`)
	if !youtubePattern1.MatchString(url) && !youtubePattern2.MatchString(url) && !youtubePattern3.MatchString(url) {
		return fmt.Errorf("Given URL is not a youtube valid URL: `%s`", url)
	}
	return nil
}

func extractVideoID(url string) (string, error) {
	// youtubePattern := regexp.MustCompile(`(?:youtu\.be\/|youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=))([a-zA-Z0-9_-]{11})`)
	youtubePattern := regexp.MustCompile(`^(?:https?:\/\/)?(?:youtu\.be\/|(?:www\.)?youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:watch\/|v|e(?:mbed)?)\/|\S*?[?&]v=|shorts\/))([a-zA-Z0-9_-]{11})`)
	matches := youtubePattern.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("video ID not found in URL")
	}
	// The video ID is the second element in the matches slice
	videoID := matches[1]
	return videoID, nil
}
