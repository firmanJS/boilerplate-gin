package registerhandler

import (
	"net/http"

	"github.com/firmanJS/boilerplate-gin/model"
	"github.com/firmanJS/boilerplate-gin/usecase/auth/register"
	"github.com/firmanJS/boilerplate-gin/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	var input model.InputRegister
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	} else {

		resultRegister, errRegister := h.service.RegisterService(&input)

		switch errRegister {

		case util.CONFLICT:
			util.APIResponse(ctx, http.StatusConflict, "Username already exist", nil)
			return

		case util.FAILED:
			util.APIResponse(ctx, http.StatusForbidden, "Register new account failed", nil)
			return

		default:
			var secretKey = viper.GetString("GO_JWT_SECRET")

			accessTokenData := map[string]interface{}{"id": resultRegister.ID, "username": resultRegister.Username}
			_, errToken := util.Sign(accessTokenData, secretKey, 60)

			if errToken != nil {
				defer logrus.Error(errToken.Error())
				util.APIResponse(ctx, http.StatusBadRequest, "Generate accessToken failed", nil)
				return
			}

			util.APIResponse(ctx, http.StatusCreated, "Register new account successfully", input)
		}
	}
}
