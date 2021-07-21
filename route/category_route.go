package route

import (
	middleware "github.com/firmanJS/boilerplate-gin/middleware"
	handlerReadCategory "github.com/firmanJS/boilerplate-gin/transport/http/category/read"
	readCategory "github.com/firmanJS/boilerplate-gin/usecase/category/read"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCategoryRoutes(db *gorm.DB, route *gin.Engine) {
	/**
	  @description All Handler Category
	*/
	readCategoryRepository := readCategory.NewRepositoryRead(db)
	readCategoryService := readCategory.NewServiceRead(readCategoryRepository)
	readCategoryHandler := handlerReadCategory.NewHandlerReadCategory(readCategoryService)
	/**
	  @description All Category Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.GET("/category", readCategoryHandler.ReadCategoryHandler)
}
