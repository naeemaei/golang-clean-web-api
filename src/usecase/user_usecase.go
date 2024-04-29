package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/common"
	"github.com/naeemaei/golang-clean-web-api/config"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	"github.com/naeemaei/golang-clean-web-api/pkg/service_errors"
	dto "github.com/naeemaei/golang-clean-web-api/usecase/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	logger       logging.Logger
	cfg          *config.Config
	otpUsecase   *OtpUsecase
	tokenUsecase *TokenUsecase
	repository   repository.UserRepository
}

func NewUserUsecase(cfg *config.Config, repository repository.UserRepository) *UserUsecase {
	logger := logging.NewLogger(cfg)
	return &UserUsecase{
		cfg:          cfg,
		repository:   repository,
		logger:       logger,
		otpUsecase:   NewOtpUsecase(cfg),
		tokenUsecase: NewTokenUsecase(cfg),
	}
}

// Login by username
func (u *UserUsecase) LoginByUsername(ctx context.Context, username string, password string) (*dto.TokenDetail, error) {
	user, err := u.repository.FetchUserInfo(ctx, username, password)

	if err != nil {
		return nil, err
	}
	tokenDto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, ur.Role.Name)
		}
	}

	token, err := u.tokenUsecase.GenerateToken(tokenDto)

	if err != nil {
		return nil, err
	}
	return token, nil

}

// Register by username
func (u *UserUsecase) RegisterByUsername(ctx context.Context, req dto.RegisterUserByUsername) error {
	user := dto.ToUserModel(req)

	exists, err := u.repository.ExistsEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = u.repository.ExistsUsername(ctx, req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	user.Password = string(hp)
	_, err = u.repository.CreateUser(ctx, user)
	return err

}

// Register/login by mobile number
func (u *UserUsecase) RegisterAndLoginByMobileNumber(ctx context.Context, mobileNumber string, otp string) (*dto.TokenDetail, error) {
	err := u.otpUsecase.ValidateOtp(mobileNumber, otp)
	if err != nil {
		return nil, err
	}
	exists, err := u.repository.ExistsMobileNumber(ctx, mobileNumber)
	if err != nil {
		return nil, err
	}

	user := model.User{MobileNumber: mobileNumber, Username: mobileNumber}

	if exists {
		user, err = u.repository.FetchUserInfo(ctx, user.Username, user.Password)
		if err != nil {
			return nil, err
		}

		return u.generateToken(user)
	}

	// Register and login
	bp := []byte(common.GeneratePassword())
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return nil, err
	}
	user.Password = string(hp)

	user, err = u.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return u.generateToken(user)

}

func (u *UserUsecase) generateToken(user model.User) (*dto.TokenDetail, error) {
	tokenDto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, ur.Role.Name)
		}
	}

	token, err := u.tokenUsecase.GenerateToken(tokenDto)
	if err != nil {
		return nil, err
	}
	return token, nil
}
