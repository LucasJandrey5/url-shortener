package service

import (
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/logger"
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/rest_err"
	model "github.com/LucasJandrey5/url-shortener/backend/src/model/user"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	return nil
}
