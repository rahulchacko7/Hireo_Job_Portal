// employer_repository.go
package repository

import (
	"Auth/pkg/domain"
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"
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
