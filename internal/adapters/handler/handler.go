package handler

import (
	"SandBox/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/tracks", h.getAllTracks)
	router.GET("/tracks/user/:id", h.getUserPlaylist)
	return router
}
