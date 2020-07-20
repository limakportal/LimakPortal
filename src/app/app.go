package app

import (
	"log"
	"net/http"

	"limakcv/src/app/handler"
	"limakcv/src/app/model"
	"limakcv/src/config"
	"limakcv/src/token"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
	// 	config.DB.Username,
	// 	config.DB.Password,
	// 	config.DB.Host,
	// 	config.DB.Port,
	// 	config.DB.Name,
	// 	config.DB.Charset,
	// )

	//db, err := gorm.Open(config.DB.Dialect, dbURI)
	db, err := gorm.Open("postgres", "host=kandula.db.elephantsql.com port=5432 user=plrvuppn password=DyhDQ6VlBGElGdX-qTJSjB5mR1fAvkrd dbname=plrvuppn")
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/persons", a.handleRequest(token.ValidateToken(handler.GetAllPersons)))
	a.Post("/persons", a.handleRequest(token.ValidateToken(handler.CreatePerson)))
	a.Get("/persons/{PersonID}", a.handleRequest(token.ValidateToken(handler.GetPerson)))
	a.Put("/persons/{PersonID}", a.handleRequest(token.ValidateToken(handler.UpdatePerson)))
	a.Delete("/persons/{PersonID}", a.handleRequest(token.ValidateToken(handler.DeletePerson)))

	a.Get("/genders", a.handleRequest(token.ValidateToken(handler.GetAllGenders)))
	a.Post("/genders", a.handleRequest(token.ValidateToken(handler.CreateGender)))
	a.Get("/genders/{GenderID}", a.handleRequest(token.ValidateToken(handler.GetGender)))
	a.Put("/genders/{GenderID}", a.handleRequest(token.ValidateToken(handler.UpdateGender)))
	a.Delete("/genders/{GenderID}", a.handleRequest(token.ValidateToken(handler.DeleteGender)))

	a.Get("/cities", a.handleRequest(token.ValidateToken(handler.GetAllCities)))
	a.Post("/cities", a.handleRequest(token.ValidateToken(handler.CreateCity)))
	a.Get("/cities/{CityID}", a.handleRequest(token.ValidateToken(handler.GetCity)))
	a.Put("/cities/{CityID}", a.handleRequest(token.ValidateToken(handler.UpdateCity)))
	a.Delete("/cities/{CityID}", a.handleRequest(token.ValidateToken(handler.DeleteCity)))

	a.Get("/statuses", a.handleRequest(token.ValidateToken(handler.GetAllStatues)))
	a.Post("/statuses", a.handleRequest(token.ValidateToken(handler.CreateStatus)))
	a.Get("/statuses/{StatusID}", a.handleRequest(token.ValidateToken(handler.GetStatus)))
	a.Put("/statuses/{StatusID}", a.handleRequest(token.ValidateToken(handler.UpdateStatus)))
	a.Delete("/statuses/{StatusID}", a.handleRequest(token.ValidateToken(handler.DeleteStatus)))

	a.Get("/maritalstatuses", a.handleRequest(token.ValidateToken(handler.GetAllMaritalStatus)))
	a.Post("/maritalstatuses", a.handleRequest(token.ValidateToken(handler.CreateMaritalStatus)))
	a.Get("/maritalstatuses/{MaritalStatusID}", a.handleRequest(token.ValidateToken(handler.GetMaritalStatus)))
	a.Put("/maritalstatuses/{MaritalStatusID}", a.handleRequest(token.ValidateToken(handler.UpdateMaritalStatus)))
	a.Delete("/maritalstatuses/{MaritalStatusID}", a.handleRequest(token.ValidateToken(handler.DeleteMaritalStatus)))

	a.Get("/nationalities", a.handleRequest(token.ValidateToken(handler.GetAllNationalities)))
	a.Post("/nationalities", a.handleRequest(token.ValidateToken(handler.CreateNationality)))
	a.Get("/nationalities/{NationalityID}", a.handleRequest(token.ValidateToken(handler.GetNationality)))
	a.Put("/nationalities/{NationalityID}", a.handleRequest(token.ValidateToken(handler.UpdateNationality)))
	a.Delete("/nationalities/{NationalityID}", a.handleRequest(token.ValidateToken(handler.DeleteNationality)))

	a.Get("/users", a.handleRequest(token.ValidateToken(handler.GetAllUsers)))
	a.Post("/users", a.handleRequest(token.ValidateToken(handler.CreateUser)))
	a.Get("/users/{UserID}", a.handleRequest(token.ValidateToken(handler.GetUser)))
	a.Put("/users/{UserID}", a.handleRequest(token.ValidateToken(handler.UpdateUser)))
	a.Delete("/users/{UserID}", a.handleRequest(token.ValidateToken(handler.DeleteUser)))

	a.Get("/staff", a.handleRequest(token.ValidateToken(handler.GetAllStaffs)))
	a.Post("/staff", a.handleRequest(token.ValidateToken(handler.CreateStaff)))
	a.Get("/staff/{StaffID}", a.handleRequest(token.ValidateToken(handler.GetStaff)))
	a.Put("/staff/{StaffID}", a.handleRequest(token.ValidateToken(handler.UpdateStaff)))
	a.Delete("/staff/{StaffID}", a.handleRequest(token.ValidateToken(handler.DeleteStaff)))

	a.Get("/district", a.handleRequest(token.ValidateToken(handler.GetAllDistricts)))
	a.Post("/district", a.handleRequest(token.ValidateToken(handler.CreateDistrict)))
	a.Get("/district/{DistrictID}", a.handleRequest(token.ValidateToken(handler.GetDistrict)))
	a.Put("/district/{DistrictID}", a.handleRequest(token.ValidateToken(handler.UpdateDistrict)))
	a.Delete("/district/{DistrictID}", a.handleRequest(token.ValidateToken(handler.DeleteDistrict)))

	a.Get("/personalinformations", a.handleRequest(token.ValidateToken(handler.GetAllPersonelInformations)))
	a.Post("/personalinformations", a.handleRequest(token.ValidateToken(handler.CreatePersonelInformation)))
	a.Get("/personalinformations/{PersonID}", a.handleRequest(token.ValidateToken(handler.GetPersonelInformation)))
	a.Put("/personalinformations/{PersonID}", a.handleRequest(token.ValidateToken(handler.UpdatePersonelInformation)))
	a.Delete("/personalinformations/{PersonID}", a.handleRequest(token.ValidateToken(handler.DeletePersonelInformation)))

	a.Get("/personhistories", a.handleRequest(token.ValidateToken(handler.GetAllPersonHistories)))
	a.Post("/personhistories", a.handleRequest(token.ValidateToken(handler.CreatePersonHistory)))
	a.Get("/personhistories/{PersonHistoryID}", a.handleRequest(token.ValidateToken(handler.GetPersonHistory)))
	a.Put("/personhistories/{PersonHistoryID}", a.handleRequest(token.ValidateToken(handler.UpdatePersonHistory)))
	a.Delete("/personhistories/{PersonHistoryID}", a.handleRequest(token.ValidateToken(handler.DeletePersonHistory)))

	a.Get("/titles", a.handleRequest(token.ValidateToken(handler.GetAllTitles)))

	a.Post("/titles", a.handleRequest(token.ValidateToken(handler.CreateTitle)))

	a.Get("/titles/{TitleID}", a.handleRequest(token.ValidateToken(handler.GetTitle)))

	a.Put("/titles/{TitleID}", a.handleRequest(token.ValidateToken(handler.UpdateTitle)))

	a.Delete("/titles/{TitleID}", a.handleRequest(token.ValidateToken(handler.DeleteTitle)))

	a.Get("/rights", a.handleRequest(token.ValidateToken(handler.GetAllRights)))

	a.Post("/rights", a.handleRequest(token.ValidateToken(handler.CreateRight)))

	a.Get("/rights/{RightID}", a.handleRequest(token.ValidateToken(handler.GetRight)))

	a.Put("/rights/{RightID}", a.handleRequest(token.ValidateToken(handler.UpdateRight)))

	a.Delete("/rights/{RightID}", a.handleRequest(token.ValidateToken(handler.DeleteRight)))

	a.Get("/rightHistoryes", a.handleRequest(token.ValidateToken(handler.GetAllRightHistoryes)))

	a.Post("/rightHistoryes", a.handleRequest(token.ValidateToken(handler.CreateRightHistory)))

	a.Get("/rightHistoryes/{RightHistoryID}", a.handleRequest(token.ValidateToken(handler.GetRightHistory)))

	a.Put("/rightHistoryes/{RightHistoryID}", a.handleRequest(token.ValidateToken(handler.UpdateRightHistory)))

	a.Delete("/rightHistoryes/{RightHistoryID}", a.handleRequest(token.ValidateToken(handler.DeleteRightHistory)))

	a.Get("/rightStatuses", a.handleRequest(token.ValidateToken(handler.GetAllRightStatuses)))

	a.Post("/rightStatuses", a.handleRequest(token.ValidateToken(handler.CreateRightStatus)))

	a.Get("/rightStatuses/{RightStatusID}", a.handleRequest(token.ValidateToken(handler.GetRightStatus)))

	a.Put("/rightStatuses/{RightStatusID}", a.handleRequest(token.ValidateToken(handler.UpdateRightStatus)))

	a.Delete("/rightStatuses/{RightStatusID}", a.handleRequest(token.ValidateToken(handler.DeleteRightStatus)))

	a.Get("/rightTypes", a.handleRequest(token.ValidateToken(handler.GetAllRightTypes)))

	a.Post("/rightTypes", a.handleRequest(token.ValidateToken(handler.CreateRightType)))

	a.Get("/rightTypes/{RightTypeID}", a.handleRequest(token.ValidateToken(handler.GetRightType)))

	a.Put("/rightTypes/{RightTypeID}", a.handleRequest(token.ValidateToken(handler.UpdateRightType)))

	a.Delete("/rightTypes/{RightTypeID}", a.handleRequest(token.ValidateToken(handler.DeleteRightType)))

	a.Get("/organizations", a.handleRequest(token.ValidateToken(handler.GetAllOrganizations)))

	a.Post("/organizations", a.handleRequest(token.ValidateToken(handler.CreateOrganization)))

	a.Get("/organizations/{OrganizationID}", a.handleRequest(token.ValidateToken(handler.GetOrganization)))

	a.Put("/organizations/{OrganizationID}", a.handleRequest(token.ValidateToken(handler.UpdateOrganization)))

	a.Delete("/organizations/{OrganizationID}", a.handleRequest(token.ValidateToken(handler.DeleteOrganization)))

	a.Get("/organizationTypes", a.handleRequest(token.ValidateToken(handler.GetAllOrganizationTypes)))

	a.Post("/organizationTypes", a.handleRequest(token.ValidateToken(handler.CreateOrganizationType)))

	a.Get("/organizationTypes/{OrganizationTypeID}", a.handleRequest(token.ValidateToken(handler.GetOrganizationType)))

	a.Put("/organizationTypes/{OrganizationTypeID}", a.handleRequest(token.ValidateToken(handler.UpdateOrganizationType)))

	a.Delete("/organizationTypes/{OrganizationTypeID}", a.handleRequest(token.ValidateToken(handler.DeleteOrganizationType)))

	a.Post("/login", a.handleRequest(handler.Login))

}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) handleRequest(handler token.RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}