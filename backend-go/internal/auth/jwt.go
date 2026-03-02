package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(userID uint64) (string, error){
	signingKey := []byte("JWT_SECRET_KEY")

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer: "TeamUP",
		},
		
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("Sign token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, signingKey []byte) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString, 
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
		
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return signingKey, nil
	},
)

if err != nil {
	return nil, err
}


claims, ok := token.Claims.(*Claims) 
if !ok {
	return nil, fmt.Errorf("invalid token")

}
return claims, nil
}