package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"pass-crypto/modules"
)

// VerifyMiddleware jwt验证
func VerifyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		// Authorization 为空
		if tokenStr == "" {
			fmt.Println("Authorization is null")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"data": nil,
				"msg":  "Unauthorized",
			})

			ctx.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenStr, &modules.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("LToken"), nil
		})
		// token认证失败
		if _, ok := token.Claims.((*modules.MyCustomClaims)); !ok || !token.Valid || err != nil {
			fmt.Println(ok, err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"data": nil,
				"msg":  "Unauthorized",
			})
			ctx.Abort()
			return
		}

		origin := ctx.GetHeader("Origin")
		fmt.Printf("验证通过，%s\n", origin)
		ctx.Next()
	}
}
