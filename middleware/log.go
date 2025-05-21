package middleware

import (
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// 日志中间件，记录请求和响应
func LogRequestResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		clientIP := c.ClientIP()

		var reqBody []byte
		if c.Request.Body != nil {
			reqBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(strings.NewReader(string(reqBody)))
		}

		log.Printf("[REQUEST] %s %s?%s from %s body: %s", method, path, query, clientIP, string(reqBody))

		w := &bodyWriter{body: &strings.Builder{}, ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		status := c.Writer.Status()
		log.Printf("[RESPONSE] %s %s status: %d body: %s", method, path, status, w.body.String())
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	body *strings.Builder
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
