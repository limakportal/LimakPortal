package model

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model
	Name                 string    `json:"Name"`
	Surname              string    `json:"Surname"`
	IdentityID           string    `json:"IdentityID"`
	NationalityID        int       `json:"nationality_id"`
	Address              string    `json:"Address"`
	Telephone            string    `json:"Telephone"`
	BirthDate            time.Time `json:"BirthDate"`
	State                int       `json:"State"`
	GenderID             int       `json:"gender_id"`
	MaritalStatusID      int       `json:"maritalstatus_id"`
	RegisteredProvinceID int       `json:"registered_province_id"`
	PlaceOfRegistryID    int       `json:"place_of_registry_id"`
	IdentitySerialNumber string    `json:"IdentitySerialNumber"`
	IdentityVolumeNo     string    `json:"IdentityVolumeNo"`
	MothersName          string    `json:"MothersName"`
	FathersName          string    `json:"FathersName"`
	BloodType            string    `json:"BloodType"`
	Email                string    `json:"Email"`
	Picture              []byte    `json:"Picture"`
	PersonelInformation  []PersonelInformation `gorm:"ForeignKey:PersonID" json:"PersonelInformation"`
}

type Right struct {
	gorm.Model
	PersonID      int       `json:"person_id"`
	StartDate     time.Time `json:"StartDate"`
	EndDate       time.Time `json:"EndDate"`
	DateOfReturn  time.Time `json:"DateOfReturn"`
	Address       string    `json:"Address"`
	Telephone     string    `json:"Telephone"`
	Approver1     int       `json:"Approver1"`
	Approver2     int       `json:"Approver2"`
	RightTypeID   int       `json:"righttype_id"`
	RightStatusID int       `json:"rightstatus_id"`
	RightNumber   int       `json:"RightNumber"`
	DenyExplanation   string     `json:"DenyExplanation"`
	Person  []Person `gorm:"ForeignKey:PersonID" json:"Person"`
	RightStatus []RightStatus  `gorm:"ForeignKey:RightStatusID" json:"RightStatus"`
	RightType   []RightType  `gorm:"ForeignKey:RightTypeID" json:"RightType"`
}

type RightHistory struct {
	gorm.Model
	RightID       int       `json:"right_id"`
	ChangedBy     int       `json:"ChangedBy"`
	ChangedDate   time.Time `json:"ChangedDate"`
	RightStatusID int       `json:"rightstatus_id"`
}

type RightStatus struct {
	gorm.Model
	Name     string `json:"Name"`
	StatusID int    `json:"status_id"`
}

type RightType struct {
	gorm.Model
	Name     string `json:"Name"`
	StatusID int    `json:"status_id"`
}

type PersonHistory struct {
	gorm.Model
	PersonID        int       `json:"person_id"`
	StaffID         int       `json:"staff_id"`
	EntryDate       time.Time `json:"EntryDate"`
	TerminationDate time.Time `json:"TerminationDate"`
}

type PersonelInformation struct {
	PersonID       int       `gorm:"primary_key;auto_increment:false" json:"person_id"`
	RegisterNo     string    `json:"RegisterNo"`
	SGKRegisterNo  string    `json:"SGKRegisterNo"`
	SGKEnterDate   time.Time `json:"SGKEnterDate"`
	LimakEnterDate time.Time `json:"LimakEnterDate"`
}

