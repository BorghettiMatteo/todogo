package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// uso questo per verificare lo status del login
	return func(c *gin.Context) {
		// faccio c.next se  effettivamente il token passato è valido, senno c.abort
		token := c.Query("token")
		if VerifyToken(token) {
			c.Next()
			c.IndentedJSON(http.StatusOK, gin.H{"message": "access granted"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "this token is not valid, please re-login / register"})
			c.Abort()
		}
	}
}
