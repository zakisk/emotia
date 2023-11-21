package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zakisk/emotia/core/youtube"
	"github.com/zakisk/emotia/handlers"
	"github.com/zakisk/emotia/pkg/errors"
	"github.com/zakisk/emotia/router"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	youtubeV3 "google.golang.org/api/youtube/v3"
)

func main() {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		logger.Fatal(errors.ErrEmptyAPIKey(errors.EmptyError))
	}

	svc, err := youtubeV3.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		logger.Fatal(errors.ErrYoutubeServiceCreate(err))
	}

	commentThreadService := youtubeV3.NewCommentThreadsService(svc)
	service := youtube.NewYoutubeService(commentThreadService)

	h := handlers.NewHandler(logger, &service)
	router := router.NewRouter(h)
	//creating custom server in order to set the fields as per my requirement
	s := &http.Server{
		Addr:         ":9254",
		Handler:      router.R,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Info("Got signal: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