type User struct {
	PersonID   int    `gorm:"primary_key;auto_increment:false" json:"person_id"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	IsApproved int    `json:"IsApproved"`
	IsActive   int    `json:"IsActive"`
	IsLocked   int    `json:"IsLocked"`
}

type District struct {
	gorm.Model
	Name     string `json:"Name"`
	CityID   int    `json:"city_id"`
	StatusID int    `json:"status_id"`
}

type Gender struct {
	gorm.Model
	Name string `json:"Name"`
}

type MaritalStatus struct {
	gorm.Model
	Name string `json:"Name"`
}

type Nationality struct {
	gorm.Model
	Name string `json:"Name"`
}

type City struct {
	gorm.Model
	Name     string `json:"Name"`
	StatusID int    `json:"status_id"`
}

type Staff struct {
	gorm.Model
	TitleID        int `json:"title_id"`
	OrganizationID int `json:"organization_id"`
	StatusID       int `json:"status_id"`
	PersonID       int `json:"person_id"`
}

type StaffHistory struct {
	gorm.Model
	StaffID        int `json:"staff_id"`
	TitleID        int `json:"title_id"`
	OrganizationID int `json:"organization_id"`
	PersonID       int `json:"person_id"`
}

type Title struct {
	gorm.Model
	Name     string `json:"Name"`
	StatusID int    `json:"status_id"`
}

type Status struct {
	gorm.Model
	Name string `json:"Name"`
}

type OrganizationType struct {
	gorm.Model
	Name     string `json:"Name"`
	StatusID int    `json:"status_id"`
}

type Organization struct {
	gorm.Model
	Name                string `json:"Name"`
	StatusID            int    `json:"status_id"`
	OrganizationTypeID  int    `json:"organization_type_id"`
	UpperOrganizationID int    `json:"upper_organization_id"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Person{}, &Right{}, &RightHistory{}, &RightStatus{}, &RightType{}, &PersonHistory{}, &PersonelInformation{}, &User{}, &District{}, &Gender{}, &MaritalStatus{}, &Nationality{}, &City{}, &Staff{}, &Title{}, &Status{}, &StaffHistory{}, &Organization{}, &OrganizationType{})

	db.Model(&Person{}).AddForeignKey("nationality_id", "nationalities(id)", "RESTRICT", "RESTRICT")
	db.Model(&Person{}).AddForeignKey("gender_id", "genders(id)", "RESTRICT", "RESTRICT")
	db.Model(&Person{}).AddForeignKey("marital_status_id", "marital_statuses(id)", "RESTRICT", "RESTRICT")
	db.Model(&Person{}).AddForeignKey("place_of_registry_id", "cities(id)", "RESTRICT", "RESTRICT")
	db.Model(&Person{}).AddForeignKey("registered_province_id", "cities(id)", "RESTRICT", "RESTRICT")

	db.Model(&Right{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")
	db.Model(&Right{}).AddForeignKey("right_type_id", "right_types(id)", "RESTRICT", "RESTRICT")
	db.Model(&Right{}).AddForeignKey("right_status_id", "right_statuses(id)", "RESTRICT", "RESTRICT")

	db.Model(&RightHistory{}).AddForeignKey("right_id", "rights(id)", "RESTRICT", "RESTRICT")
	db.Model(&RightHistory{}).AddForeignKey("right_status_id", "right_statuses(id)", "RESTRICT", "RESTRICT")

	db.Model(&RightStatus{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")

	db.Model(&RightType{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")

	db.Model(&PersonHistory{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")
	db.Model(&PersonHistory{}).AddForeignKey("staff_id", "staffs(id)", "RESTRICT", "RESTRICT")

	// db.Model(&PersonelInformation{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")

	//db.Model(&User{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")

	db.Model(&District{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")
	db.Model(&District{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")

	db.Model(&City{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")

	db.Model(&Staff{}).AddForeignKey("title_id", "titles(id)", "RESTRICT", "RESTRICT")
	db.Model(&Staff{}).AddForeignKey("organization_id", "organizations(id)", "RESTRICT", "RESTRICT")
	db.Model(&Staff{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")
	db.Model(&Staff{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")

	db.Model(&StaffHistory{}).AddForeignKey("title_id", "titles(id)", "RESTRICT", "RESTRICT")
	db.Model(&StaffHistory{}).AddForeignKey("organization_id", "organizations(id)", "RESTRICT", "RESTRICT") // İlk organizasyon mauel eklenecek
	db.Model(&StaffHistory{}).AddForeignKey("staff_id", "staffs(id)", "RESTRICT", "RESTRICT")
	db.Model(&StaffHistory{}).AddForeignKey("person_id", "people(id)", "RESTRICT", "RESTRICT")

	db.Model(&Title{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")

	db.Model(&OrganizationType{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")

	db.Model(&Organization{}).AddForeignKey("upper_organization_id", "organizations(id)", "RESTRICT", "RESTRICT")
	db.Model(&Organization{}).AddForeignKey("organization_type_id", "organization_types(id)", "RESTRICT", "RESTRICT")
	db.Model(&Organization{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")

	return db
}