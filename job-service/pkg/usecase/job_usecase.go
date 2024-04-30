package usecase

import (
	interfaces "Auth/pkg/repository/interface"
	services "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
)

type jobUseCase struct {
	jobRepository interfaces.JobRepository
}

func NewJobUseCase(repository interfaces.JobRepository) services.JobUseCase {
	return &jobUseCase{
		jobRepository: repository,
	}
}

func (ju *jobUseCase) PostJob(job models.JobOpening) (models.JobOpeningResponse, error) {
	jobData, err := ju.jobRepository.PostJob(job)
	if err != nil {
		return models.JobOpeningResponse{}, err
	}

	return jobData, nil
}
