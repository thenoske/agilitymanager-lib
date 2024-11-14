package agilitymanager_lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
)

// LicenceClaims this structure hold the data that has been saved in token
type LicenceClaims struct {

	// standard jwt token claims
	jwt.StandardClaims

	// Additional fields
	Customer string `json:"customer"`
	Code     string `json:"code"`
	Type     string `json:"type"`
	MaxTeams int    `json:"maxTeams"`
}

func (l LicenceClaims) ToJson() []byte {
	b, err := json.Marshal(&l)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}

	return b
}

// VerifyToken verify authenticate token
func VerifyToken(token string) (claims *LicenceClaims, err error) {

	var (
		t *jwt.Token
	)

	// parse with claims
	if t, err = jwt.ParseWithClaims(token, &LicenceClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return false, fmt.Errorf("%w: %v",
					errors.New("invalid token"), token.Header["alg"])
			}
			return SECURE, nil
		}); err != nil {

		return
	}

	// if token is valid
	if t == nil {
		return nil, errors.New("invalid token")
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	return t.Claims.(*LicenceClaims), nil
}
