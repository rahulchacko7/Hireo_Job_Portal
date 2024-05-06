package repository

import (
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type jobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(DB *gorm.DB) interfaces.JobRepository {
	return &jobRepository{
		DB: DB,
	}
}

func (jr *jobRepository) PostJob(jobDetails models.JobOpening, employerID int32) (models.JobOpeningResponse, error) {
	// Get the current time for posted on
	postedOn := time.Now()

	job := models.JobOpeningResponse{
		ID:                  jobDetails.ID,
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		PostedOn:            postedOn,
		EmployerID:          int(employerID),
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: jobDetails.ApplicationDeadline,
	}

	// Insert the job into the database
	if err := jr.DB.Create(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	// Return the created job with the generated ID
	return job, nil
}

func (jr *jobRepository) GetAllJobs(employerID int32) ([]models.AllJob, error) {
	var jobs []models.AllJob

	// Execute select query to retrieve all jobs
	if err := jr.DB.Model(&models.JobOpeningResponse{}).Select("id, title, application_deadline, employer_id").Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (jr *jobRepository) GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ? AND employer_id = ?", jobId, employerID).First(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) IsJobExist(jobID int32) (bool, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
func (jr *jobRepository) DeleteAJob(employerIDInt, jobID int32) error {
	// First, check if the job exists
	var job models.JobOpeningResponse
	if err := jr.DB.First(&job, jobID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("job with ID %d does not exist", jobID)
		}
		return err // Other errors
	}

	// Update the job's is_deleted status to true
	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).Update("is_deleted", true).Error; err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}
