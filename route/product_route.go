package route

import (
	middleware "github.com/firmanJS/boilerplate-gin/middleware"
	handlerCreateProduct "github.com/firmanJS/boilerplate-gin/transport/http/product/create"
	handlerDeleteProduct "github.com/firmanJS/boilerplate-gin/transport/http/product/delete"
	handlerReadProduct "github.com/firmanJS/boilerplate-gin/transport/http/product/read"
	handlerUpdateProduct "github.com/firmanJS/boilerplate-gin/transport/http/product/update"
	createProduct "github.com/firmanJS/boilerplate-gin/usecase/product/create"
	deleteProduct "github.com/firmanJS/boilerplate-gin/usecase/product/delete"
	readProduct "github.com/firmanJS/boilerplate-gin/usecase/product/read"
	updateProduct "github.com/firmanJS/boilerplate-gin/usecase/product/update"
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

	readProductRepository := readProduct.NewRepositoryRead(db)
	readProductService := readProduct.NewServiceRead(readProductRepository)
	readProductHandler := handlerReadProduct.NewHandlerReadProduct(readProductService)

	updateProductRepository := updateProduct.NewRepositoryUpdate(db)
	updateProductService := updateProduct.NewServiceUpdate(updateProductRepository)
	updateProductHandler := handlerUpdateProduct.NewHandlerUpdateProduct(updateProductService)

	deleteProductRepository := deleteProduct.NewRepositoryDelete(db)
	deleteProductService := deleteProduct.NewServiceDelete(deleteProductRepository)
	deleteProductHandler := handlerDeleteProduct.NewHandlerDeleteProduct(deleteProductService)
	/**
	  @description All Product Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/product", createProductHandler.CreateProductHandler)
	groupRoute.GET("/product", readProductHandler.ReadProductHandler)
	groupRoute.PUT("/product/:id", updateProductHandler.UpdateProductHandler)
	groupRoute.DELETE("/product/:id", deleteProductHandler.DeleteProductHandler)
}
