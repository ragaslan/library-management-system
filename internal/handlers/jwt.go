package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type JwtHandler struct {
	RDB        *redis.Client
	jwtSecret  string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

type Claims struct {
	UID  uint   `json:"uid"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (h *JwtHandler) signAccess(uid uint, role string) (string, time.Time, error) {
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

func (h *JwtHandler) hashPassword(pw string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(b), err
}

func (h *JwtHandler) checkPassword(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
