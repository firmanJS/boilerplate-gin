package handlerCreateProduct

import (
	"net/http"

	structs "github.com/firmanJS/boilerplate-gin/usecase/product"
	createProduct "github.com/firmanJS/boilerplate-gin/usecase/product/create"
	util "github.com/firmanJS/boilerplate-gin/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service createProduct.Service
}

func NewHandlerCreateProduct(service createProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateProductHandler(ctx *gin.Context) {

	var input structs.InputCreateProduct
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	} else {

		_, errCreateProduct := h.service.CreateProductService(&input)

		switch errCreateProduct {

		case util.CONFLICT:
			util.APIResponse(ctx, http.StatusConflict, "ID Product already exist", nil)
			return

		case util.FAILED:
			util.APIResponse(ctx, http.StatusForbidden, "Create new Product failed", nil)
			return

		default:
			util.APIResponse(ctx, http.StatusCreated, "Create new Product successfully", input)
		}
	}
}
