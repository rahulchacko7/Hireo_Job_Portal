package helper

import (
	"Auth/pkg/utils/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authCustomClaimsEmployer struct {
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	jwt.StandardClaims
}

func EmployerPasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashPassword)
	return hash, nil
}

func GenerateTokenEmployer(employer models.EmployerDetailsResponse) (string, error) {
	claims := &authCustomClaimsEmployer{
		CompanyName: employer.CompanyName,
		Email:       employer.ContactEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your_signing_key")) // Update with your signing key
	if err != nil {
		fmt.Println("Error is", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenEmployer(tokenString string) (*authCustomClaimsEmployer, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsEmployer{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your_signing_key"), nil // Update with your signing key
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsEmployer); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
