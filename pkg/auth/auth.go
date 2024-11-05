package auth

import (
	"context"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type clientHTTP struct {
	secretKey []byte
}

type Authenticator interface {
	ValidateToken(ctx context.Context, signedToken string) (*AuthorizedUser, error)
}

func NewClientHTTP(secretKey []byte) Authenticator {
	return &clientHTTP{
		secretKey: secretKey,
	}
}

type AuthorizedUser struct {
	UserID        string
	Username      string
	UserPackageID uint
	jwt.RegisteredClaims
}

func (c *clientHTTP) GenerateToken(args *AuthorizedUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":      args.Username,
			"userID":        args.UserID,
			"userPackageID": args.UserPackageID,
			"exp":           time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(c.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (c *clientHTTP) ValidateToken(ctx context.Context, tokenString string) (*AuthorizedUser, error) {
	claims := &AuthorizedUser{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return c.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Pastikan token valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
