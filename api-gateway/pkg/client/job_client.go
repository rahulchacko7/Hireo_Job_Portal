package client

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/config"
	pb "HireoGateWay/pkg/pb/job"
	"HireoGateWay/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type jobClient struct {
	Client pb.JobClient
}

// NewJobClient creates a new instance of JobClient.
func NewJobClient(cfg config.Config) interfaces.JobClient {
	grpcConnection, err := grpc.Dial(cfg.HireoJob, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewJobClient(grpcConnection)

	return &jobClient{
		Client: grpcClient,
	}
}

// PostJobOpening posts a job opening.
func (jc *jobClient) PostJobOpening(jobDetails models.JobOpening, EmployerID int) (models.JobOpeningResponse, error) {

	applicationDeadline := timestamppb.New(jobDetails.ApplicationDeadline)

	job, err := jc.Client.PostJob(context.Background(), &pb.JobOpeningRequest{
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: applicationDeadline,
		EmployerID:          uint32(EmployerID), 
	})
	if err != nil {
		return models.JobOpeningResponse{}, err
	}

	postedOnTime := job.PostedOn.AsTime()
	applicationDeadlineTime := job.ApplicationDeadline.AsTime()

	return models.JobOpeningResponse{
		ID:                  uint(job.Id),
		Title:               job.Title,
		Description:         job.Description,
		Requirements:        job.Requirements,
		PostedOn:            postedOnTime,
		Location:            job.Location,
		EmploymentType:      job.EmploymentType,
		Salary:              job.Salary,
		SkillsRequired:      job.SkillsRequired,
		ExperienceLevel:     job.ExperienceLevel,
		EducationLevel:      job.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          uint32(EmployerID),
	}, nil
}
