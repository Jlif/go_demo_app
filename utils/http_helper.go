package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, msg string, data any) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

func Success(ctx *gin.Context, msg string, data any) {
	Response(ctx, http.StatusOK, 200, msg, data)
}

func Fail(ctx *gin.Context, msg string, data any) {
	Response(ctx, http.StatusOK, 400, msg, data)
}
