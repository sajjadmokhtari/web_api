package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constants"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"GOLANG_CLEAN_WEB_API/src/pkg/service_errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type TokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	UserName     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *TokenService) GenerateToken(token *TokenDto) (*dto.TokenDetail, error) {
	fmt.Println("[TokenService] شروع ساخت توکن برای کاربر:", token.UserId)

	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	// Access Token Claims
	atc := jwt.MapClaims{}
	atc[constants.UserIdKey] = token.UserId
	atc[constants.FirstNameKey] = token.FirstName
	atc[constants.LastNameKey] = token.LastName
	atc[constants.UserNameKey] = token.UserName
	atc[constants.EmailKey] = token.Email
	atc[constants.MobileNumberKey] = token.MobileNumber
	atc[constants.RoleKey] = token.Roles
	atc[constants.ExpireTimeKey] = td.AccessTokenExpireTime

	fmt.Println("[AccessToken Claims]")
	for k, v := range atc {
		fmt.Printf("  %s: %v\n", k, v)
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		fmt.Println("[Error] امضای AccessToken شکست خورد:", err)
		return nil, err
	}
	fmt.Println("[AccessToken] ساخته شد:", td.AccessToken)

	// Refresh Token Claims
	rtc := jwt.MapClaims{}
	rtc[constants.UserIdKey] = token.UserId
	rtc[constants.ExpireTimeKey] = td.RefreshTokenExpireTime

	fmt.Println("[RefreshToken Claims]")
	for k, v := range rtc {
		fmt.Printf("  %s: %v\n", k, v)
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)
	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.JWT.RefreshSecret))
	if err != nil {
		fmt.Println("[Error] امضای RefreshToken شکست خورد:", err)
		return nil, err
	}
	fmt.Println("[RefreshToken] ساخته شد:", td.RefreshToken)

	fmt.Printf("[TokenService] توکن‌ها ساخته شدند برای UserId=%d\n", token.UserId)
	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &service_errors.ServiceError{
				EndUserMessage: service_errors.UnExpectedError,
			}
		}
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}
	verifyToken, err := s.VerifyToken(token)
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
