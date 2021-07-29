package route

import (
	authHandler "github.com/firmanJS/boilerplate-gin/transport/http/auth_handler"
	"github.com/firmanJS/boilerplate-gin/usecase/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	  @description All Handler Auth
	*/
	LoginRepository := auth.NewRepositoryLogin(db)
	loginService := auth.NewServiceLogin(LoginRepository)
	loginHandler := authHandler.NewHandlerLogin(loginService)
	registerRepository := auth.NewRepositoryRegister(db)
	registerService := auth.NewServiceRegister(registerRepository)
	registerHandler := authHandler.NewHandlerRegister(registerService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
