package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pass-crypto/modules"
)

type Login struct {
}

// LoginValid 用户登录验证
func (l *Login) LoginValid(ctx *gin.Context) {
	var reqBody modules.UserInfo

	if err := ctx.ShouldBind(&reqBody); err != nil {
		modules.ErrWithMsg("参数错误", ctx)
		return
	}

	// 账号查询是否存在
	var user modules.SystemUser
	res := modules.DB.Where("username = ?", reqBody.UserName).First(&user)
	if res.Error != nil {
		fmt.Println("账号不存在")
		modules.ErrOtWithMsg("登录失败，账号或密码不正确", ctx)
		return
	}

	// 查询密码是否正确
	if !modules.ComparePass([]byte(user.Password), []byte(reqBody.Password)) {
		modules.ErrOtWithMsg("登录失败，账号或密码不正确", ctx)
		fmt.Println("密码验证失败")
		return
	}

	// 验证成功，生成token
	b := modules.GenerateToken(user.Id, reqBody.UserName)
	var data = map[string]string{"LToken": string(b)}

	modules.OKWithAllResp(data, "登录成功", ctx)
}
