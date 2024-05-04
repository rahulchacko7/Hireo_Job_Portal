package interfaces

import (
	"Auth/pkg/utils/models"
)

type JobRepository interface {
	PostJob(jobDetails models.JobOpening, employerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerID int32) ([]models.AllJob, error)
	GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error)
}
