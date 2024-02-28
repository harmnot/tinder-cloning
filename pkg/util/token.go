package util

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type AccountDataClaims struct {
	AccountID string `json:"account_id"`
}

func GenerateTokenJWT(accountId string) (string, error) {
	// Set expired token to 8 hours
	expiredToken := time.Now().Add(time.Hour * 8).Unix()

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id": accountId,
		"exp":        expiredToken,
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(jwtKey)
}

func VerifyTokenJWT(tokenString string) (AccountDataClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtKey, nil
	})
	if err != nil {
		return AccountDataClaims{}, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountData := AccountDataClaims{
			AccountID: claims["account_id"].(string),
		}
		return accountData, nil
	} else {
		return AccountDataClaims{}, err
	}
}

func GetClaimsFromContext(ctx context.Context) (AccountDataClaims, error) {
	claims, ok := ctx.Value("accountData").(AccountDataClaims)
	if !ok {
		return AccountDataClaims{}, errors.New("invalid Payload Token")
	}
	return claims, nil
}
