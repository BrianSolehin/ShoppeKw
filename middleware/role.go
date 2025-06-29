package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// RequireRole mengecek apakah role user sesuai dengan yang diizinkan
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role tidak ditemukan"})
			c.Abort()
			return
		}

		// Konversi role ke string
		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Format role tidak valid"})
			c.Abort()
			return
		}

		// Cek apakah role user termasuk dalam allowedRoles
		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak: role tidak diizinkan"})
		c.Abort()
	}
}
