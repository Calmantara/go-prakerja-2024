package middleware

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/Calmantara/go-prakerja-2024/deploy/helper"
	"github.com/gin-gonic/gin"
)

const (
	PASSWORD = "mysupersecretpassword"
	USERNAME = "myusername"
)

// Perbedaan Authentication v Authorization
// Authentication => public / kalau sudah login
// Authorization => mengetahui roles yang diperbolehkan untuk mengakses suatu resource

// Basic Authentication
func BasicAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		// Basic bXl1c2VybmFtZTpteXN1cGVyc2VjcmV0cGFzc3dvcmQ=
		// encoding base64
		// Bearer

		// get the encoded string
		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization header",
			})
			return
		}

		// check basic
		if splitToken[0] != "Basic" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization method",
			})
			return
		}
		// decode base64
		var decodedHeader []byte
		decodedHeader, err := base64.StdEncoding.DecodeString(splitToken[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization token",
			})
			return
		}
		splitDecodedToken := strings.Split(string(decodedHeader), ":")
		if len(splitDecodedToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "malformed token",
			})
			return
		}
		username := splitDecodedToken[0]
		password := splitDecodedToken[1]

		if username != USERNAME || password != PASSWORD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid username or password",
			})
			return
		}
		log.Println("BasicAuthorization")
		// penting untuk middleware
		ctx.Next()
	}
}

func BearerAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		// {Authorization: Bearer jwt_token}
		// get the encoded string
		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization header",
			})
			return
		}

		// check basic
		if splitToken[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization method",
			})
			return
		}
		// validate jwt
		valid := helper.ValidateUserJWT(splitToken[1])
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "malformed token",
			})
			return
		}
		ctx.Next()
	}
}
