package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		verb := c.Request.Method
		time := time.Now()
		path := c.Request.RequestURI

		//Process request
		c.Next()

		var size int
		if c.Writer != nil {
			size = c.Writer.Size()
		}

		fmt.Printf("time: %v\npath: %s\nverb: %s\nsize: %d", time, path, verb, size)
	}
}
