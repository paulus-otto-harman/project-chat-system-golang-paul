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

// Create Chat Room endpoint
// @Summary Create chat room
// @Description create chat room (channel) for users
// @Tags Chat
// @Accept  json
// @Produce  json
// @Param Room body Room true " "
// @Success 200 {object} handler.Response "room created successfully"
// @Failure 401 {object} handler.Response "unauthorized"
// @Failure 422 {object} handler.Response "invalid input"
// @Failure 500 {object} handler.Response "failed to create chat room"
// @Router  /rooms [post]
func (ctrl *RoomController) Create(c *gin.Context) {
	var room domain.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		log.Println(err.Error())
		BadResponse(c, "invalid input", http.StatusUnprocessableEntity)
		return
	}

	if err := ctrl.service.SaveRoom(&room); err != nil {
		log.Println(err.Error())
		BadResponse(c, "failed to create chat room", http.StatusInternalServerError)
		return
	}
	GoodResponseWithData(c, "room created successfully", http.StatusOK, room)
}

type Room struct {
	Name  string `json:"name" example:"Alumni Lumoshive Academy"`
	Users []User `json:"users"`
}

type User struct {
	ID uint `json:"id" example:"6"`
}
