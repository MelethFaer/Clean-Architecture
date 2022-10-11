package handler

import (
	"SandBox/internal/domain/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) getAllUser(ctx *gin.Context) {
	tracks := &[]entity.Track{}

	tracks, err := h.services.Track.GetAll()
	if err != nil {
		log.Println(err)
	}

	byteData, err := json.MarshalIndent(tracks, "", "    ")
	if err != nil {
		log.Println(err)
	}

	ctx.Data(http.StatusOK, "application/json", byteData)
}
