package jwt

import (
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
)

type Claimer struct {
	req struct {
		accountId int64
	}

	payload struct {
		AccountId int64 `json:"account_id"`
		jwt.StandardClaims
	}
}

func (c *Claimer) GetTokens(accountId, expire int64) string {
	c.payload.StandardClaims.Issuer = "AuthServices"
	c.payload.StandardClaims.ExpiresAt = expire
	c.payload.AccountId = accountId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.payload)

	tokenString, err := token.SignedString([]byte(JWTSecretString))
	ok := true
	if err != nil {
		logs.Error("[Controller:Security][CreateJWT] Cannot create token. Details: ", err.Error())
		ok = false
	}
	return tokenString, ok
}

func (c *Claimer) Validate(token string) bool {
	return false
}

func NewJWTClaimer() *Claimer {
	return &Claimer{}
}
