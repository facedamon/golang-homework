package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

func CheckPasswd(pwd string, hashedPasswd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(pwd)) == nil
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedToken, err
}

func ParseJWT(token string) (string, error) {
	if len(token) > 7 && token[:7] == "Bearer " {
		tokenString := token[7:]
		t, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method, expected HMAC")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			return "", err
		}
		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			username, ok := claims["username"].(string)
			if !ok {
				return "", errors.New("username claim is not a string")
			}
			return username, nil
		}
	}
	return "", errors.New("非法token")
}
