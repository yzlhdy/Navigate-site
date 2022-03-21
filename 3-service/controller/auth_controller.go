package controller

import (
	"navigate/dto"
	"navigate/entity"
	"navigate/helper"
	"navigate/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthController(authService service.AuthService, jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (a *authController) Login(c *gin.Context) {
	var loginForm dto.LoginDto
	c.ShouldBind(&loginForm)
	valid, err := helper.BindAndValid(c, &loginForm)
	if !valid {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		c.JSON(http.StatusOK, response)
		return
	}
	res := a.authService.VerifyCredential(loginForm.Email, loginForm.Password)
	if v, ok := res.(entity.User); ok {
		genToken := a.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = genToken

		response := helper.BuildResponse(200, "登录成功", v)
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := helper.BuildErrorResponse(401, "账号或密码错误", nil, helper.EmptyObjectResponse{})
		c.JSON(http.StatusOK, response)
	}
}

func (a *authController) Register(c *gin.Context) {
	var registerForm dto.RegisterDto
	c.ShouldBind(&registerForm)
	valid, err := helper.BindAndValid(c, &registerForm)
	if !valid {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		c.JSON(http.StatusOK, response)
		return
	}
	if !a.authService.IsDuplicateEmail(registerForm.Email) {
		response := helper.BuildErrorResponse(401, "邮箱已被注册", nil, helper.EmptyObjectResponse{})
		c.JSON(http.StatusOK, response)
		return
	} else {
		user := a.authService.CreateUser(registerForm)
		genToken := a.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))
		user.Token = genToken
		response := helper.BuildResponse(200, "注册成功", user)
		c.JSON(http.StatusOK, response)
		return
	}

}
