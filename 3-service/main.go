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

	// resourceTypeController --start
	resourceTypeRepository repository.ResourceTypeRepository = repository.NewResourceTypeRepository(db)
	resourceTypeService    service.ResourceTypeService       = service.NewResourceTypeRepository(resourceTypeRepository)
	resourceTypeController controller.ResourceTypeController = controller.NewResourceTypeController(resourceTypeService)
	// resourceController -- start
	resourceRepository repository.ResourceRepository = repository.NewResourceRepository(db)
	resourceService    service.ResourceService       = service.NewResourceService(resourceRepository)
	resourceController controller.ResourceController = controller.NewResourceController(resourceService)
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

	resourceType := route.Group("/api/v1", middleware.AuthorizeJWT(jwtService))

	{
		resourceType.GET("/resource_type", resourceTypeController.ResourceList)
		resourceType.POST("/resource_type", resourceTypeController.CreateResourceType)
		resourceType.PUT("/resource_type/:id", resourceTypeController.UpdateResourceType)
		resourceType.GET("/resource_type/:id", resourceTypeController.FindResourceType)
		resourceType.DELETE("/resource_type/:id", resourceTypeController.DeleteResourceType)
	}

	resource := route.Group("/api/v1", middleware.AuthorizeJWT(jwtService))
	{
		resource.GET("/resource", resourceController.GetAll)
		resource.POST("/resource", resourceController.Create)
		resource.PUT("/resource/:id", resourceController.Update)
		resource.GET("/resource/:id", resourceController.Get)
		resource.DELETE("/resource/:id", resourceController.Delete)
	}

	route.Run(":8080")

}
