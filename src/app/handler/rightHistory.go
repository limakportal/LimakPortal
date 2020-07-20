package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllRightHistoryes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rightHistoryes := []model.RightHistory{}
	db.Find(&rightHistoryes)
	respondJSON(w, http.StatusOK, rightHistoryes)
}

func CreateRightHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rightHistory := model.RightHistory{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rightHistory); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rightHistory).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, rightHistory)
}

func GetRightHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightHistoryID"])
	if err != nil {
		return
	}
	rightHistory := getRightHistoryOr404(db, id, w, r)
	if rightHistory == nil {
		return
	}
	respondJSON(w, http.StatusOK, rightHistory)
}

func UpdateRightHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightHistoryID"])
	if err != nil {
		return
	}
	rightHistory := getRightHistoryOr404(db, id, w, r)
	if rightHistory == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rightHistory); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rightHistory).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rightHistory)
}

func DeleteRightHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightHistoryID"])
	if err != nil {
		return
	}
	rightHistory := getRightHistoryOr404(db, id, w, r)
	if rightHistory == nil {
		return
	}
	if err := db.Delete(&rightHistory).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getRightHistoryOr404(db *gorm.DB, rightHistoryID int, w http.ResponseWriter, r *http.Request) *model.RightHistory {
	rightHistory := model.RightHistory{}
	if err := db.First(&rightHistory, model.RightHistory{Model: gorm.Model{ID: uint(rightHistoryID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &rightHistory
}
