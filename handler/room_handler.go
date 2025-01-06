package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"homework/service"
)

type RoomController struct {
	service service.RoomService
	logger  *zap.Logger
}

func NewRoomController(service service.RoomService, logger *zap.Logger) *RoomController {
	return &RoomController{service: service, logger: logger}
}

func (ctrl *RoomController) Create(c *gin.Context) {

}
