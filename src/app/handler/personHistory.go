package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//GetAllPersonHistories ...
func GetAllPersonHistories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	personHistories := []model.PersonHistory{}
	db.Find(&personHistories)
	respondJSON(w, http.StatusOK, personHistories)
}

//CreatePersonHistory ...
func CreatePersonHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	personHistory := model.PersonHistory{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&personHistory); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&personHistory).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, personHistory)
}

//GetPersonHistory ...
func GetPersonHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonHistoryID"])
	if err != nil {
		return
	}
	personHistory := getPersonHistoryOr404(db, id, w, r)
	if personHistory == nil {
		return
	}
	respondJSON(w, http.StatusOK, personHistory)
}

//UpdatePersonHistory ...
func UpdatePersonHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonHistoryID"])
	if err != nil {
		return
	}
	personHistory := getPersonHistoryOr404(db, id, w, r)
	if personHistory == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&personHistory); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&personHistory).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, personHistory)
}

//DeletePersonHistory ...
func DeletePersonHistory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonHistoryID"])
	if err != nil {
		return
	}
	personHistory := getPersonHistoryOr404(db, id, w, r)
	if personHistory == nil {
		return
	}
	if err := db.Delete(&personHistory).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getPersonHistoryOr404(db *gorm.DB, PersonHistoryID int, w http.ResponseWriter, r *http.Request) *model.PersonHistory {
	personHistory := model.PersonHistory{}
	if err := db.First(&personHistory, model.PersonHistory{Model: gorm.Model{ID: uint(PersonHistoryID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &personHistory
}
