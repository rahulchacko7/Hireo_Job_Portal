package interfaces

import (
	"Auth/pkg/utils/models"
)

type JobUseCase interface {
	PostJob(job models.JobOpening) (models.JobOpeningResponse, error)
}
