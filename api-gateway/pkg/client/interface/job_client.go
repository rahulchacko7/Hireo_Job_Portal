package interfaces

import "HireoGateWay/pkg/utils/models"

type JobClient interface {
	PostJobOpening(jobDetails models.JobOpening, EmployerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerIDInt int32) ([]models.AllJob, error)
	GetAJob(employerIDInt, jobId int32) (models.JobOpeningResponse, error)
	DeleteAJob(employerIDInt, jobID int32) error
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	GetJobDetails(jobID int32)
	UpdateAJob(employerIDInt int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error)
}
