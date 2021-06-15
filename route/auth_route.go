package route

import (
	handlerLogin "github.com/firmanJS/boilerplate-gin/transport/http/auth/login"
	handlerRegister "github.com/firmanJS/boilerplate-gin/transport/http/auth/register"
	loginAuth "github.com/firmanJS/boilerplate-gin/usecase/auth/login"
	registerAuth "github.com/firmanJS/boilerplate-gin/usecase/auth/register"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	  @description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)
	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
