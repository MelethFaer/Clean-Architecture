package handler

import (
	"SandBox/internal/domain/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func responseEntityToJSON(entity interface{}, ctx *gin.Context) {
	byteData, err := json.MarshalIndent(entity, "", "    ")
	if err != nil {
		log.Println(err)
	}
	ctx.Data(http.StatusOK, "application/json", byteData)
}

func (h *Handler) getAllTracks(ctx *gin.Context) {
	tracks := &[]entity.Track{}
	tracks, err := h.services.Track.GetAll()
	if err != nil {
		log.Println(err)
	}
	responseEntityToJSON(tracks, ctx)
}

func (h *Handler) getUserPlaylist(ctx *gin.Context) {
	tracks := &[]entity.Track{}

	userId := ctx.Param("id")
	tracks, err := h.services.Track.GetUserPlaylist(userId)
	if err != nil {
		log.Println(err)
	}
	responseEntityToJSON(tracks, ctx)
}
