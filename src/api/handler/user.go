package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/constant"
	"github.com/naeemaei/golang-clean-web-api/dependency"
	"github.com/naeemaei/golang-clean-web-api/usecase"
)

type UsersHandler struct {
	userUsecase  *usecase.UserUsecase
	otpUsecase   *usecase.OtpUsecase
	tokenUsecase *usecase.TokenUsecase
	config       *config.Config
}

func NewUserHandler(cfg *config.Config) *UsersHandler {
	userUsecase := usecase.NewUserUsecase(cfg, dependency.GetUserRepository(cfg))
	otpUsecase := usecase.NewOtpUsecase(cfg)
	tokenUsecase := usecase.NewTokenUsecase(cfg)
	return &UsersHandler{userUsecase: userUsecase, otpUsecase: otpUsecase, tokenUsecase: tokenUsecase, config: cfg}
}

// LoginByUsername godoc
// @Summary LoginByUsername
// @Description LoginByUsername
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login-by-username [post]
func (h *UsersHandler) LoginByUsername(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	token, err := h.userUsecase.LoginByUsername(c, req.Username, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// Set the refresh token in a cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     constant.RefreshTokenCookieName,
		Value:    token.RefreshToken,
		MaxAge:   int(h.config.JWT.RefreshTokenExpireDuration * 60),
		Path:     "/",
		Domain:   h.config.Server.Domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}

// RegisterByUsername godoc
// @Summary RegisterByUsername
// @Description RegisterByUsername
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/register-by-username [post]
func (h *UsersHandler) RegisterByUsername(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	err = h.userUsecase.RegisterByUsername(c, req.ToRegisterUserByUsername())
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// RegisterLoginByMobileNumber godoc
// @Summary RegisterLoginByMobileNumber
// @Description RegisterLoginByMobileNumber
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterLoginByMobileRequest true "RegisterLoginByMobileRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login-by-mobile [post]
func (h *UsersHandler) RegisterLoginByMobileNumber(c *gin.Context) {
	req := new(dto.RegisterLoginByMobileRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	token, err := h.userUsecase.RegisterAndLoginByMobileNumber(c, req.MobileNumber, req.Otp)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// Set the refresh token in a cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     constant.RefreshTokenCookieName,
		Value:    token.RefreshToken,
		MaxAge:   int(h.config.JWT.RefreshTokenExpireDuration * 60),
		Path:     "/",
		Domain:   h.config.Server.Domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UsersHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	err = h.otpUsecase.SendOtp(req.MobileNumber)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	// TODO: Call internal SMS service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// RefreshToken godoc
// @Summary RefreshToken
// @Description RefreshToken
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 401 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/refresh-token [post]
func (h *UsersHandler) RefreshToken(c *gin.Context) {
	token, err := h.tokenUsecase.RefreshToken(c)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	// Set the refresh token in a cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     constant.RefreshTokenCookieName,
		Value:    token.RefreshToken,
		MaxAge:   int(h.config.JWT.RefreshTokenExpireDuration * 60),
		Path:     "/",
		Domain:   h.config.Server.Domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(token, true, helper.Success))
}
