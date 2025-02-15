package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aidenfine/go_carmeet_backend/config"
	user_types "github.com/aidenfine/go_carmeet_backend/services/user/types"
	"github.com/aidenfine/go_carmeet_backend/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token`
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var payload RefreshTokenPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	refreshToken, err := VerifyRefreshToken(payload.RefreshToken)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid refresh token"))
		return
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token claims"))
		return
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid subject in token"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	newClaims := jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to generate new token"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"access_token": tokenString})
}

func VerifyRefreshToken(refreshToken string) (*jwt.Token, error) {
	// In a real scenario, you might verify against a database or other persistent storage
	// For now, assuming it's valid as a simple check.
	secret := []byte(config.Envs.JWTSecret)
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ExtractIDFromToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("auth header is missing or empty")
	}
	tokenString := strings.Split(authHeader, " ")[1]

	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Envs.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	log.Println(token, "token ")
	log.Println(claims, "claims ")

	userID, ok := (*claims)["userId"].(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in token")
	}

	return userID, nil
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user user_types.UserLoginPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Check credentials (this part should verify the email/username and password)
	// Here you can check user in the DB and verify password with the hashed password

	if !VerifyPassword(user.Password, "hashedPassword") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	accessToken := generateAccessToken(user.Email)

	refreshToken := generateRefreshToken(user.Email)

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func generateAccessToken(email string) string {
	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	secret := []byte(config.Envs.JWTSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func generateRefreshToken(email string) string {
	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	secret := []byte(config.Envs.JWTSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)

	return tokenString
}
