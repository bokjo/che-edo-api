package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/bokjo/test_edo/handlers"
	"gitlab.com/bokjo/test_edo/model"
)

// API struct
type API struct {
	Router *mux.Router
	Db     *sql.DB
}

//Init - inilstialize the API
func (api *API) Init(username, password, dbname, host string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", username, password, dbname, host)

	var err error

	api.Db, err = model.DBConnect(connectionString)

	if err != nil {
		log.Fatal(err)
	}

	api.Router = mux.NewRouter()
	api.initRoutes()

}

// Run the API
func (api *API) Run(port string) {
	log.Fatal(http.ListenAndServe(":"+port, api.Router))
}

// TODO: move to separate endpoints file
// initRoutes - initiates the api routes
func (api *API) initRoutes() {

	defaultHandler := handlers.DefaultHandler{}

	api.Router.HandleFunc("/", defaultHandler.GetDefault).Methods(http.MethodGet)

	versionService := model.VersionService{}
	versionHandler := handlers.VersionHandler{VersionService: versionService}

	api.Router.HandleFunc("/version", versionHandler.GetVersion).Methods(http.MethodGet)

	jobsService := model.JobsService{DB: api.Db}
	jobsHandler := handlers.JobsHandler{JobsService: jobsService}

	api.Router.HandleFunc("/jobs", jobsHandler.GetJobs).Methods(http.MethodGet)
	api.Router.HandleFunc("/jobs", jobsHandler.CreateJob).Methods(http.MethodPost)

}
