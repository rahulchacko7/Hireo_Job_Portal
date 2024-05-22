package usecase

import (
	"Auth/pkg/config"
	"Auth/pkg/helper"
	interfaces "Auth/pkg/repository/interface"
	services "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

	isJobExist, err := ju.jobRepository.IsJobExist(jobId)
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningResponse{}, fmt.Errorf("job with ID %d does not exist", jobId)
	}

	jobData, err := ju.jobRepository.GetAJob(employerID, jobId)
	if err != nil {
		return models.JobOpeningResponse{}, err
	}
	return jobData, nil
}

func (ju *jobUseCase) DeleteAJob(employerIDInt, jobID int32) error {

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
func (ju *jobUseCase) UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error) {

	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningResponse{}, fmt.Errorf("job with ID %d does not exist", jobID)
	}

	updatedJob, err := ju.jobRepository.UpdateAJob(employerID, jobID, jobDetails)
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to update job: %v", err)
	}

	return updatedJob, nil
}

func (ju *jobUseCase) JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error) {

	jobs, err := ju.jobRepository.JobSeekerGetAllJobs(keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %v", err)
	}

	var jobSeekerJobs []models.JobSeekerGetAllJobs
	for _, job := range jobs {

		jobSeekerJob := models.JobSeekerGetAllJobs{
			ID:    job.ID,
			Title: job.Title,
		}
		jobSeekerJobs = append(jobSeekerJobs, jobSeekerJob)
	}

	return jobSeekerJobs, nil
}

func (ju *jobUseCase) GetJobDetails(jobID int32) (models.JobOpeningResponse, error) {

	isJobExist, err := ju.jobRepository.IsJobExist(jobID)
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to check if job exists: %v", err)
	}

	if !isJobExist {
		return models.JobOpeningResponse{}, fmt.Errorf("job with ID %d does not exist", jobID)
	}

	jobData, err := ju.jobRepository.GetJobDetails(jobID)
	if err != nil {
		return models.JobOpeningResponse{}, err
	}

	return jobData, nil
}
func (ju *jobUseCase) ApplyJob(jobApplication models.ApplyJob, resumeData []byte) (models.ApplyJobResponse, error) {

	if jobApplication.JobID == 0 || jobApplication.JobseekerID == 0 || jobApplication.CoverLetter == "" {
		return models.ApplyJobResponse{}, errors.New("invalid input data")
	}

	fileUID := uuid.New()
	fileName := fileUID.String()
	h := helper.NewHelper(config.Config{})

	url, err := h.AddImageToAwsS3([]byte(jobApplication.ResumeURL), fileName)
	if err != nil {
		return models.ApplyJobResponse{}, err
	}

	fmt.Println("url", url)

	Data, err := ju.jobRepository.ApplyJob(jobApplication, url)
	if err != nil {
		return models.ApplyJobResponse{}, err
	}

	return Data, nil
}
