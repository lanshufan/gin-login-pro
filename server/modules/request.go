package modules

// UserInfo login and register user info
type UserInfo struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
