package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"homework/domain"
	"homework/service"
	"net/http"
)

type UserController struct {
	service service.UserService
	logger  *zap.Logger
}

func NewUserController(service service.UserService, logger *zap.Logger) *UserController {
	return &UserController{service: service, logger: logger}
}

func (ctrl *UserController) All(c *gin.Context) {
	searchParam := domain.User{Email: c.Query("email")}

	if searchParam.Email == "" {
		BadResponse(c, "invalid parameter", http.StatusBadRequest)
		return
	}

	if _, err := ctrl.service.All(searchParam); err != nil {
		BadResponse(c, err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "email is valid", http.StatusOK, nil)
}

// Registration endpoint
// @Summary User Registration
// @Description register user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param domain.User body domain.User true " "
// @Success 200 {object} handler.Response "login successfully"
// @Failure 500 {object} handler.Response "server error"
// @Router  /register [post]
func (ctrl *UserController) Registration(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := ctrl.service.Register(&user); err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "user registered", http.StatusCreated, gin.H{"name": user.Name, "email": user.Email})
}

// Reset Password endpoint
// @Summary Reset Password
// @Description reset user password. user can only reset their password after successfully validated the OTP.
// @Description use 1052c225-9a44-4f61-a340-040ef44e8022 to reset the password using data from the seeder
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "OTP sent"
// @Failure 500 {object} handler.Response "failed to send OTP"
// @Router  /user/:id [put]
func (ctrl *UserController) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		BadResponse(c, "invalid parameter", http.StatusBadRequest)
		return
	}

	var newPassword NewPassword
	if err = c.ShouldBindJSON(&newPassword); err != nil {
		BadResponse(c, "invalid password", http.StatusUnprocessableEntity)
		return
	}

	if err = ctrl.service.UpdatePassword(id, newPassword.Password); err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "password changed", http.StatusOK, nil)
}

type NewPassword struct {
	Password        string `binding:"required" json:"password"`
	ConfirmPassword string `binding:"required,eqfield=Password" json:"confirm_password"`
}
