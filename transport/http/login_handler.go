package http

import (
	"net/http"

	"github.com/firmanJS/boilerplate-gin/model"
	"github.com/firmanJS/boilerplate-gin/util"
	"github.com/firmanJS/boilerplate-gin/usecase/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type handlerLogin struct {
	service auth.ServiceLogin
}

func NewHandlerLogin(service auth.ServiceLogin) *handlerLogin {
	return &handlerLogin{service: service}
}

func (h *handlerLogin) LoginHandler(ctx *gin.Context) {

	var input model.InputLogin
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	} else {

		resultLogin, errLogin := h.service.LoginService(&input)

		switch errLogin {

		case util.NOT_FOUND:
			util.APIResponse(ctx, http.StatusNotFound, "User account is not registered", nil)
			return

		case util.FAILED:
			util.APIResponse(ctx, http.StatusForbidden, "Username or password is wrong", nil)
			return

		default:
			accessTokenData := map[string]interface{}{"id": resultLogin.ID, "username": resultLogin.Username}
			accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

			if errToken != nil {
				defer logrus.Error(errToken.Error())
				util.APIResponse(ctx, http.StatusBadRequest, "Generate accessToken failed", nil)
				return
			}

			responseToken := map[string]string{"accessToken": accessToken}

			util.APIResponse(ctx, http.StatusOK, "Login successfully", responseToken)
		}
	}
}
