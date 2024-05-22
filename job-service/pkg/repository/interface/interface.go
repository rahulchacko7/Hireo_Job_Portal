package interfaces

import (
	"Auth/pkg/utils/models"
)

type JobRepository interface {
	PostJob(jobDetails models.JobOpening, employerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerID int32) ([]models.AllJob, error)
	GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error)
	IsJobExist(jobID int32) (bool, error)
	DeleteAJob(employerIDInt, jobID int32) error
	JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningResponse, error)
	GetJobDetails(jobID int32) (models.JobOpeningResponse, error)
	ApplyJob(application models.ApplyJob, resumeURL string) (models.ApplyJobResponse, error)
	UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error)
}
