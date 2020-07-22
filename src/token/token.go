package token

import (
	"encoding/json"
	"fmt"
	"limakcv/src/app/model"
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
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()

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
				respondErrorToken(w, http.StatusUnauthorized, "Token expired")
			}
			if token.Valid {
				next(db, w, r)
			}

		} else {
			respondErrorToken(w, http.StatusUnauthorized, "Not Authorized")
		}

	}
}

func respondErrorToken(w http.ResponseWriter, code int, message string) {
	respondJSONToken(w, code, map[string]string{"error": message})
}

func respondJSONToken(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
