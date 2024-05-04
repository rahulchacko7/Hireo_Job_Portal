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
		EmployerID:          EmployerID, // Uncomment this line if you need to set EmployerID
	}, nil
}

func (jc *jobClient) GetAllJobs(employerIDInt int32) ([]models.AllJob, error) {
	// Make the gRPC call to get all jobs
	resp, err := jc.Client.GetAllJobs(context.Background(), &pb.GetAllJobsRequest{EmployerIDInt: employerIDInt})
	if err != nil {
		return nil, fmt.Errorf("failed to get all jobs: %v", err)
	}

	// Convert gRPC response to models.AllJob
	var allJobs []models.AllJob
	for _, job := range resp.Jobs {
		// Convert timestamp fields to Go time.Time
		applicationDeadlineTime := job.ApplicationDeadline.AsTime()

		allJobs = append(allJobs, models.AllJob{
			ID:                  uint(job.Id),
			Title:               job.Title,
			ApplicationDeadline: applicationDeadlineTime,
			EmployerID:          employerIDInt, // Assuming the employer ID is the same for all jobs
		})
	}

	return allJobs, nil
}

func (jc *jobClient) GetAJob(employerIDInt, jobId int32) (models.JobOpeningResponse, error) {
	resp, err := jc.Client.GetAJob(context.Background(), &pb.GetAJobRequest{EmployerIDInt: employerIDInt, JobId: jobId})
	if err != nil {
		return models.JobOpeningResponse{}, fmt.Errorf("failed to get job: %v", err)
	}

	postedOnTime := resp.PostedOn.AsTime()
	applicationDeadlineTime := resp.ApplicationDeadline.AsTime()

	return models.JobOpeningResponse{
		ID:                  uint(resp.Id),
		Title:               resp.Title,
		Description:         resp.Description,
		Requirements:        resp.Requirements,
		PostedOn:            postedOnTime,
		Location:            resp.Location,
		EmploymentType:      resp.EmploymentType,
		Salary:              resp.Salary,
		SkillsRequired:      resp.SkillsRequired,
		ExperienceLevel:     resp.ExperienceLevel,
		EducationLevel:      resp.EducationLevel,
		ApplicationDeadline: applicationDeadlineTime,
		EmployerID:          employerIDInt,
	}, nil
}
