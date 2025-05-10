package service

import (
	"log"

	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/logger"
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/rest_err"
	model "github.com/LucasJandrey5/url-shortener/backend/src/model/user"
	"go.uber.org/zap"
)

func (*userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	log.Println(userDomain)

	return nil
}
