package helpers

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "4ndaS4ngatK3poD3nganK3yS4y4Yaw"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))
	return signedToken
}

func ValidateToken(tokenString string) (id uint, email string, err error) {
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("sign in to process")
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return 0, "", fmt.Errorf("token invalid")
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	intID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))

	id = uint(intID)
	email = fmt.Sprintf("%v", claims["email"])

	return id, email, nil
}
