package handlerDeleteProduct

import (
	"net/http"

	structs "github.com/firmanJS/boilerplate-gin/usecase/product"
	deleteProduct "github.com/firmanJS/boilerplate-gin/usecase/product/delete"
	util "github.com/firmanJS/boilerplate-gin/util"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service deleteProduct.Service
}

func NewHandlerDeleteProduct(service deleteProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteProductHandler(ctx *gin.Context) {

	var input structs.InputDeleteProduct
	input.Id = ctx.Param("id")

	_, err := h.service.DeleteProductService(&input)

	switch err.Code {

	case util.NOT_FOUND:
		util.APIResponse(ctx, http.StatusNotFound, err.Message, nil)
		return

	case util.FAILED:
		util.APIResponse(ctx, http.StatusForbidden, err.Message, nil)
		return

	default:
		util.APIResponse(ctx, http.StatusOK, err.Message, nil)
	}
}
