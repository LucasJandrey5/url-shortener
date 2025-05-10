package controller

import (
	"net/http"

	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/logger"
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/validation"
	"github.com/LucasJandrey5/url-shortener/backend/src/controller/model/request"
	model "github.com/LucasJandrey5/url-shortener/backend/src/model/user"
	"github.com/LucasJandrey5/url-shortener/backend/src/model/user/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		logger.Error("Error creating user", restErr,
			zap.String("journey", "createUser"),
		)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Username,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusCreated, "User created successfully")
}
