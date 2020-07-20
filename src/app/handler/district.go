package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllDistricts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	districts := []model.District{}
	db.Find(&districts)
	respondJSON(w, http.StatusOK, districts)
}

func CreateDistrict(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	district := model.District{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&district); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&district).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, district)
}

func GetDistrict(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["DistrictID"])
	if err != nil {
		return
	}
	district := getDistrictOr404(db, id, w, r)
	if district == nil {
		return
	}
	respondJSON(w, http.StatusOK, district)
}

func UpdateDistrict(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["DistrictID"])
	if err != nil {
		return
	}
	district := getDistrictOr404(db, id, w, r)
	if district == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&district); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&district).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, district)
}

func DeleteDistrict(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["DistrictID"])
	if err != nil {
		return
	}
	district := getDistrictOr404(db, id, w, r)
	if district == nil {
		return
	}
	if err := db.Delete(&district).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getDistrictOr404(db *gorm.DB, districtID int, w http.ResponseWriter, r *http.Request) *model.District {
	district := model.District{}
	if err := db.First(&district, model.District{Model: gorm.Model{ID: uint(districtID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &district
}
