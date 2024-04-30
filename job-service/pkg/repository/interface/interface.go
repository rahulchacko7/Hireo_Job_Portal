package interfaces

import (
	"Auth/pkg/utils/models"
)

type JobRepository interface {
	PostJob(jobDetails models.JobOpening) (models.JobOpeningResponse, error)
}
