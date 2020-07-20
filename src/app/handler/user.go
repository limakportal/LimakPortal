package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"limakcv/src/token"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	person := CheckPerson(db, user.PersonID, w, r)

	if person == nil {
		respondError(w, http.StatusNotFound, "Bu kullanıcı kayıtlı değil")
		return
	}
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["UserID"])
	if err != nil {
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.First(&user, "Email = ? AND Password = ?", user.Email, user.Password).Error; err != nil {
		respondError(w, http.StatusNotFound, "Kullanıcı bulunamadı")
	} else {
		person := getPersonOr404(db, user.PersonID, w, r)

		if person != nil {
			tokenString, err := token.GenerateToken(person)
			if err != nil {
				respondError(w, http.StatusNotFound, err.Error())
				return
			}
			w.Header().Set("Token", tokenString)
			respondJSON(w, http.StatusOK, person)
		} else {
			respondError(w, http.StatusNotFound, "Personel bulunamadı")
		}
	}

}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["UserID"])
	if err != nil {
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["UserID"])
	if err != nil {
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getUserOr404(db *gorm.DB, userID int, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, "person_id=?", userID).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
