package public

import (
	"github.com/gin-gonic/gin"
	"pass-crypto/api/v1/system"
)

type RegisterRouter struct {
}

func (r *RegisterRouter) RegisterRouter(group *gin.RouterGroup) {
	var register system.Register
	group.POST("register", register.UserRegister)
}
