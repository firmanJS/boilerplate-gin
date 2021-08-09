package handlerReadProduct

import (
	"net/http"

	readProduct "github.com/firmanJS/boilerplate-gin/usecase/product/read"
	util "github.com/firmanJS/boilerplate-gin/util"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service readProduct.Service
}

func NewHandlerReadProduct(service readProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ReadProductHandler(ctx *gin.Context) {

	result, err := h.service.ReadProductService()

	switch err.Code {

	case util.NOT_FOUND:
		util.APIResponse(ctx, http.StatusNotFound, "data is empty", nil)

	default:
		util.APIResponse(ctx, http.StatusOK, "sucess", result)
	}
}
