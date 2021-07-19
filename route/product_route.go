package route

import (
	middleware "github.com/firmanJS/boilerplate-gin/middleware"
	handlerCreateProduct "github.com/firmanJS/boilerplate-gin/transport/http/product/create"
	createProduct "github.com/firmanJS/boilerplate-gin/usecase/product/create"
    handlerDeleteProduct "github.com/firmanJS/boilerplate-gin/transport/http/product/delete"
	deleteProduct "github.com/firmanJS/boilerplate-gin/usecase/product/delete"
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
	
	deleteProductRepository := deleteProduct.NewRepositoryDelete(db)
	deleteProductService := deleteProduct.NewServiceDelete(deleteProductRepository)
	deleteProductHandler := handlerDeleteProduct.NewHandlerDeleteProduct(deleteProductService)
    /**
	  @description All Product Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/product", createProductHandler.CreateProductHandler)
    groupRoute.DELETE("/product/:id", deleteProductHandler.DeleteProductHandler)
}
