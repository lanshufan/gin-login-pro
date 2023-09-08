package modules

type SystemUser struct {
	Id         int    `json:"id"`
	UserName   string `json:"username" gorm:"column:username"`
	Password   string `json:"password"`
	UserStatus int    `json:"user_status"`
	CreateAt   XTime  `json:"create_at"`
}

func (s SystemUser) TableName() string {
	return "t_user"
}
