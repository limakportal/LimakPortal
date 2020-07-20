package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllTitles(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	titles := []model.Title{}
	db.Find(&titles)
	respondJSON(w, http.StatusOK, titles)
}

func CreateTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	title := model.Title{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&title); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&title).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, title)
}

func GetTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["TitleID"])
	if err != nil {
		return
	}
	title := getTitleOr404(db, id, w, r)
	if title == nil {
		return
	}
	respondJSON(w, http.StatusOK, title)
}

func UpdateTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["TitleID"])
	if err != nil {
		return
	}
	title := getTitleOr404(db, id, w, r)
	if title == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&title); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&title).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, title)
}

func DeleteTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["TitleID"])
	if err != nil {
		return
	}
	title := getTitleOr404(db, id, w, r)
	if title == nil {
		return
	}
	if err := db.Delete(&title).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getTitleOr404(db *gorm.DB, titleID int, w http.ResponseWriter, r *http.Request) *model.Title {
	title := model.Title{}
	if err := db.First(&title, model.Title{Model: gorm.Model{ID: uint(titleID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &title
}
