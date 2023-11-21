package errors

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

var (
	ErrInvalidVideoUrlCode      = "100"
	ErrEmptyAPIKeyCode          = "101"
	ErrYoutubeServiceCreateCode = "102"
)

var EmptyError error = fmt.Errorf("")

func ErrInvalidVideoUrl(err error) error {
	return errors.New(ErrInvalidVideoUrlCode, errors.Fatal,
		[]string{"Encountered an error while validating video URL"},
		[]string{err.Error()},
		[]string{"URL is incorrect"},
		[]string{"Copy URL from youtube web application."})
}

func ErrYoutubeServiceCreate(err error) error {
	return errors.New(ErrYoutubeServiceCreateCode, errors.Fatal,
		[]string{"Encountered an error while creating youtube servie"},
		[]string{err.Error()},
		[]string{"http.Client may be empty"},
		[]string{})
}

func ErrEmptyAPIKey(err error) error {
	return errors.New(ErrEmptyAPIKeyCode, errors.Fatal,
		[]string{"Youtube data API Key is empty"},
		[]string{err.Error()},
		[]string{"API_KEY env variable is not set"},
		[]string{"Set the enviroment variable with Youtube data API Key value."})
}
