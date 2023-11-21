package models

import (
	"fmt"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

// YoutubeInterface service interface
type YoutubeInterface interface {
	GetTopLevelComments(options *GetCommentsOptions) (*youtube.CommentThreadListResponse, error)
	GetRepliesOfComment(options *GetCommentsOptions) (*youtube.CommentThreadListResponse, error)
}

// Part Constants
const (
	// The `ID` that YouTube uses to uniquely identify the comment.
	ID = "id"

	// The `snippet` object contains basic details about the comment.
	Snippet = "snippet"

	// The `replies` includes replies to the base comment in the response.
	Replies = "replies"
)

// ModerationStatus constants
const (
	// Retrieve comment threads that are awaiting review by a moderator.
	// A comment thread can be included in the response if the top-level comment or
	// at least one of the replies to that comment are awaiting review.
	HeldForReview = "heldForReview"

	// Retrieve comment threads classified as likely to be spam.
	// A comment thread can be included in the response if the top-level comment or
	//  at least one of the replies to that comment is considered likely to be spam.
	LikelySpam = "likelySpam"

	// Retrieve threads of published comments. This is the default value.
	// A comment thread can be included in the response if its top-level comment has been published.
	Published = "published"
)

// Order constants
const (
	// Comment threads are ordered by time. This is the default behavior.
	Time = "time"

	// Comment threads are ordered by relevance.
	Relevance = "relevance"
)

// TestFormat constants
const (
	// Returns the comments in HTML format. This is the default value.
	Html = "html"

	// Returns the comments in plain text format.
	PlainText = "plainText"
)

type GetCommentsOptions struct {
	// Required parameters

	//The part parameter specifies a comma-separated list of one or
	// more commentThread resource properties that the API response will include.
	Parts []string

	// Filters

	// The allThreadsRelatedToChannelId parameter instructs the API to return all comment threads associated with the specified channel.
	// The response can include comments about the channel or about the channel's videos.
	AllThreadsRelatedToChannelId string

	// The channelId parameter instructs the API to return comment threads containing comments about the specified channel.
	// (The response will not include comments left on videos that the channel uploaded.)
	ChannelId string

	// The id parameter specifies a comma-separated list of comment thread IDs for the resources that should be retrieved.
	ID string

	// The videoId parameter instructs the API to return comment threads associated with the specified video ID.
	VideoID string

	// Optional Parametes

	/* The maxResults parameter specifies the maximum number of items that should be returned in the result set.
	Note: This parameter must not be used in conjuction with `id` parameter.*/
	MaxResults string

	/* Parameter to limit the returned comment threads to a particular moderation state.
	Note: This parameter must not be used in conjuction with `id` parameter.*/
	ModerationStatus string

	/* The parameter specifies the order in which the API response should list comment threads.
	Note: This parameter must not be used in conjuction with `id` parameter.*/
	Order string

	/* The pageToken parameter identifies a specific page in the result set that should be returned. In an API response, the nextPageToken property identifies the next page of the result that can be retrieved.
	Note: This parameter must not be used in conjuction with `id` parameter.*/
	PageToken string

	/* The searchTerms parameter instructs the API to limit the API response to only contain comments that contain the specified search terms.
	Note: This parameter must not be used in conjuction with `id` parameter.*/
	SearchTerms string

	// Set this parameter's value to html or plainText to instruct the API to return the comments
	// left by users in html formatted or in plain text.
	TextFormat string
}

func (o *GetCommentsOptions) GetOptions() []googleapi.CallOption {
	options := []googleapi.CallOption{}
	isIDEmpty := len(o.ID) == 0

	if len(o.AllThreadsRelatedToChannelId) > 0 {
		options = append(options, googleapi.QueryParameter("allThreadsRelatedToChannelId", o.AllThreadsRelatedToChannelId))
	}

	/* if `channelId` `videoId`, `id` are passed then `id` takes precendence and comments will be fetched for `id`
	Note: This policy is specific to this application. */
	if len(o.ChannelId) > 0 && len(o.VideoID) == 0 && isIDEmpty {
		options = append(options, googleapi.QueryParameter("channelId", o.ChannelId))
	}

	if !isIDEmpty {
		options = append(options, googleapi.QueryParameter("id", o.ID))
	}

	/* if `videoId`, `id` both are passed then `id` takes precendence and comments will be fetched for `id`
	Note: This policy is specific to this application. */
	if len(o.VideoID) > 0 && isIDEmpty {
		options = append(options, googleapi.QueryParameter("videoId", o.VideoID))
	}

	if len(o.MaxResults) > 0 && isIDEmpty {
		options = append(options, googleapi.QueryParameter("maxResults", o.MaxResults))
	}

	if len(o.ModerationStatus) > 0 && isIDEmpty {
		options = append(options, googleapi.QueryParameter("moderationStatus", o.ModerationStatus))
	}

	if len(o.Order) > 0 && isIDEmpty {
		options = append(options, googleapi.QueryParameter("order", o.Order))
	}

	if len(o.PageToken) > 0 && isIDEmpty {
		options = append(options, googleapi.QueryParameter("pageToken", o.PageToken))
	}

	if len(o.SearchTerms) > 0 && isIDEmpty {
		options = append(options, googleapi.QueryParameter("searchTerms", o.SearchTerms))
	}

	if len(o.TextFormat) > 0 {
		options = append(options, googleapi.QueryParameter("textFormat", o.TextFormat))
	}

	return options
}

func (o *GetCommentsOptions) ArePartsValid() error {
	for _, part := range o.Parts {
		if part != ID && part != Replies && part != Snippet {
			return fmt.Errorf("`%s` is invalid part name", part)
		}
	}
	return nil
}

func (o *GetCommentsOptions) isModerationStatusValid() error {
	if o.ModerationStatus != HeldForReview && o.ModerationStatus != LikelySpam && o.ModerationStatus != Published {
		return fmt.Errorf("`%s` is invalid moderationStatus", o.ModerationStatus)
	}
	return nil
}

func (o *GetCommentsOptions) isValidOrder() error {
	if o.Order != Time && o.Order != Relevance {
		return fmt.Errorf("`%s` is invalid order", o.Order)
	}
	return nil
}

func (o *GetCommentsOptions) isValidTextFormat() error {
	if o.TextFormat != Html && o.TextFormat != PlainText {
		return fmt.Errorf("`%s` is invalid text format", o.TextFormat)
	}
	return nil
}
