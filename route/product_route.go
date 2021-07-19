package route

import (
	middleware "github.com/firmanJS/boilerplate-gin/middleware"
	handlerCreateProduct "github.com/firmanJS/boilerplate-gin/transport/http/product/create"
	createProduct "github.com/firmanJS/boilerplate-gin/usecase/product/create"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProductRoutes(db *gorm.DB, route *gin.Engine) {
	/**
	  @description All Handler Product
	*/
	createProductRepository := createProduct.NewRepositoryCreate(db)
	createProductService := createProduct.NewServiceCreate(createProductRepository)
	createProductHandler := handlerCreateProduct.NewHandlerCreateProduct(createProductService)
	/**
	  @description All Product Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/product", createProductHandler.CreateProductHandler)
}
