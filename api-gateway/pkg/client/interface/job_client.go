package interfaces

import "HireoGateWay/pkg/utils/models"

type JobClient interface {
	PostJobOpening(jobDetails models.JobOpening, EmployerID int) (models.JobOpeningResponse, error)

	// // UpdateJobOpening updates an existing job opening.
	// UpdateJobOpening(jobDetails models.JobOpening) (models.JobToken, error)

	// // GetJobOpening retrieves details of a specific job opening.
	// GetJobOpening(jobID uint64) (models.JobDetails, error)

	// // DeleteJobOpening deletes a job opening.
	// DeleteJobOpening(jobID uint64) error

	// // ListJobOpenings lists all available job openings.
	// ListJobOpenings() ([]models.JobDetails, error)
}
