package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllRightTypes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rightTypes := []model.RightType{}
	db.Find(&rightTypes)
	respondJSON(w, http.StatusOK, rightTypes)
}

func CreateRightType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rightType := model.RightType{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rightType); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rightType).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, rightType)
}

func GetRightType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightTypeID"])
	if err != nil {
		return
	}
	rightType := getRightTypeOr404(db, id, w, r)
	if rightType == nil {
		return
	}
	respondJSON(w, http.StatusOK, rightType)
}

func UpdateRightType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightTypeID"])
	if err != nil {
		return
	}
	rightType := getRightTypeOr404(db, id, w, r)
	if rightType == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rightType); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rightType).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rightType)
}

func DeleteRightType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightTypeID"])
	if err != nil {
		return
	}
	rightType := getRightTypeOr404(db, id, w, r)
	if rightType == nil {
		return
	}
	if err := db.Delete(&rightType).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getRightTypeOr404(db *gorm.DB, rightTypeID int, w http.ResponseWriter, r *http.Request) *model.RightType {
	rightType := model.RightType{}
	if err := db.First(&rightType, model.RightType{Model: gorm.Model{ID: uint(rightTypeID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &rightType
}
