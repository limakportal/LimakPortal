package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllMaritalStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	maritalStatus := []model.MaritalStatus{}
	db.Find(&maritalStatus)
	respondJSON(w, http.StatusOK, maritalStatus)
}

func CreateMaritalStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	maritalStatus := model.MaritalStatus{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&maritalStatus); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&maritalStatus).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, maritalStatus)
}

func GetMaritalStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["MaritalStatusID"])
	if err != nil {
		return
	}
	maritalStatus := getMaritalStatusOr404(db, id, w, r)
	if maritalStatus == nil {
		return
	}
	respondJSON(w, http.StatusOK, maritalStatus)
}

func UpdateMaritalStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["MaritalStatusID"])
	if err != nil {
		return
	}
	maritalStatus := getMaritalStatusOr404(db, id, w, r)
	if maritalStatus == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&maritalStatus); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&maritalStatus).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, maritalStatus)
}

func DeleteMaritalStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["MaritalStatusID"])
	if err != nil {
		return
	}
	maritalStatus := getMaritalStatusOr404(db, id, w, r)
	if maritalStatus == nil {
		return
	}
	if err := db.Delete(&maritalStatus).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getMaritalStatusOr404(db *gorm.DB, MaritalStatusID int, w http.ResponseWriter, r *http.Request) *model.MaritalStatus {
	maritalStatus := model.MaritalStatus{}
	if err := db.First(&maritalStatus, model.MaritalStatus{Model: gorm.Model{ID: uint(MaritalStatusID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &maritalStatus
}
