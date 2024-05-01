package repository

import (
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"
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

	// Return the created job
	return models.JobOpeningResponse{
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
	}, nil
}
