package main

import (
	_ "fmt"

	"gorm.io/gorm"

	"github.com/Lasang3012/facilityManage/controller"
	"github.com/Lasang3012/facilityManage/middleware"
	"github.com/Lasang3012/facilityManage/repository"
	"github.com/Lasang3012/facilityManage/service"
	util "github.com/Lasang3012/facilityManage/util"
	"github.com/gin-gonic/gin"
)

var (
	db             *gorm.DB                  = util.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	roleRepository repository.RoleRepository = repository.NewRoleRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	roleService    service.RoleService       = service.NewRoleService(roleRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	roleController controller.RoleController = controller.NewRoleController(roleService)
)

func main() {
	defer util.CloseDatabaseConnection(db)

	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := router.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		// userRoutes.PUT("/profile", userController.Update)
	}

	roleRoutes := router.Group("api/auth")
	{
		roleRoutes.GET("/getUsers", roleController.All)
	}

	router.Run(":8088")

}
