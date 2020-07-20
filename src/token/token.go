package token

import (
	"limakcv/src/app/model"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var mySignedKey = []byte("mysupersecretphrase")

func GenerateToken(person *model.Person) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = person.IdentityID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignedKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func ValidateToken(next RequestHandlerFunction) RequestHandlerFunction {
	return func(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySignedKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				next(db, w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}

	}
}
