package helper

import (
	"Auth/pkg/utils/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authCustomClaimsJobSeeker struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func JobSeekerPasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashPassword)
	return hash, nil
}

func GenerateTokenJobSeeker(jobSeeker models.JobSeekerDetailsResponse) (string, error) {
	claims := &authCustomClaimsJobSeeker{
		Email: jobSeeker.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("jobseekerkey"))
	if err != nil {
		fmt.Println("Error is", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenJobSeeker(tokenString string) (*authCustomClaimsJobSeeker, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsJobSeeker{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("jobseekerkey"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsJobSeeker); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
