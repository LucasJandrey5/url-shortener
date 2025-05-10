package service

import (
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/logger"
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/rest_err"
	model "github.com/LucasJandrey5/url-shortener/backend/src/model/user"
	"go.uber.org/zap"
)

func (*userDomainService) FindUserByEmail(email string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail model", zap.String("journey", "findUserByEmail"))

	return nil, nil
}

func (*userDomainService) FindUserByUsername(username string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByUsername model", zap.String("journey", "findUserByUsername"))

	return nil, nil
}
