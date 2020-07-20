package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//GetAllPersonelInformations ...
func GetAllPersonelInformations(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	personalInformations := []model.PersonelInformation{}
	db.Find(&personalInformations)
	respondJSON(w, http.StatusOK, personalInformations)
}

//CreatePersonelInformation ...
func CreatePersonelInformation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	personalInformation := model.PersonelInformation{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&personalInformation); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	person := CheckPerson(db, personalInformation.PersonID, w, r)

	if person == nil {
		respondError(w, http.StatusNotFound, "Bu kullanıcı kayıtlı değil")
		return
	}
	if err := db.Save(&personalInformation).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, personalInformation)
}

//GetPersonelInformation ...
func GetPersonelInformation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	personalInformation := getPersonelInformationOr404(db, id, w, r)
	if personalInformation == nil {
		return
	}
	respondJSON(w, http.StatusOK, personalInformation)
}

//UpdatePersonelInformation ...
func UpdatePersonelInformation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	personalInformation := getPersonelInformationOr404(db, id, w, r)
	if personalInformation == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&personalInformation); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&personalInformation).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, personalInformation)
}

//DeletePersonelInformation ...
func DeletePersonelInformation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	personalInformation := getPersonelInformationOr404(db, id, w, r)
	if personalInformation == nil {
		return
	}
	if err := db.Delete(&personalInformation).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getPersonelInformationOr404(db *gorm.DB, PersonID int, w http.ResponseWriter, r *http.Request) *model.PersonelInformation {
	personalInformation := model.PersonelInformation{}
	if err := db.First(&personalInformation, "person_id=?", PersonID).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &personalInformation
}
