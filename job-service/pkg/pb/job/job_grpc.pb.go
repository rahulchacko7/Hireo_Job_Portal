// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: pkg/pb/job/job.proto

package job

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Job_PostJob_FullMethodName             = "/job.Job/PostJob"
	Job_GetAllJobs_FullMethodName          = "/job.Job/GetAllJobs"
	Job_GetAJob_FullMethodName             = "/job.Job/GetAJob"
	Job_DeleteAJob_FullMethodName          = "/job.Job/DeleteAJob"
	Job_UpdateAJob_FullMethodName          = "/job.Job/UpdateAJob"
	Job_JobSeekerGetAllJobs_FullMethodName = "/job.Job/JobSeekerGetAllJobs"
	Job_GetJobDetails_FullMethodName       = "/job.Job/GetJobDetails"
	Job_ApplyJob_FullMethodName            = "/job.Job/ApplyJob"
	Job_GetJobApplications_FullMethodName  = "/job.Job/GetJobApplications"
)

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	PostJob(ctx context.Context, in *JobOpeningRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error)
	GetAllJobs(ctx context.Context, in *GetAllJobsRequest, opts ...grpc.CallOption) (*GetAllJobsResponse, error)
	GetAJob(ctx context.Context, in *GetAJobRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error)
	DeleteAJob(ctx context.Context, in *DeleteAJobRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateAJob(ctx context.Context, in *UpdateAJobRequest, opts ...grpc.CallOption) (*UpdateAJobResponse, error)
	JobSeekerGetAllJobs(ctx context.Context, in *JobSeekerGetAllJobsRequest, opts ...grpc.CallOption) (*JobSeekerGetAllJobsResponse, error)
	GetJobDetails(ctx context.Context, in *GetJobDetailsRequest, opts ...grpc.CallOption) (*GetJobDetailsResponse, error)
	ApplyJob(ctx context.Context, in *ApplyJobRequest, opts ...grpc.CallOption) (*ApplyJobResponse, error)
	GetJobApplications(ctx context.Context, in *GetJobApplicationsRequest, opts ...grpc.CallOption) (*GetJobApplicationsResponse, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) PostJob(ctx context.Context, in *JobOpeningRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error) {
	out := new(JobOpeningResponse)
	err := c.cc.Invoke(ctx, Job_PostJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetAllJobs(ctx context.Context, in *GetAllJobsRequest, opts ...grpc.CallOption) (*GetAllJobsResponse, error) {
	out := new(GetAllJobsResponse)
	err := c.cc.Invoke(ctx, Job_GetAllJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetAJob(ctx context.Context, in *GetAJobRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error) {
	out := new(JobOpeningResponse)
	err := c.cc.Invoke(ctx, Job_GetAJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) DeleteAJob(ctx context.Context, in *DeleteAJobRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Job_DeleteAJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) UpdateAJob(ctx context.Context, in *UpdateAJobRequest, opts ...grpc.CallOption) (*UpdateAJobResponse, error) {
	out := new(UpdateAJobResponse)
	err := c.cc.Invoke(ctx, Job_UpdateAJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) JobSeekerGetAllJobs(ctx context.Context, in *JobSeekerGetAllJobsRequest, opts ...grpc.CallOption) (*JobSeekerGetAllJobsResponse, error) {
	out := new(JobSeekerGetAllJobsResponse)
	err := c.cc.Invoke(ctx, Job_JobSeekerGetAllJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetJobDetails(ctx context.Context, in *GetJobDetailsRequest, opts ...grpc.CallOption) (*GetJobDetailsResponse, error) {
	out := new(GetJobDetailsResponse)
	err := c.cc.Invoke(ctx, Job_GetJobDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) ApplyJob(ctx context.Context, in *ApplyJobRequest, opts ...grpc.CallOption) (*ApplyJobResponse, error) {
	out := new(ApplyJobResponse)
	err := c.cc.Invoke(ctx, Job_ApplyJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetJobApplications(ctx context.Context, in *GetJobApplicationsRequest, opts ...grpc.CallOption) (*GetJobApplicationsResponse, error) {
	out := new(GetJobApplicationsResponse)
	err := c.cc.Invoke(ctx, Job_GetJobApplications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility
type JobServer interface {
	PostJob(context.Context, *JobOpeningRequest) (*JobOpeningResponse, error)
	GetAllJobs(context.Context, *GetAllJobsRequest) (*GetAllJobsResponse, error)
	GetAJob(context.Context, *GetAJobRequest) (*JobOpeningResponse, error)
	DeleteAJob(context.Context, *DeleteAJobRequest) (*empty.Empty, error)
	UpdateAJob(context.Context, *UpdateAJobRequest) (*UpdateAJobResponse, error)
	JobSeekerGetAllJobs(context.Context, *JobSeekerGetAllJobsRequest) (*JobSeekerGetAllJobsResponse, error)
	GetJobDetails(context.Context, *GetJobDetailsRequest) (*GetJobDetailsResponse, error)
	ApplyJob(context.Context, *ApplyJobRequest) (*ApplyJobResponse, error)
	GetJobApplications(context.Context, *GetJobApplicationsRequest) (*GetJobApplicationsResponse, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (UnimplementedJobServer) PostJob(context.Context, *JobOpeningRequest) (*JobOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostJob not implemented")
}
func (UnimplementedJobServer) GetAllJobs(context.Context, *GetAllJobsRequest) (*GetAllJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllJobs not implemented")
}
func (UnimplementedJobServer) GetAJob(context.Context, *GetAJobRequest) (*JobOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAJob not implemented")
}
func (UnimplementedJobServer) DeleteAJob(context.Context, *DeleteAJobRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAJob not implemented")
}
func (UnimplementedJobServer) UpdateAJob(context.Context, *UpdateAJobRequest) (*UpdateAJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAJob not implemented")
}
func (UnimplementedJobServer) JobSeekerGetAllJobs(context.Context, *JobSeekerGetAllJobsRequest) (*JobSeekerGetAllJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerGetAllJobs not implemented")
}
func (UnimplementedJobServer) GetJobDetails(context.Context, *GetJobDetailsRequest) (*GetJobDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobDetails not implemented")
}
func (UnimplementedJobServer) ApplyJob(context.Context, *ApplyJobRequest) (*ApplyJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyJob not implemented")
}
func (UnimplementedJobServer) GetJobApplications(context.Context, *GetJobApplicationsRequest) (*GetJobApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobApplications not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_PostJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).PostJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_PostJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).PostJob(ctx, req.(*JobOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetAllJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetAllJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_GetAllJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetAllJobs(ctx, req.(*GetAllJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_GetAJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetAJob(ctx, req.(*GetAJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_DeleteAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).DeleteAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_DeleteAJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).DeleteAJob(ctx, req.(*DeleteAJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_UpdateAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).UpdateAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_UpdateAJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).UpdateAJob(ctx, req.(*UpdateAJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_JobSeekerGetAllJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerGetAllJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).JobSeekerGetAllJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_JobSeekerGetAllJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).JobSeekerGetAllJobs(ctx, req.(*JobSeekerGetAllJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetJobDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetJobDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_GetJobDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetJobDetails(ctx, req.(*GetJobDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_ApplyJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).ApplyJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_ApplyJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).ApplyJob(ctx, req.(*ApplyJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetJobApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetJobApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_GetJobApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetJobApplications(ctx, req.(*GetJobApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostJob",
			Handler:    _Job_PostJob_Handler,
		},
		{
			MethodName: "GetAllJobs",
			Handler:    _Job_GetAllJobs_Handler,
		},
		{
			MethodName: "GetAJob",
			Handler:    _Job_GetAJob_Handler,
		},
		{
			MethodName: "DeleteAJob",
			Handler:    _Job_DeleteAJob_Handler,
		},
		{
			MethodName: "UpdateAJob",
			Handler:    _Job_UpdateAJob_Handler,
		},
		{
			MethodName: "JobSeekerGetAllJobs",
			Handler:    _Job_JobSeekerGetAllJobs_Handler,
		},
		{
			MethodName: "GetJobDetails",
			Handler:    _Job_GetJobDetails_Handler,
		},
		{
			MethodName: "ApplyJob",
			Handler:    _Job_ApplyJob_Handler,
		},
		{
			MethodName: "GetJobApplications",
			Handler:    _Job_GetJobApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/job/job.proto",
}
