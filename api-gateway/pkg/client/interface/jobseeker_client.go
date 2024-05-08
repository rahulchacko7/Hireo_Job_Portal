package interfaces

import "HireoGateWay/pkg/utils/models"

type JobSeekerClient interface {
	JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.TokenJobSeeker, error)
	JobSeekerLogin(jobSeekerDetails models.JobSeekerLogin) (models.TokenJobSeeker, error)
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
}
