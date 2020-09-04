package googlejwtchecker

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims

	ClientID      string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func (c *Claims) IsValidIssuer(issuers []string) bool {
	found := false
	for _, issuer := range issuers {
		if issuer == c.Issuer {
			found = true
			break
		}
	}

	return !found
}
