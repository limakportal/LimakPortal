package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllGenders(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	genders := []model.Gender{}
	db.Find(&genders)
	respondJSON(w, http.StatusOK, genders)
}

func CreateGender(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	gender := model.Gender{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&gender); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&gender).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, gender)
}

func GetGender(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["GenderID"])
	if err != nil {
		return
	}
	gender := getGenderOr404(db, id, w, r)
	if gender == nil {
		return
	}
	respondJSON(w, http.StatusOK, gender)
}

func UpdateGender(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["GenderID"])
	if err != nil {
		return
	}
	gender := getGenderOr404(db, id, w, r)
	if gender == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&gender); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&gender).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, gender)
}

func DeleteGender(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["GenderID"])
	if err != nil {
		return
	}
	gender := getGenderOr404(db, id, w, r)
	if gender == nil {
		return
	}
	if err := db.Delete(&gender).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getGenderOr404(db *gorm.DB, genderID int, w http.ResponseWriter, r *http.Request) *model.Gender {
	gender := model.Gender{}
	if err := db.First(&gender, model.Gender{Model: gorm.Model{ID: uint(genderID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &gender
}
