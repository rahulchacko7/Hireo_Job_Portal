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

	if err := jr.DB.Create(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) GetAllJobs(employerID int32) ([]models.AllJob, error) {
	var jobs []models.AllJob

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

	if err := jr.DB.Delete(&models.JobOpeningResponse{}, jobID).Error; err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}

func (jr *jobRepository) UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error) {

	postedOn := time.Now()

	updatedJob := models.JobOpeningResponse{
		ID:                  uint(jobID),
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

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).Updates(updatedJob).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return updatedJob, nil
}

func (jr *jobRepository) JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningResponse, error) {
	var jobSeekerJobs []models.JobOpeningResponse

	if err := jr.DB.Where("title ILIKE ?", "%"+keyword+"%").Find(&jobSeekerJobs).Error; err != nil {
		return nil, fmt.Errorf("failed to query jobs: %v", err)
	}

	fmt.Println(jobSeekerJobs)

	return jobSeekerJobs, nil

}

func (jr *jobRepository) GetJobDetails(jobID int32) (models.JobOpeningResponse, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).First(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (aj *jobRepository) ApplyJob(application models.ApplyJob, resumeURL string) (models.ApplyJobResponse, error) {
	var JobResponse models.ApplyJobResponse

	result := aj.DB.Exec("INSERT INTO apply_jobs (jobseeker_id, job_id, resume, resume_url, cover_letter) VALUES (?, ?, ?, ?, ?)",
		application.JobseekerID,
		application.JobID,
		nil,
		resumeURL,
		application.CoverLetter)

	if result.Error != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("error on inserting into database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return models.ApplyJobResponse{}, errors.New("no rows were affected during insert")
	}

	err := aj.DB.Raw("select * from apply_jobs where jobseeker_id = ? and job_id = ? ", application.JobseekerID, application.JobID).Scan(&JobResponse).Error
	if err != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	fmt.Println("job", JobResponse.ID)

	return JobResponse, nil
}
