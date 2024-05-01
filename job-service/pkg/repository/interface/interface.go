package interfaces

import (
	"Auth/pkg/utils/models"
)

type JobRepository interface {
	PostJob(jobDetails models.JobOpening, employerID int32) (models.JobOpeningResponse, error)
}
