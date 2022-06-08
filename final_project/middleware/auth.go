package middleware

import (
	"final_project/helpers"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			helpers.Unauthorized(c, "missing 'Authorization' header")
			return
		}

		bearer := strings.HasPrefix(token, "Bearer")
		if !bearer {
			helpers.Unauthorized(c, "no 'Bearer'")
			return
		}

		tokenString := strings.Split(token, "Bearer ")[1]
		if tokenString == "" {
			helpers.Unauthorized(c, "token not found")
			return
		}

		claims, err := helpers.VerifyToken(tokenString)
		if err != nil {
			helpers.Unauthorized(c, err.Error())
			return
		}
		fmt.Println("====", claims.(jwt.Claims))

		var data = claims.(jwt.MapClaims)

		fmt.Println(data)

		c.Set("id", data["id"])
		c.Set("email", data["email"])
		c.Set("username", data["user_name"])
		c.Next()
	}
}
