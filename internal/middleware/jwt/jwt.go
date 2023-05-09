package jwt

import (
	"fmt"
	"net/http"
	"os"
	"post/pkg/app"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			appG  = app.Gin{C: c}
			token = strings.Replace(appG.C.Request.Header.Get("Authorization"), "Bearer ", "", 1)
			err   = ValidateToken(token)
		)

		if err != nil {
			appG.Response(http.StatusUnauthorized, "Please check your credential and try again.")
			appG.C.Abort()
			return
		}
		appG.C.Next()
	}
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil && strings.ToLower(err.Error()) == jwt.ErrTokenUsedBeforeIssued.Error() {
		return nil
	}
	return err
}
