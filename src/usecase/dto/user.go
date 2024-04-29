package dto

import model "github.com/naeemaei/golang-clean-web-api/domain/model"

type TokenDetail struct {
	AccessToken            string
	RefreshToken           string
	AccessTokenExpireTime  int64
	RefreshTokenExpireTime int64
}

type RegisterUserByUsername struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func ToUserModel(from RegisterUserByUsername) model.User {
	return model.User{Username: from.Username,
		FirstName: from.FirstName,
		LastName:  from.LastName,
		Email:     from.Email,
	}
}

type RegisterLoginByMobile struct {
	MobileNumber string
	Otp          string
}

type LoginByUsername struct {
	Username string
	Password string
}
