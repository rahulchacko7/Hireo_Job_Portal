package service

import (
	pb "Auth/pkg/pb/job"
	interfaces "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

	jobs, err := js.jobUseCase.GetAllJobs(employerID)
	if err != nil {
		return nil, err
	}

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

	return &pb.GetAllJobsResponse{Jobs: jobResponses}, nil
}

func (js *JobServer) GetAJob(ctx context.Context, req *pb.GetAJobRequest) (*pb.JobOpeningResponse, error) {
	employerID := req.EmployerIDInt
	jobId := req.JobId

	res, err := js.jobUseCase.GetAJob(employerID, jobId)
	if err != nil {
		return nil, err
	}

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
		EmployerId:          employerID,
	}

	return jobOpening, nil
}

func (js *JobServer) DeleteAJob(ctx context.Context, req *pb.DeleteAJobRequest) (*emptypb.Empty, error) {
	employerID := req.EmployerIDInt
	jobID := req.JobId

	err := js.jobUseCase.DeleteAJob(employerID, jobID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete job: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (js *JobServer) UpdateAJob(ctx context.Context, req *pb.UpdateAJobRequest) (*pb.UpdateAJobResponse, error) {
	employerID := req.EmployerIDInt
	jobID := req.JobId

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

	res, err := js.jobUseCase.UpdateAJob(employerID, jobID, jobDetails)
	if err != nil {
		return nil, err
	}

	updateResponse := &pb.UpdateAJobResponse{
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
		EmployerId:          employerID,
	}

	return updateResponse, nil
}

func (js *JobServer) JobSeekerGetAllJobs(ctx context.Context, req *pb.JobSeekerGetAllJobsRequest) (*pb.JobSeekerGetAllJobsResponse, error) {
	keyword := req.Title

	jobSeekerJobs, err := js.jobUseCase.JobSeekerGetAllJobs(keyword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get jobs for job seeker: %v", err)
	}

	var jobsResponse []*pb.JSGetAllJobsRespons
	for _, job := range jobSeekerJobs {
		jobResponse := &pb.JSGetAllJobsRespons{
			Id:    uint64(job.ID),
			Title: job.Title,
		}
		jobsResponse = append(jobsResponse, jobResponse)
	}

	response := &pb.JobSeekerGetAllJobsResponse{
		Jobs: jobsResponse,
	}

	return response, nil
}

func (js *JobServer) GetJobDetails(ctx context.Context, req *pb.GetJobDetailsRequest) (*pb.GetJobDetailsResponse, error) {
	jobId := req.JobId

	res, err := js.jobUseCase.GetJobDetails(jobId)
	if err != nil {
		return nil, err
	}

	jobDetailsResponse := &pb.GetJobDetailsResponse{
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		EmployerId:          int32(res.EmployerID),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              res.Salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
	}

	return jobDetailsResponse, nil
}

func (js *JobServer) ApplyJob(ctx context.Context, req *pb.ApplyJobRequest) (*pb.ApplyJobResponse, error) {
	fmt.Println("Applying for job...")

	jobApplication := models.ApplyJob{
		JobID:       req.JobId,
		JobseekerID: req.JobseekerId,
		CoverLetter: req.CoverLetter,
		Resume:      req.ResumeData,
	}

	Data, err := js.jobUseCase.ApplyJob(jobApplication, req.ResumeData)
	if err != nil {
		return nil, err
	}

	return &pb.ApplyJobResponse{
		Id:          int64(Data.ID),
		JobId:       Data.JobID,
		JobseekerId: Data.JobseekerID,
		CoverLetter: Data.CoverLetter,
		ResumeUrl:   Data.ResumeURL,
	}, nil
}

func (js *JobServer) GetJobApplications(ctx context.Context, req *pb.GetJobApplicationsRequest) (*pb.GetJobApplicationsResponse, error) {
	employerID, err := strconv.ParseInt(req.EmployerId, 10, 64)
	if err != nil {
		return nil, err
	}

	applications, err := js.jobUseCase.GetApplicants(employerID)
	if err != nil {
		return nil, err
	}

	var applicationResponses []*pb.JobApplication
	for _, application := range applications {
		applicationResponse := &pb.JobApplication{
			Id:          strconv.FormatUint(uint64(application.ID), 10),
			JobId:       strconv.FormatUint(uint64(application.JobID), 10),
			JobSeekerId: strconv.FormatUint(uint64(application.JobseekerID), 10),
			Resume:      application.ResumeURL,
			CoverLetter: application.CoverLetter,
		}
		applicationResponses = append(applicationResponses, applicationResponse)
	}

	return &pb.GetJobApplicationsResponse{JobApplications: applicationResponses}, nil
}

func (js *JobServer) SaveJobs(ctx context.Context, req *pb.SaveJobRequest) (*pb.SaveJobResponse, error) {
	JobID, err := strconv.ParseInt(req.JobId, 10, 64)
	if err != nil {
		return nil, err
	}

	UserID, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	savedJob, err := js.jobUseCase.SaveJobs(JobID, UserID)
	if err != nil {
		return nil, err
	}

	response := &pb.SaveJobResponse{
		Id:      strconv.FormatUint(uint64(savedJob.ID), 10),
		JobId:   strconv.FormatInt(savedJob.JobID, 10),
		UserId:  strconv.FormatInt(savedJob.JobseekerID, 10),
		Message: "Job saved successfully",
	}

	return response, nil
}
