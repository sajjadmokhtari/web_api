package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/common"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constants"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"

	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"GOLANG_CLEAN_WEB_API/src/pkg/service_errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	OtpService   *OtpService
	database     *gorm.DB
	tokenService *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:          cfg,
		database:     database,
		logger:       logger,
		OtpService:   NewOtpService(cfg),
		tokenService: NewTokenService(cfg),
	}
}

// Login by username
func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User

	err := s.database.
		Preload("UserRoles.Role").
		Where("user_name = ?", req.Username).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	tdto := TokenDto{
		UserId:       user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		UserName:     user.UserName,
	}
	fmt.Println("مقدار  یوزر نیم", user.UserName)

	for _, ur := range user.UserRoles {
		tdto.Roles = append(tdto.Roles, ur.Role.Name)
	}

	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Register by username
func (s *UserService) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{
		UserName:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if exists, err := s.existsByEmail(req.Email); err != nil {
		return err
	} else if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}

	if exists, err := s.existsByUsername(req.Username); err != nil {
		return err
	} else if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(common.GeneratePassword()), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)

	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	tx := s.database.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Register and login by mobile
func (s *UserService) RegisterLoginByMobileRequest(req *dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {
	fmt.Println("[Start] شروع ثبت و لاگین با موبایل")

	if err := s.OtpService.ValidateOtp(req.MobileNumber, req.Otp); err != nil {
		fmt.Println("[Error] OTP نامعتبر:", err)
		return nil, err
	}
	fmt.Println("[OTP] تایید موفق OTP برای شماره:", req.MobileNumber)

	exists, err := s.existsByMobileNumber(req.MobileNumber)
	if err != nil {
		fmt.Println("[Error] بررسی وجود کاربر شکست خورد:", err)
		return nil, err
	}
	fmt.Println("[UserCheck] آیا کاربر وجود دارد؟", exists)

	u := models.User{
		MobileNumber: req.MobileNumber,
		UserName:     req.MobileNumber,
	}

	if exists {
		fmt.Println("[Flow] ورود با کاربر موجود")
		var user models.User
		err = s.database.
			Preload("UserRoles.Role").
			Where("mobile_number = ?", u.MobileNumber).
			First(&user).Error
		if err != nil {
			fmt.Println("[Error] دریافت اطلاعات کاربر موجود:", err)
			return nil, err
		}
		fmt.Println("[FindUser] کاربر یافت شد:", user.Id)

		tdto := TokenDto{
			UserId:       user.Id,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			MobileNumber: user.MobileNumber,
			UserName: user.UserName,
		}

		for _, ur := range user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
		fmt.Printf("[Debug-Tdto] UserId: %d\n", tdto.UserId)
		fmt.Printf("[Debug-Tdto] FirstName: %s\n", tdto.FirstName)
		fmt.Printf("[Debug-Tdto] LastName: %s\n", tdto.LastName)
		fmt.Printf("[Debug-Tdto] Email: %s\n", tdto.Email)
		fmt.Printf("[Debug-Tdto] MobileNumber: %s\n", tdto.MobileNumber)
		fmt.Printf("[Debug-Tdto] Roles: %v\n", tdto.Roles)

		fmt.Println("[Token] توکن تولید شد برای کاربر موجود:", tdto.UserId)
		return s.tokenService.GenerateToken(&tdto)
	}

	fmt.Println("[Flow] ثبت کاربر جدید")

	hp, err := bcrypt.GenerateFromPassword([]byte(common.GeneratePassword()), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("[Error] ساخت رمز عبور شکست خورد:", err)
		return nil, err
	}
	u.Password = string(hp)

	roleId, err := s.getDefaultRole()
	if err != nil {
		fmt.Println("[Error] گرفتن نقش پیش‌فرض شکست خورد:", err)
		return nil, err
	}
	fmt.Println("[Role] نقش پیش‌فرض دریافت شد:", roleId)

	tx := s.database.Begin()
	fmt.Println("[Transaction] شروع transaction")

	if err := tx.Create(&u).Error; err != nil {
		fmt.Println("[Error] ثبت کاربر شکست خورد:", err)
		tx.Rollback()
		return nil, err
	}
	fmt.Println("[User] کاربر جدید ثبت شد با شناسه:", u.Id)

	if err := tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error; err != nil {
		fmt.Println("[Error] ثبت نقش برای کاربر شکست خورد:", err)
		tx.Rollback()
		return nil, err
	}
	fmt.Println("[UserRole] نقش کاربر ثبت شد")

	tx.Commit()
	fmt.Println("[Transaction] commit انجام شد")

	var user models.User
	err = s.database.
		Preload("UserRoles.Role").
		Where("user_name = ?", u.UserName).
		First(&user).Error
	if err != nil {
		fmt.Println("[Error] دریافت مجدد کاربر شکست خورد:", err)
		return nil, err
	}

	tdto := TokenDto{
		UserId:       user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
	}

	for _, ur := range user.UserRoles {
		tdto.Roles = append(tdto.Roles, ur.Role.Name)
	}

	fmt.Println("[Success] توکن نهایی تولید شد برای کاربر:", user.Id)
	return s.tokenService.GenerateToken(&tdto)
}

// SendOtp: OTP تولید و ذخیره
func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	return s.OtpService.SetOtp(req.MobileNumber, otp)
}

// Database checks
func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).Error
	return exists, err
}

func (s *UserService) existsByUsername(userName string) (bool, error) {
	var exists bool
	err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", userName).
		Find(&exists).Error
	return exists, err
}

func (s *UserService) existsByMobileNumber(mobileNumber string) (bool, error) {
	var exists bool
	err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).Error
	return exists, err
}

func (s *UserService) getDefaultRole() (int, error) {
	var role models.Role
	roleName := constants.DefaultRoleName

	err := s.database.
		Model(&models.Role{}).
		Where("name = ?", roleName).
		First(&role).Error

	if err != nil {
		s.logger.Error(logging.General, "DefaultRoleNotFound", fmt.Sprintf("role '%s' not found: %v", roleName, err), nil)
		return 0, fmt.Errorf("default role '%s' not found: %w", roleName, err)
	}

	s.logger.Info(logging.General, "DefaultRoleResolved", fmt.Sprintf("resolved role '%s' with id: %d", roleName, role.Id), nil)
	return role.Id, nil
}
