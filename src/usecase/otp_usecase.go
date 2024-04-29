package usecase

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/naeemaei/golang-clean-web-api/common"
	"github.com/naeemaei/golang-clean-web-api/config"
	constant "github.com/naeemaei/golang-clean-web-api/constant"
	"github.com/naeemaei/golang-clean-web-api/infra/cache"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	"github.com/naeemaei/golang-clean-web-api/pkg/service_errors"
)

type OtpUsecase struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type otpDto struct {
	Value string
	Used  bool
}

func NewOtpUsecase(cfg *config.Config) *OtpUsecase {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpUsecase{logger: logger, cfg: cfg, redisClient: redis}
}

func (u *OtpUsecase) SendOtp(mobileNumber string) error {
	otp := common.GenerateOtp()
	err := u.SetOtp(mobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpUsecase) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constant.RedisOtpDefaultKey, mobileNumber)
	val := &otpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[otpDto](s.redisClient, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OptExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}
	err = cache.Set(s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpUsecase) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constant.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[otpDto](s.redisClient, key)
	if err != nil {
		return err
	} else if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
