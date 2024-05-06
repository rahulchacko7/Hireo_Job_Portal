// employer_repository.go
package repository

import (
	"Auth/pkg/domain"
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type employerRepository struct {
	DB *gorm.DB
}

func NewEmployerRepository(DB *gorm.DB) interfaces.EmployerRepository {
	return &employerRepository{
		DB: DB,
	}
}

func (er *employerRepository) EmployerSignUp(employerDetails models.EmployerSignUp) (models.EmployerDetailsResponse, error) {
	var model models.EmployerDetailsResponse

	fmt.Println("email", model.ContactEmail)

	fmt.Println("models", model)
	if err := er.DB.Raw("INSERT INTO employers (company_name, industry, company_size, website, headquarters_address, about_company, contact_email, contact_phone_number, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, company_name, industry, company_size, website, headquarters_address, about_company, contact_email, contact_phone_number", employerDetails.CompanyName, employerDetails.Industry, employerDetails.CompanySize, employerDetails.Website, employerDetails.HeadquartersAddress, employerDetails.AboutCompany, employerDetails.ContactEmail, employerDetails.ContactPhoneNumber, employerDetails.Password).Scan(&model).Error; err != nil {
		return models.EmployerDetailsResponse{}, err
	}
	fmt.Println("inside", model.ContactEmail)
	return model, nil
}

func (er *employerRepository) CheckEmployerExistsByEmail(email string) (*domain.Employer, error) {
	var employer domain.Employer
	res := er.DB.Where(&domain.Employer{Contact_email: email}).First(&employer)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Employer{}, res.Error
	}
	return &employer, nil
}

func (er *employerRepository) FindEmployerByEmail(employer models.EmployerLogin) (models.EmployerSignUp, error) {
	var user models.EmployerSignUp
	err := er.DB.Raw("SELECT * FROM employers WHERE Contact_email=? ", employer.Email).Scan(&user).Error
	if err != nil {
		return models.EmployerSignUp{}, errors.New("error checking user details")
	}
	return user, nil
}

func (er *employerRepository) GetCompanyDetails(employerID int32) (models.EmployerDetailsResponse, error) {
	var user models.EmployerDetailsResponse
	err := er.DB.Raw("SELECT * FROM employers WHERE id=? ", employerID).Scan(&user).Error
	if err != nil {
		return models.EmployerDetailsResponse{}, errors.New("error checking user details")
	}
	return user, nil
}

func (er *employerRepository) UpdateCompany(ctx context.Context, employerIDInt int32, employerDetails models.EmployerDetails) (models.EmployerDetailsResponse, error) {
	// Prepare the SQL query to update the company details
	query := `
		UPDATE employers
		SET company_name = ?, industry = ?, company_size = ?, website = ?, headquarters_address = ?, about_company = ?, contact_email = ?, contact_phone_number = ?
		WHERE id = ?
		RETURNING id, company_name, industry, company_size, website, headquarters_address, about_company, contact_email, contact_phone_number
	`

	// Execute the SQL query
	var updatedEmployerDetails models.EmployerDetailsResponse
	result := er.DB.Raw(query,
		employerDetails.CompanyName,
		employerDetails.Industry,
		employerDetails.CompanySize,
		employerDetails.Website,
		employerDetails.HeadquartersAddress,
		employerDetails.AboutCompany,
		employerDetails.ContactEmail,
		employerDetails.ContactPhoneNumber,
		employerIDInt,
	).Scan(
		&updatedEmployerDetails.ID,
		&updatedEmployerDetails.CompanyName,
		&updatedEmployerDetails.Industry,
		&updatedEmployerDetails.CompanySize,
		&updatedEmployerDetails.Website,
		&updatedEmployerDetails.HeadquartersAddress,
		&updatedEmployerDetails.AboutCompany,
		&updatedEmployerDetails.ContactEmail,
		&updatedEmployerDetails.ContactPhoneNumber,
	)
	if result.Error != nil {
		return models.EmployerDetailsResponse{}, errors.Wrap(result.Error, "failed to update company details")
	}

	return updatedEmployerDetails, nil
}
