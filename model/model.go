package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type RequestUser struct {
	LoginID string `json:"loginid"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Passwd  string `json:"passwd"`
	Country string `json:"country"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	JwtData map[string]string `json:"JwtData"`
}
