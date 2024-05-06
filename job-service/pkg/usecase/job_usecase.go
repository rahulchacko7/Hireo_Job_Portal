package usecase

import (
	interfaces "Auth/pkg/repository/interface"
	services "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"fmt"
)

type jobUseCase struct {
	jobRepository interfaces.JobRepository
}

func NewJobUseCase(repository interfaces.JobRepository) services.JobUseCase {
	return &jobUseCase{
		jobRepository: repository,
	}
}

func (ju *jobUseCase) PostJob(job models.JobOpening, employerID int32) (models.JobOpeningResponse, error) {
	jobData, err := ju.jobRepository.PostJob(job, int32(employerID))
	if err != nil {
		return models.JobOpeningResponse{}, err
	}
	return jobData, nil
}

func (ju *jobUseCase) GetAllJobs(employerID int32) ([]models.AllJob, error) {
	jobData, err := ju.jobRepository.GetAllJobs(employerID)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}

func (ju *jobUseCase) GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error) {
	jobData, err := ju.jobRepository.GetAJob(employerID, jobId)
	if err != nil {
		return models.JobOpeningResponse{}, err
	}
	return jobData, nil
}

func (ju *jobUseCase) DeleteAJob(employerIDInt, jobID int32) error {
	// Check if the job exists
	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return fmt.Errorf("job with ID %d does not exist", jobID)
	}

	// If the job exists, proceed with deletion
	err = ju.jobRepository.DeleteAJob(employerIDInt, jobID)
	if err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}
