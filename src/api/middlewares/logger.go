package middlewares

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc{
	logger := logging.NewLogger(cfg)
	return  StructuredLogger(logger)
}

func StructuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "swagger"){
			c.Next()

		}else {
					blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		start := time.Now()
		path := c.FullPath()
		raw := c.Request.URL.RawQuery

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Writer = blw
		c.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		Keys := map[logging.ExtraKey]interface{}{}
		Keys[logging.Path] = param.Path
		Keys[logging.ClientIp] = param.ClientIP
		Keys[logging.Method] = param.Method
		Keys[logging.Latency] = param.Latency
		Keys[logging.StatusCode] = param.StatusCode
		Keys[logging.ErrorMessage] = param.ErrorMessage
		Keys[logging.BodySize] = param.BodySize
		Keys[logging.RequestBody] = string(bodyBytes)
		Keys[logging.ResponseBody] = blw.body.String()

		logger.Info(logging.RequestResponse, logging.Api, "", Keys)

		}


	}

}
