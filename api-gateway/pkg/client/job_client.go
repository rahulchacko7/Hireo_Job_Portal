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
func (jc *jobClient) PostJobOpening(jobDetails models.JobOpening, EmployerID int32) (models.JobOpeningResponse, error) {
	// Create a timestamp for the application deadline
	applicationDeadline := timestamppb.New(jobDetails.ApplicationDeadline)

	// Make the gRPC call to post the job opening
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
		EmployerId:          EmployerID,
	})
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to post job opening: %v", err)
	}

	// Convert timestamp fields to Go time.Time
	postedOnTime := job.PostedOn.AsTime()
	applicationDeadlineTime := job.ApplicationDeadline.AsTime()

	// Construct the response
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
		//EmployerID:          EmployerID, // Uncomment this line if you need to set EmployerID
	}, nil
}
