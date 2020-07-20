package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllRightStatuses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rightStatuses := []model.RightStatus{}
	db.Find(&rightStatuses)
	respondJSON(w, http.StatusOK, rightStatuses)
}

func CreateRightStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rightStatus := model.RightStatus{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rightStatus); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rightStatus).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, rightStatus)
}

func GetRightStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightStatusID"])
	if err != nil {
		return
	}
	rightStatus := getRightStatusOr404(db, id, w, r)
	if rightStatus == nil {
		return
	}
	respondJSON(w, http.StatusOK, rightStatus)
}

func UpdateRightStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightStatusID"])
	if err != nil {
		return
	}
	rightStatus := getRightStatusOr404(db, id, w, r)
	if rightStatus == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rightStatus); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rightStatus).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rightStatus)
}

func DeleteRightStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightStatusID"])
	if err != nil {
		return
	}
	rightStatus := getRightStatusOr404(db, id, w, r)
	if rightStatus == nil {
		return
	}
	if err := db.Delete(&rightStatus).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getRightStatusOr404(db *gorm.DB, rightStatusID int, w http.ResponseWriter, r *http.Request) *model.RightStatus {
	rightStatus := model.RightStatus{}
	if err := db.First(&rightStatus, model.RightStatus{Model: gorm.Model{ID: uint(rightStatusID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &rightStatus
}
