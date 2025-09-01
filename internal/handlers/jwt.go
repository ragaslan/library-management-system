package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UID  uint   `json:"uid"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (h *Handler) signAccess(uid uint, role string) (string, time.Time, error) {
	exp := time.Now().Add(h.accessTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UID:  uid,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	s, err := token.SignedString([]byte(h.jwtSecret))

	return s, exp, err
}
