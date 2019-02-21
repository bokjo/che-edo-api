package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"gitlab.com/bokjo/test_edo/model"
)

// JobsHandler handler struct
type JobsHandler struct {
	JobsService model.JobsService
}

// GetJob handles retrieving single job
func (jh *JobsHandler) GetJob(w http.ResponseWriter, r *http.Request) {

	// TODO: [DRY] - extract common functions to utils/helpers!!!
	// Parse request URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}

	Job, err := jh.JobsService.GetJob(id)

	if err != nil {

		// TODO: Move to utils/helpers!!!
		switch {
		case err == sql.ErrNoRows:
			errorRespond(w, http.StatusNotFound, fmt.Sprintf("Job with ID: %d not found", id))
		default:
			errorRespond(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	jsonRespond(w, http.StatusOK, Job)

}

// CreateJob handles single job creation
func (jh *JobsHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	// newJob := &jh.JobsService.Job
	// newJob.Name = "JOB: 2"
	// newJob.State = "COMPLETED"
	// newJob.StartDate = time.Now().Local().String()
	// newJob.CompletionDate = time.Now().Local().Add(time.Second * 5).String()
	// newJob.Priority = 0

	// test := jh.JobsService.Job

	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&newJob); err != nil {
	// 	errorRespond(w, http.StatusBadRequest, "Invalid request payload for creating new Job!")
	// 	return
	// }

	// defer r.Body.Close()

	//priority := r.FormValue(priority)

	//vars := mux.Vars(r)
	var priority int
	var err error
	priorityParam := r.FormValue("priority")

	if priorityParam == "" {
		priority = 0
	} else {
		priority, err = strconv.Atoi(priorityParam)

		if err != nil {
			errorRespond(w, http.StatusBadRequest, "Invalid priority value [Please provide integer value '< 0 lower priority, = 0 default, > 0 higher priority]")
			return
		}
	}

	//TEMP
	type ID struct {
		ID string `json:"id"`
	}

	var jobID ID

	if jobID.ID, err = jh.JobsService.CreateJob(priority); err != nil {
		log.Print(err.Error())
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	//TODO: put output in log file
	log.Printf("[JOB] CREATED: %s", jobID.ID)
	jsonRespond(w, http.StatusOK, jobID)

}

// UpdateJob handles updating single job
func (jh *JobsHandler) UpdateJob(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}

	updateJob := &jh.JobsService.Job

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(updateJob); err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid request payload for updating a Job!")
		return
	}

	defer r.Body.Close()

	if err := jh.JobsService.UpdateJob(id); err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	jsonRespond(w, http.StatusOK, updateJob)

}

// DeleteJob handles deleting single job
func (jh *JobsHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}

	ret, err := jh.JobsService.DeleteJob(id)
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	cnt, err := ret.RowsAffected()
	if cnt == 0 {
		errorRespond(w, http.StatusInternalServerError, "NOTHING TO DELETE: The Job with the specified ID does not exists!")
		return
	}

	jsonRespond(w, http.StatusOK, map[string]string{"Result": "SUCESS: Job successfully deleted!"})

}

// GetJobs handles retrieving all the jobs
func (jh *JobsHandler) GetJobs(w http.ResponseWriter, r *http.Request) {

	sortBy := r.FormValue("sortBy")

	Jobs, err := jh.JobsService.GetJobs(sortBy)

	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(Jobs) == 0 {
		jsonRespond(w, http.StatusOK, map[string]string{"[INFO]": "No jobs found... Add one with POST request on the same url ;)"})
	} else {
		jsonRespond(w, http.StatusOK, Jobs)
	}

}

// TODO: Move to utils/helpers !!!
// errorRespond - responds with custom JSON error message
func errorRespond(w http.ResponseWriter, statusCode int, message string) {
	jsonRespond(w, statusCode, map[string]string{"error": message})
}

// jsonRespond - returns valid JSON response
func jsonRespond(w http.ResponseWriter, statusCode int, payload interface{}) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
