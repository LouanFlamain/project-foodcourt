package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"foodcourt/app/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("votre-clef-secrète")

type UserClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(userID int, username, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &UserClaims{
		Id:       userID,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*model.UserItem, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		user := &model.UserItem{
			Id:       claims.Id,
			Username: claims.Username,
			Email:    claims.Email,
		}
		return user, nil
	}

	return nil, errors.New("Invalid token")
}

func AuthenticateUser(db *sql.DB, email, password string) (*model.UserItem, error) {
	var user model.UserItem
	err := db.QueryRow("SELECT id, username, email, password, picture, roles FROM users WHERE email = ?", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Roles)
	if err != nil {
		return nil, fmt.Errorf("User not found: %v", err)
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("Invalid password")
	}

	return &user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
