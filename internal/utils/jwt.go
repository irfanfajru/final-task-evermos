package utils

import (
	"fmt"
	"time"
	"tugas_akhir_example/internal/daos"

	jtoken "github.com/golang-jwt/jwt/v4"
)

// @TODO : make function create jwt token and validate
func CreateToken(payload daos.User, secret string) string {
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"UserId":  fmt.Sprintf("%v", payload.ID),
		"IsAdmin": payload.IsAdmin,
		"exp":     time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte(secret))
	// Return the token
	return t
}
