package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllCities(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	cities := []model.City{}
	db.Find(&cities)
	respondJSON(w, http.StatusOK, cities)
}

func CreateCity(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	city := model.City{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&city); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&city).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, city)
}

func GetCity(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["CityID"])
	if err != nil {
		return
	}
	city := getCityOr404(db, id, w, r)
	if city == nil {
		return
	}
	respondJSON(w, http.StatusOK, city)
}

func UpdateCity(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["CityID"])
	if err != nil {
		return
	}
	city := getCityOr404(db, id, w, r)
	if city == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&city); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&city).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, city)
}

func DeleteCity(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["CityID"])
	if err != nil {
		return
	}
	city := getCityOr404(db, id, w, r)
	if city == nil {
		return
	}
	if err := db.Delete(&city).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getCityOr404(db *gorm.DB, cityID int, w http.ResponseWriter, r *http.Request) *model.City {
	city := model.City{}
	if err := db.First(&city, model.City{Model: gorm.Model{ID: uint(cityID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &city
}
