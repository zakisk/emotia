package handlers

import (
	"github.com/zakisk/emotia/models"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.SugaredLogger
	Service models.YoutubeInterface
}

func NewHandler(logger *zap.SugaredLogger, service models.YoutubeInterface) models.HandlerInterface {
	return &Handler{logger: logger, Service: service}
}
