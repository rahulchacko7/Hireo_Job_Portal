package domain

import "time"

type JobOpening struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title" gorm:"type:varchar(255);not null"`
	Description         string    `json:"description" gorm:"type:text;not null"`
	Requirements        string    `json:"requirements" gorm:"type:text;not null"`
	PostedOn            time.Time `json:"posted_on" gorm:"not null"`
	TotalApplications   int       `json:"total_applications" gorm:"not null;default:0"`
	CompanyName         string    `json:"company_name" gorm:"type:varchar(255);not null"`
	PostedByID          uint      `json:"posted_by_id" gorm:"not null"`
	PostedBy            Employer  `json:"posted_by" references employers(id)`
	Location            string    `json:"location" gorm:"type:varchar(255);not null"`
	EmploymentType      string    `json:"employment_type" gorm:"type:varchar(100);not null"`
	SalaryRange         string    `json:"salary_range" gorm:"type:varchar(100);not null"`
	SkillsRequired      []string  `json:"skills_required" gorm:"-"`
	ExperienceLevel     string    `json:"experience_level" gorm:"type:varchar(100);not null"`
	EducationLevel      string    `json:"education_level" gorm:"type:varchar(100);not null"`
	ApplicationDeadline time.Time `json:"application_deadline" gorm:"not null"`
}

type Employer struct {
	ID                   uint   `json:"id" gorm:"uniquekey; not null"`
	Company_name         string `json:"company_name" gorm:"validate:required"`
	Industry             string `json:"industry" gorm:"validate:required"`
	Company_size         int    `json:"company_size" gorm:"validate:required"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company" gorm:"type:text"`
	Contact_email        string `json:"contact_email" gorm:"validate:required"`
	Contact_phone_number uint   `json:"contact_phone_number" gorm:"type:numeric"`
	Password             string `json:"password" gorm:"validate:required"`
}
