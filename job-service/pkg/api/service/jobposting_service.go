package service

import (
	pb "Auth/pkg/pb/job"
	interfaces "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobServer struct {
	jobUseCase interfaces.JobUseCase
	pb.UnimplementedJobServer
}

func NewJobServer(useCase interfaces.JobUseCase) pb.JobServer {
	return &JobServer{
		jobUseCase: useCase,
	}
}
func (js *JobServer) PostJob(ctx context.Context, req *pb.JobOpeningRequest) (*pb.JobOpeningResponse, error) {
	// Extract EmployerID from context or request, assuming it's available in the request for now
	employerID := int32(req.EmployerId)

	jobDetails := models.JobOpening{
		Title:               req.Title,
		Description:         req.Description,
		Requirements:        req.Requirements,
		Location:            req.Location,
		EmploymentType:      req.EmploymentType,
		Salary:              req.Salary,
		SkillsRequired:      req.SkillsRequired,
		ExperienceLevel:     req.ExperienceLevel,
		EducationLevel:      req.EducationLevel,
		ApplicationDeadline: req.ApplicationDeadline.AsTime(),
	}

	fmt.Println("service", jobDetails)

	res, err := js.jobUseCase.PostJob(jobDetails, employerID)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	jobOpening := &pb.JobOpeningResponse{
		Id:                  uint64(res.ID),
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		PostedOn:            timestamppb.New(res.PostedOn),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              res.Salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
		EmployerId:          int32(req.EmployerId), // Set the EmployerId field
	}

	return jobOpening, nil
}

func (js *JobServer) GetAllJobs(ctx context.Context, req *pb.GetAllJobsRequest) (*pb.GetAllJobsResponse, error) {
	employerID := int32(req.EmployerIDInt)

	// Call the use case to get all jobs
	jobs, err := js.jobUseCase.GetAllJobs(employerID)
	if err != nil {
		return nil, err
	}

	// Convert jobs to protobuf response format
	var jobResponses []*pb.JobOpeningResponse
	for _, job := range jobs {
		jobResponse := &pb.JobOpeningResponse{
			Id:                  uint64(job.ID),
			Title:               job.Title,
			ApplicationDeadline: timestamppb.New(job.ApplicationDeadline),
			EmployerId:          job.EmployerID,
		}
		jobResponses = append(jobResponses, jobResponse)
	}

	// Create and return the response
	return &pb.GetAllJobsResponse{Jobs: jobResponses}, nil
}
func (js *JobServer) GetAJob(ctx context.Context, req *pb.GetAJobRequest) (*pb.JobOpeningResponse, error) {
	employerID := req.EmployerIDInt
	jobId := req.JobId

	res, err := js.jobUseCase.GetAJob(employerID, jobId)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	jobOpening := &pb.JobOpeningResponse{
		Id:                  uint64(res.ID),
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		PostedOn:            timestamppb.New(res.PostedOn),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              res.Salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
		EmployerId:          employerID, // Set the EmployerId field
	}

	return jobOpening, nil
}
