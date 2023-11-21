package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/zakisk/emotia/pkg/errors"
	// "github.com/zakisk/emotia/pkg/utils"
	// "google.golang.org/api/googleapi"
	// "google.golang.org/api/option"
	// "google.golang.org/api/youtube/v3"
)

func (h *Handler) GetEmotions(w http.ResponseWriter, r *http.Request, videoUrl string) {
	videoID, err := extractVideoID(videoUrl)
	if err != nil {
		h.logger.Error(errors.ErrInvalidVideoUrl(err))
		http.Error(w, errors.ErrInvalidVideoUrl(err).Error(), http.StatusUnprocessableEntity)
		return
	}
	 // something with `videoID`
	 fmt.Println(videoID)
	
}

func extractVideoID(url string) (string, error) {
	youtubePattern := regexp.MustCompile(`(?:youtu\.be\/|youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=))([a-zA-Z0-9_-]{11})`)
	matches := youtubePattern.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("video ID not found in URL")
	}
	// The video ID is the second element in the matches slice
	videoID := matches[1]
	return videoID, nil
}
