package route

import (
	handlerLogin "github.com/firmanJS/boilerplate-gin/transport/http/auth/login"
	loginAuth "github.com/firmanJS/boilerplate-gin/usecase/auth/login"
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

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
