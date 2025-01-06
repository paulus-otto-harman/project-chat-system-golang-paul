package handler

import (
	"homework/database"
	"homework/domain"
	"homework/infra/jwt"
	"homework/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	AuthHandler          AuthController
	ChatHandler          ChatController
	PasswordResetHandler PasswordResetController
	RoomHandler          RoomController
	UserHandler          UserController
}

func NewHandler(service service.Service, logger *zap.Logger, rdb database.Cacher, jwt jwt.JWT) *Handler {
	return &Handler{
		AuthHandler:          *NewAuthController(service.Auth, logger, rdb, jwt),
		ChatHandler:          *NewChatController(service.Chat, logger, rdb),
		PasswordResetHandler: *NewPasswordResetController(service, logger),
		RoomHandler:          *NewRoomController(service.Room, logger),
		UserHandler:          *NewUserController(service.User, logger),
	}
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func GoodResponseWithPage(c *gin.Context, message string, statusCode, total, totalPages, page, Limit int, data interface{}) {
	c.JSON(statusCode, domain.DataPage{
		Status:      true,
		Message:     message,
		Total:       int64(total),
		Pages:       totalPages,
		CurrentPage: uint(page),
		Limit:       uint(Limit),
		Data:        data,
	})
}
