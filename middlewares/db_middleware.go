package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DBMiddleware is a middleware function that injects the database connection into the context
func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
		c.Next()
	}
}
