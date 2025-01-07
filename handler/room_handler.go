package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"homework/domain"
	"homework/service"
	"log"
	"net/http"
)

type RoomController struct {
	service service.RoomService
	logger  *zap.Logger
}

func NewRoomController(service service.RoomService, logger *zap.Logger) *RoomController {
	return &RoomController{service: service, logger: logger}
}

func (ctrl *RoomController) Create(c *gin.Context) {
	var room domain.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		log.Println(err.Error())
		BadResponse(c, "invalid input", http.StatusBadRequest)
		return
	}

	if err := ctrl.service.SaveRoom(&room); err != nil {
		log.Println(err.Error())
		BadResponse(c, "failed to create chat room", http.StatusInternalServerError)
		return
	}
	GoodResponseWithData(c, "room created successfully", http.StatusOK, room)
}
