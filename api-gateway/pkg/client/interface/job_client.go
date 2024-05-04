package interfaces

import "HireoGateWay/pkg/utils/models"

type JobClient interface {
	PostJobOpening(jobDetails models.JobOpening, EmployerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerIDInt int32) ([]models.AllJob, error)
	GetAJob(employerIDInt, jobId int32) (models.JobOpeningResponse, error)

	// // UpdateJobOpening updates an existing job opening.
	// UpdateJobOpening(jobDetails models.JobOpening) (models.JobToken, error)

	// // DeleteJobOpening deletes a job opening.
	// DeleteJobOpening(jobID uint64) error

	// // ListJobOpenings lists all available job openings.
	// ListJobOpenings() ([]models.JobDetails, error)
}
