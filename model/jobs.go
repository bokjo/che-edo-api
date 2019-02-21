package model

import (
	"database/sql"
	"fmt"
	"time"
)

//Job model
type Job struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Priority       int    `json:"prioriry"`
	State          string `json:"state"`
	StartDate      string `json:"start_date"`
	CompletionDate string `json:"completion_date"`
}

//JobsService for handling Jobs
type JobsService struct {
	DB  *sql.DB
	Job *Job
}

// GetJob - retrieve single job
func (js *JobsService) GetJob(id int) (*Job, error) {
	job := Job{}
	const GetSingleJobQuery = `SELECT id, name, priority, state, start_date, completion_date FROM jobs WHERE id=$1`
	err := js.DB.QueryRow(GetSingleJobQuery, id).Scan(&job.ID, &job.Name, &job.Priority, &job.State, &job.StartDate, &job.CompletionDate)

	return &job, err
}

// CreateJob - create single job
func (js *JobsService) CreateJob(priority int) (string, error) {
	//job := js.Job
	currentTime := time.Now().Local()
	completedTime := time.Now().Local().Add(time.Second * 5)
	const CreateSingleJobQuery = `INSERT INTO jobs(name, priority, state, start_date, completion_date) VALUES($1, $2, $3, $4, $5) RETURNING id`
	// err := js.DB.QueryRow(CreateSingleJobQuery, job.Name, job.Priority, job.State, job.StartDate, job.CompletionDate).Scan(&job.ID)
	var jobID string
	err := js.DB.QueryRow(CreateSingleJobQuery, "JOB: test", priority, "COMPLETED", currentTime, completedTime).Scan(&jobID)

	return jobID, err
}

// UpdateJob - updates single job
func (js *JobsService) UpdateJob(id int) error {

	job := js.Job
	const UpdateJobQuery = "UPDATE jobs SET name=$1, priority=$2, state=$3, start_date=$4, completed_date=$5"
	_, err := js.DB.Exec(UpdateJobQuery, job.Name, job.Priority, job.State, job.StartDate, job.CompletionDate)

	return err
}

// DeleteJob - updates single job
func (js *JobsService) DeleteJob(id int) (sql.Result, error) {

	const DeleteJobQuery = "DELETE FROM jobs WHERE id=$1"
	deleted, err := js.DB.Exec(DeleteJobQuery, id)

	return deleted, err
}

// GetJobs - return all jobs
func (js *JobsService) GetJobs(sortBy string) ([]Job, error) {

	SelectAllJobsQuery := "SELECT id, name, priority, state, start_date, completion_date FROM jobs "

	if sortBy != "" {
		SelectAllJobsQuery = SelectAllJobsQuery + fmt.Sprintf("ORDER BY %s DESC", sortBy)
	}

	rows, err := js.DB.Query(SelectAllJobsQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobs := []Job{}

	for rows.Next() {
		var job Job

		if err := rows.Scan(&job.ID, &job.Name, &job.Priority, &job.State, &job.StartDate, &job.CompletionDate); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)

	}

	return jobs, nil

}
