package interfaces

import (
	"Auth/pkg/domain"
	"Auth/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
}

type EmployerUseCase interface {
	EmployerSignUp(employerDetails models.EmployerSignUp) (*domain.TokenEmployer, error)
	EmployerLogin(employerDetails models.EmployerLogin) (*domain.TokenEmployer, error)
}
