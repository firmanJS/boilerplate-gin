package handlerUpdateProduct

import (
	"net/http"

	structs "github.com/firmanJS/boilerplate-gin/usecase/product"

	updateProduct "github.com/firmanJS/boilerplate-gin/usecase/product/update"
	util "github.com/firmanJS/boilerplate-gin/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service updateProduct.Service
}

func NewHandlerUpdateProduct(service updateProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateProductHandler(ctx *gin.Context) {

	var input structs.InputUpdateProduct
	input.Id = ctx.Param("id")
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	} else {

		_, err := h.service.UpdateProductService(&input)

		switch err.Code {

		case util.NOT_FOUND:
			util.APIResponse(ctx, http.StatusNotFound, err.Message, nil)

		case util.FAILED:
			util.APIResponse(ctx, http.StatusForbidden, err.Message, nil)

		default:
			util.APIResponse(ctx, http.StatusOK, err.Message, input)
		}
	}
}
