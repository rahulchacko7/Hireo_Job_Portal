package interfaces

import "HireoGateWay/pkg/utils/models"

type EmployerClient interface {
	EmployerSignUp(employerDetails models.EmployerSignUp) (models.TokenEmployer, error)
	EmployerLogin(employerDetails models.EmployerLogin) (models.TokenEmployer, error)
}
