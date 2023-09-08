package modules

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type MyCustomClaims struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// GeneratePass 生成密码加密字符串
func GeneratePass(pass []byte) []byte {
	b, err := bcrypt.GenerateFromPassword(pass, 4)
	if err != nil {
		fmt.Println("生成密码失败: ", err)
		return nil
	}

	return b
}

// ComparePass 验证密码
func ComparePass(hashPass, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPass, pass)
	if err != nil {
		return false
	}

	return true
}

// GenerateToken 根据用户名和ID生成Token
func GenerateToken(id int, name string) []byte {
	// 固定的字符串
	mySigningKey := []byte("LToken")

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		id,
		name,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 生成时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
			Issuer:    "L",                                                // 发布者
			Subject:   "LSystem",                                          // 接收者
			ID:        "1",
			Audience:  []string{"LUser"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("生成token失败， ", err)
		return nil
	}

	return []byte(ss)
}

// ValidToken 返回用户ID
func ValidToken(tokenStr string) int {
	mySigningKey := []byte("LToken")

	token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if cliams, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		// token验证生效，返回用户ID
		return cliams.Id
	} else {
		return 0
	}
}
