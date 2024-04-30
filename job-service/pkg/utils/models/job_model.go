package models

import (
	"time"
)

type JobOpening struct {
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	SalaryRange         string    `json:"salary_range"`
	SkillsRequired      []string  `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}

type JobOpeningResponse struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	EmployerId          string    `json:"employer_id"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	SalaryRange         string    `json:"salary_range"`
	SkillsRequired      []string  `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}
