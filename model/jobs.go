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
	Priority       int    `json:"priority"`
	State          string `json:"state"`
	StartDate      string `json:"started"`
	CompletionDate string `json:"completed"`
}

//JobsService for handling Jobs
type JobsService struct {
	DB  *sql.DB
	Job *Job
}

// GetJob - retrieve single job
func (js *JobsService) GetJob(id int) (*Job, error) {
	job := Job{}
	const GetSingleJobQuery = `SELECT id, name, priority, state, started, completed FROM jobs WHERE id=$1`
	err := js.DB.QueryRow(GetSingleJobQuery, id).Scan(&job.ID, &job.Name, &job.Priority, &job.State, &job.StartDate, &job.CompletionDate)

	return &job, err
}

// CreateJob - create single job
func (js *JobsService) CreateJob(priority int) (string, error) {

	var jobID string
	const CreateSingleJobQuery = `INSERT INTO jobs(name, priority, state, started, completed) VALUES($1, $2, $3, $4, $5) RETURNING id`

	jobName := "[Generic job name]"
	currentTime := time.Now().Local()
	completedTime := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)

	stage := assignJobStage()

	if stage == "COMPLETED" {
		completedTime = generateRandomWorkTime(currentTime)
	}

	err := js.DB.QueryRow(CreateSingleJobQuery, jobName, priority, stage, currentTime, completedTime).Scan(&jobID)

	return jobID, err
}

// UpdateJob - updates single job
func (js *JobsService) UpdateJob(id int) error {

	job := js.Job
	const UpdateJobQuery = "UPDATE jobs SET name=$1, priority=$2, state=$3, started=$4, completed=$5"
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

	SelectAllJobsQuery := "SELECT id, name, priority, state, started, completed FROM jobs "

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
