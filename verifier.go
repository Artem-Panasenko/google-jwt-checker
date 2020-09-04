package googlejwtchecker

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var (
	// Issuers is the allowed oauth token issuers
	Issuers = []string{
		"accounts.google.com",
		"https://accounts.google.com",
	}
)

func Verify(idToken string) (*jwt.Token, error) {
	return new(Verifier).Verify(idToken)
}

func VerifyRequest(r *http.Request) (*jwt.Token, error) {
	return new(Verifier).Verify(extractToken(r))
}

type Verifier struct{}

func (v *Verifier) Verify(idToken string) (*jwt.Token, error) {
	certs, err := getCerts()
	if err != nil {
		return nil, err
	}

	return verifyWithCerts(idToken, certs, Issuers)
}

func verifyWithCerts(tokenString string, certs *Certs, issuers []string) (*jwt.Token, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {

		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, errors.New("Wrong token header")
		}

		publicKey := certs.Keys[kid]
		if publicKey == nil {
			return nil, errors.New("No public key found for given kid")
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("ParseWithClaims Error: %s", err.Error())
	}

	if _, ok := token.Claims.(Claims); !ok && !token.Valid {
		return nil, errors.New("Invalid token claims")
	}

	if !claims.IsValidIssuer(issuers) {
		return nil, fmt.Errorf("Wrong issuer: %s", claims.Issuer)
	}

	return token, nil
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}
