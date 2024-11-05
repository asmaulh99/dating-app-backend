package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware untuk memeriksa JWT di header Authorization
func Authorize(authCl Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		// Ambil token dengan menghapus prefix "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validasi token
		claims, err := authCl.ValidateToken(ctx, tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// Simpan klaim di context untuk digunakan di endpoint yang dilindungi
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
