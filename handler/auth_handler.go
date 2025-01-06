package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"homework/database"
	"homework/domain"
	"homework/infra/jwt"
	"homework/service"
	"log"
	"net/http"
	"strconv"
)

type AuthController struct {
	service service.AuthService
	logger  *zap.Logger
	jwt     jwt.JWT
	cacher  database.Cacher
}

func NewAuthController(service service.AuthService, logger *zap.Logger, cacher database.Cacher, jwt jwt.JWT) *AuthController {
	return &AuthController{service, logger, jwt, cacher}
}

// Login endpoint
// @Summary User login
// @Description authenticate user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param domain.User body domain.User true " "
// @Success 200 {object} handler.Response "user authenticated"
// @Failure 401 {object} handler.Response "invalid username and/or password"
// @Failure 500 {object} handler.Response "server error"
// @Router  /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var login domain.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Println(err)
		BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	// Ambil IP address dari request
	ip := c.ClientIP()
	if ip == "" {
		ip = "unknown" // Default jika IP tidak ditemukan
		ctrl.logger.Warn("Failed to retrieve client IP")
	}

	user, err := ctrl.service.Login(login.Email, login.Password)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}

	ctrl.logger.Info("found", zap.Any("user", user))

	// Buat token JWT
	token, err := ctrl.jwt.CreateToken(user.Email, ip, strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		ctrl.logger.Error("Failed to create JWT token", zap.Error(err))
		BadResponse(c, "failed to create token", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("User logged in successfully", zap.String("email", user.Email))
	GoodResponseWithData(c, "user authenticated", http.StatusOK, gin.H{"token": token})
}
