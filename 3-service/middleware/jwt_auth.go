package middleware

import (
	"log"
	"navigate/helper"
	"navigate/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse(401, "Unauthorized", "请求未携带token，无权限访问", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		// token 验证
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			log.Println(err)
			response := helper.BuildErrorResponse(401, "Unauthorized", "token验证失败", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		// 获取用户id
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
