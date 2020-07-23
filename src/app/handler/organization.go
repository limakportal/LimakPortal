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

func getMainOrganizationOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request) *model.Organization {
	organization := model.Organization{}
	if err := db.First(&organization, model.Organization{UpperOrganizationID: 0}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &organization
}

type OrganizationTree struct {
	Id                  int                 `json:"Id"`
	Name                string              `json:"Name"`
	UpperOrganizationID int                 `json:"upper_organization_id"`
	Organization        []*OrganizationTree `json:"Organization,omitempty"`
	Expanded            bool                `json:"expanded,omitempty"`
	ClassName           string              `json:"className"`
	Label               string              `json:"label,omitempty"`
	Type                string              `json:"type"`
}

func GetOrganizationTree(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var mainorganization = getMainOrganizationOr404(db, w, r)
	var root *OrganizationTree = &OrganizationTree{int(mainorganization.ID), mainorganization.Name, 0, nil, true, "", mainorganization.Name, ""}
	organizations := []model.Organization{}
	db.Find(&organizations)
	for i, _ := range organizations {
		data := []*OrganizationTree{
			&OrganizationTree{int(organizations[i].ID), organizations[i].Name, organizations[i].UpperOrganizationID, nil, true, "", organizations[i].Name, ""},
		}
		root.Add(data...)
	}

	respondJSON(w, http.StatusOK, root)
}
func (this *OrganizationTree) SliceAyarla() int {
	var i_slice int = len(this.Organization)
	for _, c := range this.Organization {
		i_slice += c.SliceAyarla()
	}
	return i_slice
}
func (this *OrganizationTree) Add(nodes ...*OrganizationTree) bool {
	var size = this.SliceAyarla()
	for _, n := range nodes {
		if n.UpperOrganizationID == this.Id {
			this.Organization = append(this.Organization, n)
		} else {
			for _, c := range this.Organization {
				if c.Add(n) {
					break
				}
			}
		}
	}
	return this.SliceAyarla() == size+len(nodes)
}
