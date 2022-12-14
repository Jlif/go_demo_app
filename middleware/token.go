package middleware

import (
	"github.com/gin-gonic/gin"
	"jiangchen.tech/demo_app/config"
	"jiangchen.tech/demo_app/utils"
	"net/http"
)

// CheckToken 验证token
func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 Authorization header 头
		tokenString := ctx.GetHeader("Authorization")
		// 验证token非空
		if tokenString == "" {
			utils.Response(ctx, http.StatusUnauthorized, 401, "未设置token", nil)
			ctx.Abort()
			return
		}
		// token验证是否失效
		token, claims, err := config.ParseToken(tokenString)
		if err != nil || !token.Valid {
			utils.Response(ctx, http.StatusUnauthorized, 401, "Token已过期", nil)
			ctx.Abort()
			return
		}
		//如果用户存在 将user信息存入上下文
		ctx.Set("userID", claims.UserId)
		ctx.Next()
	}
}
