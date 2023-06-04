package util

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CachedToken struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}

type CustomClaim struct {
	UserID   string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UID      string `json:"uid"`
	jwt.RegisteredClaims
}

type tokenResult struct {
	token string
	expat *jwt.NumericDate
	uid   string
}

type jwtResult struct {
	accessToken  string
	refreshToken string
	expired_at   *jwt.NumericDate
}

func GenerateJwt(userID, roleCode, issuer string) jwtResult {
	access := createToken(userID, roleCode, issuer, "secretAccess", 1)
	refresh := createToken(userID, roleCode, issuer, "secretRefresh", 6)
	cachedJson, err := json.Marshal(CachedToken{
		AccessUID:  access.uid,
		RefreshUID: refresh.uid,
	})
	PanicIfNeeded(err)
	fmt.Println(cachedJson)
	return jwtResult{
		accessToken:  access.token,
		refreshToken: refresh.token,
		expired_at:   access.expat,
	}
}

func createToken(userID, roleCode, issuer, secret string, expHour int16) tokenResult {
	uid := uuid.NewString()
	expat := jwt.NewNumericDate(time.Now().Add(time.Duration(expHour) * time.Hour))
	claims := CustomClaim{
		userID,
		roleCode,
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: expat,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}
	mySigningKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	ss, err := token.SignedString(mySigningKey)
	PanicIfNeeded(err)
	return tokenResult{
		token: ss,
		expat: expat,
		uid:   uid,
	}
}