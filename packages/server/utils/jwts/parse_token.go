package jwts

import (
	"errors"
	"fmt"
	"web/global"

	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
  MySecret = []byte(global.Config.Jwt.Secret)
  token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
    return MySecret, nil
  })
  if err != nil {
    global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
    // fmt.Println(token)
    return nil, err
  }
  if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
    return claims, nil
  }
  return nil, errors.New("invalid token")
}
