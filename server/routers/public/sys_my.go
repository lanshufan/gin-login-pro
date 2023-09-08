package public

import (
	"github.com/gin-gonic/gin"
	"pass-crypto/api/v1/system"
)

type MyRouter struct {
}

func (m *MyRouter) BlogLists(group *gin.RouterGroup) {
	var page system.List

	group.POST("my", page.SysBlogList)
}
