package main

import (
	"navigate/config"
	"navigate/controller"
	"navigate/middleware"
	"navigate/repository"
	"navigate/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetUpDatabaseConnection()
	// authController --start
	jwtService     service.JwtService        = service.NewJwtService()
	responseAuth   repository.UserRepository = repository.NewUserRepository(db)
	authService    service.AuthService       = service.NewAuthService(responseAuth)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	// authController --end
)

func main() {

	defer config.TearDownDatabaseConnection(db)
	route := gin.New()
	// route.Use(middleware.Loggers())      // 日志
	route.Use(middleware.Cors())         // 跨域
	route.Use(gin.Recovery())            // 异常处理
	route.Use(middleware.Translations()) // 语言
	// 路由
	authRoutes := route.Group("/api/v1")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	route.Run(":8080")

}
