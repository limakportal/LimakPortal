package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllStaffs(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	staffs := []model.Staff{}
	db.Find(&staffs)
	respondJSON(w, http.StatusOK, staffs)
}

func CreateStaff(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	staff := model.Staff{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&staff); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&staff).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, staff)
}

func GetStaff(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["StaffID"])
	if err != nil {
		return
	}
	staff := getStaffOr404(db, id, w, r)
	if staff == nil {
		return
	}
	respondJSON(w, http.StatusOK, staff)
}

func UpdateStaff(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["StaffID"])
	if err != nil {
		return
	}
	staff := getStaffOr404(db, id, w, r)
	if staff == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&staff); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&staff).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, staff)
}

func DeleteStaff(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["StaffID"])
	if err != nil {
		return
	}
	staff := getStaffOr404(db, id, w, r)
	if staff == nil {
		return
	}
	if err := db.Delete(&staff).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getStaffOr404(db *gorm.DB, staffID int, w http.ResponseWriter, r *http.Request) *model.Staff {
	staff := model.Staff{}
	if err := db.First(&staff, model.Staff{Model: gorm.Model{ID: uint(staffID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &staff
}
