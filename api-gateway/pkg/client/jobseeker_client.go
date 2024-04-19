package client

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/config"
	pb "HireoGateWay/pkg/pb/jobseeker"
	"HireoGateWay/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type jobSeekerClient struct {
	Client pb.JobSeekerClient
}

func NewJobSeekerClient(cfg config.Config) interfaces.JobSeekerClient {
	grpcConnection, err := grpc.Dial(cfg.HireoAuth, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewJobSeekerClient(grpcConnection)

	return &jobSeekerClient{
		Client: grpcClient,
	}
}
func (jc *jobSeekerClient) JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.TokenJobSeeker, error) {
	jobSeeker, err := jc.Client.JobSeekerSignup(context.Background(), &pb.JobSeekerSignupRequest{
		Email:            jobSeekerDetails.Email,
		Password:         jobSeekerDetails.Password,
		FirstName:        jobSeekerDetails.FirstName,
		LastName:         jobSeekerDetails.LastName,
		PhoneNumber:      jobSeekerDetails.PhoneNumber,
		Address:          jobSeekerDetails.Address,
		DateOfBirth:      jobSeekerDetails.DateOfBirth,
		Gender:           jobSeekerDetails.Gender,
		Bio:              jobSeekerDetails.Bio,
		SocialMediaLinks: jobSeekerDetails.SocialMedia,
		Skills:           jobSeekerDetails.Skills,
		EducationHistory: jobSeekerDetails.Education,
		WorkExperience:   jobSeekerDetails.WorkExperience,
		LanguagesSpoken:  jobSeekerDetails.LanguagesSpoken,
		Interests:        jobSeekerDetails.Interests,
		Projects:         jobSeekerDetails.Projects,
	})
	if err != nil {
		return models.TokenJobSeeker{}, err
	}
	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerSignUpResponse{
			ID:              uint(jobSeeker.JobSeekerDetails.ID),
			Email:           jobSeeker.JobSeekerDetails.Email,
			FirstName:       jobSeeker.JobSeekerDetails.FirstName,
			LastName:        jobSeeker.JobSeekerDetails.LastName,
			PhoneNumber:     jobSeeker.JobSeekerDetails.PhoneNumber,
			Address:         jobSeeker.JobSeekerDetails.Address,
			DateOfBirth:     jobSeeker.JobSeekerDetails.DateOfBirth,
			Gender:          jobSeeker.JobSeekerDetails.Gender,
			Bio:             jobSeeker.JobSeekerDetails.Bio,
			SocialMedia:     jobSeeker.JobSeekerDetails.SocialMediaLinks,
			Skills:          jobSeeker.JobSeekerDetails.Skills,
			Education:       jobSeeker.JobSeekerDetails.EducationHistory,
			WorkExperience:  jobSeeker.JobSeekerDetails.WorkExperience,
			LanguagesSpoken: jobSeeker.JobSeekerDetails.LanguagesSpoken,
			Interests:       jobSeeker.JobSeekerDetails.Interests,
			Projects:        jobSeeker.JobSeekerDetails.Projects,
		},
		Token: jobSeeker.Token,
	}, nil
}

func (jc *jobSeekerClient) JobSeekerLogin(jobSeekerDetails models.JobSeekerLogin) (models.TokenJobSeeker, error) {
	jobSeeker, err := jc.Client.JobSeekerLogin(context.Background(), &pb.JobSeekerLoginRequest{
		Email:    jobSeekerDetails.Email,
		Password: jobSeekerDetails.Password,
	})

	if err != nil {
		return models.TokenJobSeeker{}, err
	}
	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerSignUpResponse{
			ID:              uint(jobSeeker.JobSeekerDetails.ID),
			Email:           jobSeeker.JobSeekerDetails.Email,
			FirstName:       jobSeeker.JobSeekerDetails.FirstName,
			LastName:        jobSeeker.JobSeekerDetails.LastName,
			PhoneNumber:     jobSeeker.JobSeekerDetails.PhoneNumber,
			Address:         jobSeeker.JobSeekerDetails.Address,
			DateOfBirth:     jobSeeker.JobSeekerDetails.DateOfBirth,
			Gender:          jobSeeker.JobSeekerDetails.Gender,
			Bio:             jobSeeker.JobSeekerDetails.Bio,
			SocialMedia:     jobSeeker.JobSeekerDetails.SocialMediaLinks,
			Skills:          jobSeeker.JobSeekerDetails.Skills,
			Education:       jobSeeker.JobSeekerDetails.EducationHistory,
			WorkExperience:  jobSeeker.JobSeekerDetails.WorkExperience,
			LanguagesSpoken: jobSeeker.JobSeekerDetails.LanguagesSpoken,
			Interests:       jobSeeker.JobSeekerDetails.Interests,
			Projects:        jobSeeker.JobSeekerDetails.Projects,
		},
		Token: jobSeeker.Token,
	}, nil
}
