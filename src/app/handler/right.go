package handler

import (
	"encoding/json"
	"limakcv/src/app/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllRights(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rights := []model.Right{}
	db.Find(&rights)
	respondJSON(w, http.StatusOK, rights)
}

func CreateRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	right := model.Right{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&right); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, right)
}

func GetRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightID"])
	if err != nil {
		return
	}
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}
	respondJSON(w, http.StatusOK, right)
}

func UpdateRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightID"])
	if err != nil {
		return
	}
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&right); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, right)
}

func DeleteRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightID"])
	if err != nil {
		return
	}
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}
	if err := db.Delete(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}


func GetPersonsAllRights(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	right := []model.Right{}
	if err := db.Model(&person).Related(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}	
	respondJSON(w, http.StatusOK, right)
}

func CreatePersonsRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	right :=   model.Right{PersonID: int(person.ID)}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&right); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, right)
}

func GetPersonsRights(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	personid, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, personid, w, r)
	if person == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}	
	respondJSON(w, http.StatusOK, right)
}

func UpdatePersonsRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	personid, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, personid, w, r)
	if person == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&right); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, right)
}

func DeletePersonsRights(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	personid, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, personid, w, r)
	if person == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}

	if err := db.Delete(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetPersonsRightDesc(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	personid, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, personid, w, r)
	if person == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}
	db.Model(right).Related(&right.Person)
	db.Model(right).Related(&right.RightStatus)
	db.Model(right).Related(&right.RightType)
	respondJSON(w, http.StatusOK, right)
}

func GetPersonsAllRightsDesc(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["PersonID"])
	if err != nil {
		return
	}
	person := getPersonOr404(db, id, w, r)
	if person == nil {
		return
	}

	right := []model.Right{}
	if err := db.Model(&person).Related(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	for i, _ := range right {
		db.Model(right[i]).Related(&right[i].Person)
		db.Model(right[i]).Related(&right[i].RightStatus)
		db.Model(right[i]).Related(&right[i].RightType)
	}
	respondJSON(w, http.StatusOK, right)
}


func getRightOr404(db *gorm.DB, rightID int, w http.ResponseWriter, r *http.Request) *model.Right {
	right := model.Right{}
	if err := db.First(&right, model.Right{Model: gorm.Model{ID: uint(rightID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &right
}
