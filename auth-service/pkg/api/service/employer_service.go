// employer_service.go
package service

import (
	pb "Auth/pkg/pb/auth"
	interfaces "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"context"
	"fmt"
)

type EmployerServer struct {
	employerUseCase interfaces.EmployerUseCase
	pb.UnimplementedEmployerServer
}

func NewEmployerServer(useCase interfaces.EmployerUseCase) pb.EmployerServer {
	return &EmployerServer{
		employerUseCase: useCase,
	}
}

func (es *EmployerServer) EmployerSignup(ctx context.Context, req *pb.EmployerSignupRequest) (*pb.EmployerSignupResponse, error) {
	employerSignup := models.EmployerSignUp{
		CompanyName:         req.CompanyName,
		Industry:            req.Industry,
		CompanySize:         int(req.CompanySize),
		Website:             req.Website,
		HeadquartersAddress: req.HeadquartersAddress,
		AboutCompany:        req.AboutCompany,
		ContactEmail:        req.ContactEmail,
		ContactPhoneNumber:  uint(req.ContactPhoneNumber),
		Password:            req.Password,
	}

	fmt.Println("service", employerSignup)

	res, err := es.employerUseCase.EmployerSignUp(employerSignup)
	if err != nil {
		return &pb.EmployerSignupResponse{}, err
	}

	employerDetails := &pb.EmployerDetails{
		Id:                  uint64(res.Employer.ID),
		CompanyName:         res.Employer.CompanyName,
		Industry:            res.Employer.Industry,
		CompanySize:         int32(res.Employer.CompanySize),
		Website:             res.Employer.Website,
		HeadquartersAddress: res.Employer.HeadquartersAddress,
		AboutCompany:        res.Employer.AboutCompany,
		ContactEmail:        res.Employer.ContactEmail,
		ContactPhoneNumber:  uint64(res.Employer.ContactPhoneNumber),
	}

	return &pb.EmployerSignupResponse{
		Status:          201,
		EmployerDetails: employerDetails,
		Token:           res.Token,
	}, nil
}

func (es *EmployerServer) EmployerLogin(ctx context.Context, req *pb.EmployerLoginInRequest) (*pb.EmployerLoginResponse, error) {
	employerLogin := models.EmployerLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	employer, err := es.employerUseCase.EmployerLogin(employerLogin)
	if err != nil {
		return &pb.EmployerLoginResponse{}, err
	}

	employerDetails := &pb.EmployerDetails{
		Id:                  uint64(employer.Employer.ID),
		CompanyName:         employer.Employer.CompanyName,
		Industry:            employer.Employer.Industry,
		CompanySize:         int32(employer.Employer.CompanySize),
		Website:             employer.Employer.Website,
		HeadquartersAddress: employer.Employer.HeadquartersAddress,
		AboutCompany:        employer.Employer.AboutCompany,
		ContactEmail:        employer.Employer.ContactEmail,
		ContactPhoneNumber:  uint64(employer.Employer.ContactPhoneNumber),
	}

	return &pb.EmployerLoginResponse{
		Status:          200,
		EmployerDetails: employerDetails,
		Token:           employer.Token,
	}, nil
}
