package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllStatues(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	statues := []model.Status{}
	db.Find(&statues)
	respondJSON(w, http.StatusOK, statues)
}

func CreateStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	status := model.Status{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&status); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&status).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, status)
}

func GetStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["StatusID"])
	if err != nil {
		return
	}
	status := getStatusOr404(db, id, w, r)
	if status == nil {
		return
	}
	respondJSON(w, http.StatusOK, status)
}

func UpdateStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["StatusID"])
	if err != nil {
		return
	}
	status := getStatusOr404(db, id, w, r)
	if status == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&status); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&status).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, status)
}

func DeleteStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["StatusID"])
	if err != nil {
		return
	}
	status := getStatusOr404(db, id, w, r)
	if status == nil {
		return
	}
	if err := db.Delete(&status).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getStatusOr404(db *gorm.DB, StatusID int, w http.ResponseWriter, r *http.Request) *model.Status {
	status := model.Status{}
	if err := db.First(&status, model.Status{Model: gorm.Model{ID: uint(StatusID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &status
}
