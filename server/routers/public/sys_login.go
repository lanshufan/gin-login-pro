package public

import (
	"github.com/gin-gonic/gin"
	"pass-crypto/api/v1/system"
)

type LoginRouter struct {
}

func (l *LoginRouter) LoginRouter(group *gin.RouterGroup) {
	var login system.Login
	group.POST("login", login.LoginValid)
}
