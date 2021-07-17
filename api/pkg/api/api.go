package api

import (
	"some/api/pkg/api/code"
	"some/api/pkg/api/response"
	"github.com/gin-gonic/gin"
)

func SetResponse(c *gin.Context, httpCode, responseCode int, data interface{}) {
	base := response.NewResponse(c)
	base.SetResponse(httpCode, responseCode, data)

	if responseCode != code.Success {
		c.Abort()
	}

	WriterAccessLog()
}

// 记录api访问日志
func WriterAccessLog() {
	return
}
