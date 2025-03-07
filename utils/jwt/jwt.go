package jwt

import (
	model "Beckend_Student2025/models"
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyToken(raw string) (map[string]any, error) {
	godotenv.Load()
	token, err := jwt.Parse(raw, func(token *jwt.Token) (
		interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token singing method")
		}
		secret := []byte(os.Getenv("TOKEN_SECRET"))
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}

func GenerateTokenUser(ctx context.Context, user *model.Users) (string, error) {
	godotenv.Load()
	tokenDurationStr := os.Getenv("TOKEN_DURATION")
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"sub": jwt.MapClaims{
			"id":                user.ID,
			"firstname":         user.Firstname,
			"lastname":          user.Lastname,
			"nickname":          user.Nickname,
			"email":             user.Email,
			"password":          user.Password,
			"student_id":        user.StudentID,
			"faculty":           user.Faculty,
			"medical_condition": user.MedicalCondition,
			"food_allergies":    user.FoodAllergies,
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(tokenDuration).Unix(),
	})

	secret := []byte(os.Getenv("TOKEN_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	return tokenString, nil
}

func GenerateTokenAdmin(ctx context.Context, admin *model.Admins) (string, error) {
	godotenv.Load()
	tokenDurationStr := os.Getenv("TOKEN_DURATION")
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"sub": jwt.MapClaims{
			"id":       admin.ID,
			"username": admin.Username,
			"password": admin.Password,
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(tokenDuration).Unix(),
	})

	secret := []byte(os.Getenv("TOKEN_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	return tokenString, nil
}

func GenerateTokenStaff(ctx context.Context, staff *model.Staffs) (string, error) {
	godotenv.Load()
	tokenDurationStr := os.Getenv("TOKEN_DURATION")
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"sub": jwt.MapClaims{
			"id":       staff.ID,
			"username": staff.Username,
			"password": staff.Password,
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(tokenDuration).Unix(),
	})

	secret := []byte(os.Getenv("TOKEN_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	return tokenString, nil
}
