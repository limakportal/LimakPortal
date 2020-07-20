package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllOrganizations(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	organizations := []model.Organization{}
	db.Find(&organizations)
	respondJSON(w, http.StatusOK, organizations)
}

func CreateOrganization(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	organization := model.Organization{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&organization); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&organization).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, organization)
}

func GetOrganization(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["OrganizationID"])
	if err != nil {
		return
	}
	organization := getOrganizationOr404(db, id, w, r)
	if organization == nil {
		return
	}
	respondJSON(w, http.StatusOK, organization)
}

func UpdateOrganization(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["OrganizationID"])
	if err != nil {
		return
	}
	organization := getOrganizationOr404(db, id, w, r)
	if organization == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&organization); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&organization).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, organization)
}

func DeleteOrganization(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["OrganizationID"])
	if err != nil {
		return
	}
	organization := getOrganizationOr404(db, id, w, r)
	if organization == nil {
		return
	}
	if err := db.Delete(&organization).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getOrganizationOr404(db *gorm.DB, organizationID int, w http.ResponseWriter, r *http.Request) *model.Organization {
	organization := model.Organization{}
	if err := db.First(&organization, model.Organization{Model: gorm.Model{ID: uint(organizationID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &organization
}
