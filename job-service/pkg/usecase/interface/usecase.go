package interfaces

import "Auth/pkg/utils/models"

type JobUseCase interface {
	PostJob(job models.JobOpening, employerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerID int32) ([]models.AllJob, error)
}
