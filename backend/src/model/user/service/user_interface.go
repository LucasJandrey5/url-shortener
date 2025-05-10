package service

import (
	"github.com/LucasJandrey5/url-shortener/backend/src/configuration/rest_err"
	model "github.com/LucasJandrey5/url-shortener/backend/src/model/user"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	FindUserByEmail(string) (*model.UserDomainInterface, *rest_err.RestErr)
	FindUserByUsername(string) (*model.UserDomainInterface, *rest_err.RestErr)
}
