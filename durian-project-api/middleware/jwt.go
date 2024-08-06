package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bryantang1107/DuriYum/dbutils"
	"github.com/bryantang1107/DuriYum/models"
	"github.com/bryantang1107/DuriYum/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		utils.HandleErrorResponse(c, http.StatusUnauthorized, "Please login again", err)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		utils.HandleErrorResponse(c, http.StatusUnauthorized, "Please login again", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp, ok := claims["exp"].(float64)

		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			utils.HandleErrorResponse(c, http.StatusUnauthorized, "Please login again", err)
			return
		}
		if time.Now().Unix() > int64(exp) {
			c.AbortWithStatus(http.StatusUnauthorized)
			utils.HandleErrorResponse(c, http.StatusUnauthorized, "Please login again", err)
			return
		}
		user_id := c.Query("userId")

		conditions := map[string]interface{}{
			"id": user_id,
		}
		var user models.User
		err := dbutils.SelectOne(&user, conditions)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			utils.HandleErrorResponse(c, http.StatusUnauthorized, "User Profile Not Found", err)
			return
		}

		c.Set("user", user.Id)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		utils.HandleErrorResponse(c, http.StatusUnauthorized, "Please login again", err)
		return
	}
}
