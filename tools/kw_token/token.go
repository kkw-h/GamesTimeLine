//Package kw_token
/*
@Title: token.go
@Description
@Author: kkw 2023/5/24 15:32
*/
package kw_token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.kkw.top/gamesTimeLine/internal/config"
	"go.kkw.top/gamesTimeLine/pkg"
	"time"
)

func GenerateToken(id int64, phone string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"phone": phone,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // 设置过期时间，这里设置为1小时
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义的密钥对令牌进行签名
	signedToken, _ := token.SignedString([]byte(config.Get("app.jwtSecretKey")))

	return signedToken
}

type TokenContent struct {
	Id    int64
	Phone string
}

func ValidateToken(tokenString string) (*TokenContent, error) {
	// 解析 JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名方法无效")
		}

		// 返回密钥
		return []byte(config.Get("app.jwtSecretKey")), nil
	})

	// 验证解析结果
	if err != nil {
		pkg.KwLogger.Debug("无效的Token", err)
		return nil, err
	}

	// 判断 token 是否有效
	if token.Valid {
		// 获取 JWT 中的声明信息
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			pkg.KwLogger.Debug("获取Token数据错误")
			return nil, errors.New("获取Token数据错误")
		}
		// 获取特定声明字段
		return &TokenContent{
			Id:    int64(claims["id"].(float64)),
			Phone: claims["phone"].(string),
		}, nil
	}
	return nil, errors.New("token无效")
}
