package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func getTimeStamp() int64 {
	return time.Now().Unix()
}

//TokenClaim token 加密的结构体
type TokenClaim struct {
	ID        string
	Username  string
	Timestamp int64
}

//CreateToken 获取token
func CreateToken(tokenClaim TokenClaim, secret string) (token string, err error) {

	claim := jwt.MapClaims{
		"id":        tokenClaim.ID,
		"username":  tokenClaim.Username,
		"timestamp": tokenClaim.Timestamp,
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = tok.SignedString([]byte(secret))
	return
}

//ParseToken .
func ParseToken(token, secret string) (tokenClaim TokenClaim, err error) {
	keyfunc := func() jwt.Keyfunc {
		return func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}
	}

	tok, err := jwt.Parse(token, keyfunc())
	if err != nil {
		return
	}

	claim, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}

	if !tok.Valid {
		err = errors.New("token is invalid")
		return
	}

	tokenClaim.ID = claim["id"].(string)
	tokenClaim.Username = claim["username"].(string)
	tokenClaim.Timestamp = int64(claim["timestamp"].(float64))
	return
}
