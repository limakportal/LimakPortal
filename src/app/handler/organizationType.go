package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllOrganizationTypes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	organizationTypes := []model.OrganizationType{}
	db.Find(&organizationTypes)
	respondJSON(w, http.StatusOK, organizationTypes)
}

func CreateOrganizationType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	organizationType := model.OrganizationType{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&organizationType); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&organizationType).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, organizationType)
}

func GetOrganizationType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["OrganizationTypeID"])
	if err != nil {
		return
	}
	organizationType := getOrganizationTypeOr404(db, id, w, r)
	if organizationType == nil {
		return
	}
	respondJSON(w, http.StatusOK, organizationType)
}

func UpdateOrganizationType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["OrganizationTypeID"])
	if err != nil {
		return
	}
	organizationType := getOrganizationTypeOr404(db, id, w, r)
	if organizationType == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&organizationType); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&organizationType).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, organizationType)
}

func DeleteOrganizationType(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["OrganizationTypeID"])
	if err != nil {
		return
	}
	organizationType := getOrganizationTypeOr404(db, id, w, r)
	if organizationType == nil {
		return
	}
	if err := db.Delete(&organizationType).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getOrganizationTypeOr404(db *gorm.DB, organizationTypeID int, w http.ResponseWriter, r *http.Request) *model.OrganizationType {
	organizationType := model.OrganizationType{}
	if err := db.First(&organizationType, model.OrganizationType{Model: gorm.Model{ID: uint(organizationTypeID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &organizationType
}
