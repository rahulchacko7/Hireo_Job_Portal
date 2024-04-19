package models

type JobSeekerLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type JobSeekerSignUp struct {
	Email            string            `json:"email" binding:"required" validate:"required,email"`
	Password         string            `json:"password" binding:"required" validate:"min=6,max=20"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	PhoneNumber      string            `json:"phone_number"`
	Address          string            `json:"address"`
	DateOfBirth      string            `json:"date_of_birth"`
	Gender           string            `json:"gender"`
	Bio              string            `json:"bio"`
	SocialMedia      map[string]string `json:"social_media_links"`
	Skills           []string          `json:"skills"`
	Education        []Education       `json:"education_history"`
	WorkExperience   []Experience      `json:"work_experience"`
	LanguagesSpoken  []string          `json:"languages_spoken"`
	Interests        []string          `json:"interests"`
	Projects         []Project         `json:"projects"`
}

type Education struct {
	Degree         string `json:"degree"`
	University     string `json:"university"`
	GraduationYear string `json:"graduation_year"`
}

type Experience struct {
	Company     string `json:"company"`
	Position    string `json:"position"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type Project struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	TechnologiesUsed []string `json:"technologies_used"`
	StartDate        string   `json:"start_date"`
	EndDate          string   `json:"end_date"`
}

type JobSeekerSignUpResponse struct {
	ID              uint              `json:"id"`
	Email           string            `json:"email"`
	FirstName       string            `json:"first_name"`
	LastName        string            `json:"last_name"`
	PhoneNumber     string            `json:"phone_number"`
	Address         string            `json:"address"`
	DateOfBirth     string            `json:"date_of_birth"`
	Gender          string            `json:"gender"`
	Bio             string            `json:"bio"`
	SocialMedia     map[string]string `json:"social_media_links"`
	Skills          []string          `json:"skills"`
	Education       []Education       `json:"education_history"`
	WorkExperience  []Experience      `json:"work_experience"`
	LanguagesSpoken []string          `json:"languages_spoken"`
	Interests       []string          `json:"interests"`
	Projects        []Project         `json:"projects"`
}

type TokenJobSeeker struct {
	JobSeeker JobSeekerSignUpResponse
	Token     string
}
