package graph

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/marklude/go-sovtech/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	Results *model.Results
}

type UserInfo struct {
	Username string
}

type CustomClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func generateKeyPair(bits int) *rsa.PrivateKey {
	// This method requires a random number of bits.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return privateKey
}
