package system

import (
	"github.com/gin-gonic/gin"
	"pass-crypto/modules"
)

type List struct {
}
type ProductList struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Date   string `json:"date"`
}

func (l *List) SysBlogList(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	// 分析token获取用户id
	id := modules.ValidToken(tokenStr)
	// 无效权限
	if id == 0 {
		modules.ErrWithMsg("用户权限无效，请联系管理员", ctx)
		return
	}

	list := []ProductList{}

	// test
	//fmt.Println("用户id：", id)

	// 管理员用户权限
	if id == 1 {
		list = []ProductList{
			{
				Id:     1,
				Name:   "Java实践",
				Author: "xxk",
				Date:   "2023-01-01",
			},
			{
				Id:     2,
				Name:   "Golang深入浅出",
				Author: "xxk",
				Date:   "2023-05-01",
			},
			{
				Id:     3,
				Name:   "Python从入门到放弃",
				Author: "xxk",
				Date:   "2023-06-01",
			},
			{
				Id:     4,
				Name:   "H5 学习深入",
				Author: "xxk",
				Date:   "2023-01-03",
			},
			{
				Id:     5,
				Name:   "JavaScript基础学习",
				Author: "xxk",
				Date:   "2023-01-01",
			},
		}
	} else {
		// 普通权限
		list = []ProductList{
			{
				Id:     4,
				Name:   "H5 学习深入",
				Author: "xxk",
				Date:   "2023-01-03",
			},
			{
				Id:     5,
				Name:   "JavaScript基础学习",
				Author: "xxk",
				Date:   "2023-01-01",
			},
		}
	}

	modules.OKWithAllResp(list, "请求成功", ctx)
}
