package service

import (
	"fmt"
	"navigate/dto"
	"navigate/entity"
	"navigate/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	// 登录
	VerifyCredential(email string, password string) interface{}
	// 注册
	CreateUser(user dto.RegisterDto) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func comparePassword(hash string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), plainPassword)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (a *authService) VerifyCredential(email string, password string) interface{} {
	res := a.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		compare := comparePassword(v.Password, []byte(password))
		if compare && v.Email == email {
			return v
		} else {
			return false
		}
	}
	return false
}

func (a *authService) CreateUser(user dto.RegisterDto) entity.User {
	return a.userRepository.InsertUser(entity.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Avatar:   user.Avatar,
	})

}
func (a *authService) IsDuplicateEmail(email string) bool {
	res := a.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
