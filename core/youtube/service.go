package youtube

import (
	"github.com/zakisk/emotia/models"
	"google.golang.org/api/youtube/v3"
)

type YoutubeService struct {
	Service *youtube.CommentThreadsService
}

func NewYoutubeService(service *youtube.CommentThreadsService) models.YoutubeInterface {
	return &YoutubeService{Service: service}
}

func (ys *YoutubeService) GetTopLevelComments(options *models.GetCommentsOptions) (*youtube.CommentThreadListResponse, error) {
	if err := options.ArePartsValid(); err != nil {
		return nil, err
	}
	call := ys.Service.List(options.Parts)
	return call.Do(options.GetOptions()...)
}

func (ys *YoutubeService) GetRepliesOfComment(options *models.GetCommentsOptions) (*youtube.CommentThreadListResponse, error) {
	if err := options.ArePartsValid(); err != nil {
		return nil, err
	}

	call := ys.Service.List(options.Parts)
	return call.Do(options.GetOptions()...)
}
