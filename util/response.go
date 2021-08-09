package util

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(ctx *gin.Context, StatusCode int, Message string, Data interface{}) {

	jsonResponse := Responses{
		Message: Message,
		Data:    Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

type CatchError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
