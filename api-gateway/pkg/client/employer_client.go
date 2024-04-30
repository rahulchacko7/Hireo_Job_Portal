package client

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/config"
	pb "HireoGateWay/pkg/pb/auth"
	"HireoGateWay/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type employerClient struct {
	Client pb.EmployerClient
}

func NewEmployerClient(cfg config.Config) interfaces.EmployerClient {
	grpcConnection, err := grpc.Dial(cfg.HireoAuth, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewEmployerClient(grpcConnection)

	return &employerClient{
		Client: grpcClient,
	}
}

func (ec *employerClient) EmployerSignUp(employerDetails models.EmployerSignUp) (models.TokenEmployer, error) {
	employer, err := ec.Client.EmployerSignup(context.Background(), &pb.EmployerSignupRequest{
		CompanyName:         employerDetails.Company_name,
		Industry:            employerDetails.Industry,
		CompanySize:         int32(employerDetails.Company_size),
		Website:             employerDetails.Website,
		HeadquartersAddress: employerDetails.Headquarters_address,
		AboutCompany:        employerDetails.About_company,
		ContactEmail:        employerDetails.Contact_email,
		ContactPhoneNumber:  uint64(employerDetails.Contact_phone_number),
		Password:            employerDetails.Password,
	})
	if err != nil {
		return models.TokenEmployer{}, err
	}
	return models.TokenEmployer{
		Employer: models.EmployerDetailsResponse{
			ID:                   uint(employer.EmployerDetails.Id),
			Company_name:         employer.EmployerDetails.CompanyName,
			Industry:             employer.EmployerDetails.Industry,
			Company_size:         int(employer.EmployerDetails.CompanySize),
			Website:              employer.EmployerDetails.Website,
			Headquarters_address: employer.EmployerDetails.HeadquartersAddress,
			About_company:        employer.EmployerDetails.AboutCompany,
			Contact_email:        employer.EmployerDetails.ContactEmail,
			Contact_phone_number: uint(employer.EmployerDetails.ContactPhoneNumber),
		},
		Token: employer.Token,
	}, nil
}

func (ec *employerClient) EmployerLogin(employerDetails models.EmployerLogin) (models.TokenEmployer, error) {
	employer, err := ec.Client.EmployerLogin(context.Background(), &pb.EmployerLoginInRequest{
		Email:    employerDetails.Email,
		Password: employerDetails.Password,
	})

	if err != nil {
		return models.TokenEmployer{}, err
	}
	return models.TokenEmployer{
		Employer: models.EmployerDetailsResponse{
			ID:                   uint(employer.EmployerDetails.Id),
			Company_name:         employer.EmployerDetails.CompanyName,
			Industry:             employer.EmployerDetails.Industry,
			Company_size:         int(employer.EmployerDetails.CompanySize),
			Website:              employer.EmployerDetails.Website,
			Headquarters_address: employer.EmployerDetails.HeadquartersAddress,
			About_company:        employer.EmployerDetails.AboutCompany,
			Contact_email:        employer.EmployerDetails.ContactEmail,
			Contact_phone_number: uint(employer.EmployerDetails.ContactPhoneNumber),
		},
		Token: employer.Token,
	}, nil
}
