package jwt

import (
	"time"

	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/env"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id"`
}

func Create(userID uuid.UUID) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "hackathon-fiber-starter",
			Subject:   userID.String(),
			Audience:  jwt.ClaimStrings{"hackathon-fiber-starter"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(env.AppEnv.JwtExpTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
		UserID: userID,
	}

	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJWT, err := unsignedJWT.SignedString(env.AppEnv.JwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func Decode(tokenString string, claims *Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (any, error) {
		return env.AppEnv.JwtSecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	return nil
}
