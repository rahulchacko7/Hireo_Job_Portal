package repository

import (
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"

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

func (jr *jobRepository) PostJob(jobDetails models.JobOpening) (models.JobOpeningResponse, error) {
	// Create a new domain.JobOpeningResponse object to hold the data
	job := models.JobOpeningResponse{
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		PostedOn:            jobDetails.PostedOn,
		CompanyName:         jobDetails.CompanyName,
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		SalaryRange:         jobDetails.SalaryRange,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: jobDetails.ApplicationDeadline,
	}

	// Insert the job into the database
	if err := jr.DB.Create(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	// Return the created job
	return job, nil
}
