package handlers

import (
	rediscache "github.com/golanguzb70/redis-cache"
	"github.com/Javokhdev/Yelp-Project/config"
	"github.com/Javokhdev/Yelp-Project/internal/usecase"
	"github.com/Javokhdev/Yelp-Project/pkg/logger"
)

type Handler struct {
	Logger  *logger.Logger
	Config  *config.Config
	UseCase *usecase.UseCase
	Redis   rediscache.RedisCache
}

func NewHandler(l *logger.Logger, c *config.Config, useCase *usecase.UseCase, redis rediscache.RedisCache) *Handler {
	return &Handler{
		Logger:  l,
		Config:  c,
		UseCase: useCase,
		Redis:   redis,
	}
}
