package token

import (
	"time"

	"github.com/daymenu/shipping/user/model"

	pb "github.com/daymenu/shipping/user/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("adecfkijknl()*&")
)

// CustomClaims CustomClaims 结构体
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

// Authable 定义接口
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

// TokenService 定义token结构体
type TokenService struct {
	User *model.UserModel
}

// Decode 解码
func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if clasims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return clasims, nil
	}
	return nil, err
}

// Encode 编码
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "shippy.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
