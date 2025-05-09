package usecase

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/constant"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	"github.com/naeemaei/golang-clean-web-api/pkg/service_errors"
	dto "github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type TokenUsecase struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenUsecase(cfg *config.Config) *TokenUsecase {
	logger := logging.NewLogger(cfg)
	return &TokenUsecase{
		cfg:    cfg,
		logger: logger,
	}
}

func (u *TokenUsecase) GenerateToken(token tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(u.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(u.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{}

	atc[constant.UserIdKey] = token.UserId
	atc[constant.FirstNameKey] = token.FirstName
	atc[constant.LastNameKey] = token.LastName
	atc[constant.UsernameKey] = token.Username
	atc[constant.EmailKey] = token.Email
	atc[constant.MobileNumberKey] = token.MobileNumber
	atc[constant.RolesKey] = token.Roles
	atc[constant.ExpireTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.AccessToken, err = at.SignedString([]byte(u.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[constant.UserIdKey] = token.UserId
	rtc[constant.FirstNameKey] = token.FirstName
	rtc[constant.LastNameKey] = token.LastName
	rtc[constant.UsernameKey] = token.Username
	rtc[constant.EmailKey] = token.Email
	rtc[constant.MobileNumberKey] = token.MobileNumber
	rtc[constant.RolesKey] = token.Roles
	rtc[constant.ExpireTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(u.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (u *TokenUsecase) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(u.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (u *TokenUsecase) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := u.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}

func (s *TokenUsecase) RefreshToken(c *gin.Context) (*dto.TokenDetail, error) {
	refreshToken, err := c.Cookie(constant.RefreshTokenCookieName)
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.InvalidRefreshToken}
	}

	claims, err := s.GetClaims(refreshToken)
	if err != nil {
		return nil, err
	}

	// Convert roles to []string
	rolesInterface, ok := claims[constant.RolesKey].([]interface{})
	if !ok {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.InvalidRolesFormat}
	}

	roles := make([]string, len(rolesInterface))
	for i, role := range rolesInterface {
		roles[i], ok = role.(string)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.InvalidRolesFormat}
		}
	}

	tokenDto := tokenDto{
		UserId:       int(claims[constant.UserIdKey].(float64)),
		FirstName:    claims[constant.FirstNameKey].(string),
		LastName:     claims[constant.LastNameKey].(string),
		Username:     claims[constant.UsernameKey].(string),
		MobileNumber: claims[constant.MobileNumberKey].(string),
		Email:        claims[constant.EmailKey].(string),
		Roles:        roles,
	}
	newTokenDetail, err := s.GenerateToken(tokenDto)
	if err != nil {
		return nil, err
	}

	return newTokenDetail, nil
}
