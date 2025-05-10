package service

import (
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/logger"
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(userID string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
