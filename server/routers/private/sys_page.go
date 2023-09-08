package private

import (
	"github.com/gin-gonic/gin"
	"pass-crypto/api/v1/system"
)

type PageRouter struct {
}

func (p *PageRouter) PageRouter(group *gin.RouterGroup) {
	var page system.List

	group.POST("/list", page.SysBlogList)
}
