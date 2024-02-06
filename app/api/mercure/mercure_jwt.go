package mercure

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT() (string, error) {
    claims := jwt.MapClaims{
        "mercure": map[string]interface{}{
            "publish": []string{"*"}, // à modifier pour la sécurité de l'app
        },
        "exp": time.Now().Add(time.Hour * 1).Unix(), //exp du token
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtKey := "8c1f0ba78cabd4ea856a4945f9ec94f7ddbaecca3315ab52d175c3a47415dcfb"
    tokenString, err := token.SignedString([]byte(jwtKey))

    return tokenString, err
}
