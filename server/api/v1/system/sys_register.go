package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pass-crypto/modules"
	"time"
)

type Register struct {
}

// UserRegister 用户注册
func (r *Register) UserRegister(ctx *gin.Context) {
	var reqBody modules.UserInfo

	if err := ctx.ShouldBind(&reqBody); err != nil {
		modules.ErrWithMsg("参数错误", ctx)
		return
	}
	fmt.Println(reqBody)

	// 账号查询是否存在
	var user modules.SystemUser
	res := modules.DB.Where("username = ?", reqBody.UserName).First(&user)
	if res.Error == nil {
		fmt.Println("账号已存在")
		modules.ErrOtWithMsg("账号已存在", ctx)
		return
	}

	// 加密密码字符串
	cryptoPass := modules.GeneratePass([]byte(reqBody.Password))

	// 创建账号
	var userCreate = modules.SystemUser{
		UserName:   reqBody.UserName,
		Password:   string(cryptoPass),
		UserStatus: 0,
		CreateAt:   modules.XTime(time.Now()),
	}
	resCreate := modules.DB.Create(&userCreate)
	if resCreate.Error != nil {
		fmt.Println("创建账号失败, ", resCreate.Error)
		modules.ErrWithMsg("账号注册失败，请联系管理员", ctx)
		return
	}

	// jwt token
	b := modules.GenerateToken(userCreate.Id, reqBody.UserName)
	var data = map[string]string{"LToken": string(b)}

	modules.OKWithAllResp(data, "注册成功", ctx)
}
