package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	SetEmail(string)
	GetPassword() string
	SetPassword(string)
	GetUsername() string
	SetUsername(string)

	EncryptPassword()
}

type userDomain struct {
	email    string
	password string
	username string
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) SetEmail(email string) {
	ud.email = email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) SetPassword(password string) {
	ud.password = password
}

func (ud *userDomain) GetUsername() string {
	return ud.username
}

func (ud *userDomain) SetUsername(username string) {
	ud.username = username
}

func NewUserDomain(email, password, username string) *userDomain {
	return &userDomain{
		email,
		password,
		username,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
