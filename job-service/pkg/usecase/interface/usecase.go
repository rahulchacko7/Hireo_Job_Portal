package interfaces

import "Auth/pkg/utils/models"

type JobUseCase interface {
	PostJob(job models.JobOpening, employerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerID int32) ([]models.AllJob, error)
	GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error)
	DeleteAJob(employerIDInt, jobID int32) error
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error)
}
