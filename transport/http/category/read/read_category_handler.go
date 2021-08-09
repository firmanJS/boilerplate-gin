package handlerReadCategory

import (
	"net/http"

	readCategory "github.com/firmanJS/boilerplate-gin/usecase/category/read"
	util "github.com/firmanJS/boilerplate-gin/util"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service readCategory.Service
}

func NewHandlerReadCategory(service readCategory.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ReadCategoryHandler(ctx *gin.Context) {

	result, err := h.service.ReadCategoryService()

	switch err.Code {

	case util.NOT_FOUND:
		util.APIResponse(ctx, http.StatusNotFound, "data is empty", nil)

	default:
		util.APIResponse(ctx, http.StatusOK, "sucess", result)
	}
}
