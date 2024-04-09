package middleware

import (
	"net/http"
	"strings"
	"thelastking/kingseafood/model"
	"thelastking/kingseafood/pkg/security"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 || fields[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := fields[1]

		token, err := jwt.ParseWithClaims(tokenString, &model.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(security.JWT_KEY), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(*model.AccessToken)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func RefreshJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 || fields[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := fields[1]

		token, err := jwt.ParseWithClaims(tokenString, &model.RefreshToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(security.JWT_KEY), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(*model.RefreshToken)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token has expired"})
			return
		}
		users := &model.Users{
			UserID: claims.UserID,
			Role:   claims.Role,
		}

		// Generate a new Access Token
		accessToken, err := security.GenerateAccessToken(users) // Assuming GenerateAccessToken expects a string user ID
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Access Token"})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Header("Authorization", "Bearer "+accessToken)

		c.Next()
	}
}
