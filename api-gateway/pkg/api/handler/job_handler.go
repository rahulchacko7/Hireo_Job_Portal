package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/logging"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	GRPC_Client interfaces.JobClient
}

func NewJobHandler(jobClient interfaces.JobClient) *JobHandler {
	return &JobHandler{
		GRPC_Client: jobClient,
	}
}
func (jh *JobHandler) PostJobOpening(c *gin.Context) {

	logEntry := logging.GetLogger().WithField("context", "PostJobOpening")
	logEntry.Info("Processing post job opening request")

	employerID, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Infof("Employer ID: %v", employerID)

	employerIDInt, ok := employerID.(int32)
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	var jobOpening models.JobOpening
	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		logEntry.Error("Details not in correct format")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	fmt.Println("id", employerIDInt, employerID)

	JobOpening, err := jh.GRPC_Client.PostJobOpening(jobOpening, employerIDInt)
	if err != nil {
		logEntry.Errorf("Failed to create job opening: %v", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to create job opening", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logEntry.Info("Job opening created successfully")
	response := response.ClientResponse(http.StatusCreated, "Job opening created successfully", JobOpening, nil)
	c.JSON(http.StatusCreated, response)
}

func (jh *JobHandler) GetAllJobs(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetAllJobs")
	logEntry.Info("Processing get all jobs request")

	employerID, ok := c.Get("id")
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Infof("Employer ID: %v", employerID)

	employerIDInt, ok := employerID.(int32)
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.GetAllJobs(employerIDInt)
	if err != nil {

		logEntry.Errorf("Failed to fetch jobs: %v", err)

		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Jobs retrieved successfully")
	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) GetAJob(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetAJob")
	logEntry.Info("Processing get a job request")

	idStr := c.Query("id")
	logEntry.Infof("Job ID: %v", idStr)

	employerID, ok := c.Get("id")
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logEntry.Errorf("Invalid job ID: %v", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.GetAJob(employerIDInt, int32(jobID))
	if err != nil {
		logEntry.Errorf("Failed to fetch jobs: %v", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Jobs retrieved successfully")
	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) DeleteAJob(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "DeleteAJob")
	logEntry.Info("Processing delete job request")

	idStr := c.Query("id")
	logEntry.Infof("Job ID: %s", idStr)

	employerID, ok := c.Get("id")
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logEntry.Errorf("Invalid job ID: %v", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = jh.GRPC_Client.DeleteAJob(employerIDInt, int32(jobID))
	if err != nil {
		logEntry.Errorf("Failed to delete job: %v", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Job deleted successfully")
	response := response.ClientResponse(http.StatusOK, "Job Deleted successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) UpdateAJob(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UpdateAJob")
	logEntry.Info("Processing update a job request")

	idStr := c.Query("id")
	logEntry.Infof("Job ID: %v", idStr)

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logEntry.Errorf("Invalid job ID: %v", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerID, ok := c.Get("id")
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	var jobOpening models.JobOpening
	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		logEntry.Errorf("Details not in correct format: %v", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	UpdateJobOpening, err := jh.GRPC_Client.UpdateAJob(employerIDInt, int32(jobID), jobOpening)
	if err != nil {
		logEntry.Errorf("Failed to update job: %v", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Job updated successfully")
	response := response.ClientResponse(http.StatusOK, "Job updated successfully", UpdateJobOpening, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) ViewAllJobs(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "ViewAllJobs")
	logEntry.Info("Processing view all jobs request")

	keyword := c.Query("Keyword")
	logEntry.Infof("Keyword: %v", keyword)

	if keyword == "" {
		logEntry.Error("Keyword parameter is required")
		errs := response.ClientResponse(http.StatusBadRequest, "Keyword parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.JobSeekerGetAllJobs(keyword)
	if err != nil {
		logEntry.Errorf("Failed to fetch jobs: %v", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	if len(jobs) == 0 {
		errMsg := "No jobs found matching your query"
		logEntry.Info(errMsg)
		errs := response.ClientResponse(http.StatusOK, errMsg, nil, nil)
		c.JSON(http.StatusOK, errs)
		return
	}
	logEntry.Info("Jobs retrieved successfully")
	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) GetJobDetails(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetJobDetails")
	logEntry.Info("Received request to get job details")

	idStr := c.Query("id")
	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logEntry.WithError(err).Error("Invalid job ID")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry = logEntry.WithField("jobID", jobID)

	jobDetails, err := jh.GRPC_Client.GetJobDetails(int32(jobID))
	if err != nil {
		logEntry.WithError(err).Error("Failed to fetch job details")
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch job details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Successfully retrieved job details")
	response := response.ClientResponse(http.StatusOK, "Job details retrieved successfully", jobDetails, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) ApplyJob(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "ApplyJob")
	logEntry.Info("Processing apply job request")

	employerID, ok := c.Get("id")
	if !ok {
		errMsg := "Invalid employer ID type"
		logEntry.Error(errMsg)
		errs := response.ClientResponse(http.StatusBadRequest, errMsg, nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, ok := employerID.(int32)
	if !ok {
		errMsg := "Invalid employer ID type"
		logEntry.Error(errMsg)
		errs := response.ClientResponse(http.StatusBadRequest, errMsg, nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	var jobApplication models.ApplyJob
	jobIDStr := c.PostForm("job_id")
	jobApplication.JobID, _ = strconv.ParseInt(jobIDStr, 10, 64)
	jobApplication.CoverLetter = c.PostForm("cover_letter")
	jobApplication.JobseekerID = int64(userIdInt)

	file, err := c.FormFile("resume")
	if err != nil {
		errMsg := "Error in getting resume file"
		logEntry.Error(errMsg)
		errorRes := response.ClientResponse(http.StatusBadRequest, errMsg, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	filePath := fmt.Sprintf("uploads/resumes/%d_%s", jobApplication.JobID, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		errMsg := "Failed to save resume file"
		logEntry.Error(errMsg)
		errorRes := response.ClientResponse(http.StatusInternalServerError, errMsg, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		errMsg := "Failed to read resume file"
		logEntry.Error(errMsg)
		errorRes := response.ClientResponse(http.StatusInternalServerError, errMsg, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	jobApplication.Resume = fileBytes

	logEntry.Info("Sending job application to gRPC client")
	res, err := jh.GRPC_Client.ApplyJob(jobApplication, file)
	if err != nil {
		errMsg := "Failed to apply for job"
		logEntry.Error(errMsg)
		errorRes := response.ClientResponse(http.StatusInternalServerError, errMsg, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	logEntry.Info("Job applied successfully")
	successRes := response.ClientResponse(http.StatusOK, "Job applied successfully", res, nil)
	c.JSON(http.StatusOK, successRes)
}

func (jh *JobHandler) GetApplicants(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetApplicants")
	logEntry.Info("Processing get applicants request")

	employerID, ok := c.Get("id")
	if !ok {
		logEntry.Warn("Failed to get employer ID from context")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, ok := employerID.(int32)
	if !ok {
		logEntry.Warnf("Invalid employer ID type: %T", employerID)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Infof("Fetching applicants for employer ID: %d", userIdInt)
	applicants, err := jh.GRPC_Client.GetApplicants(int64(userIdInt))
	if err != nil {
		logEntry.Errorf("Failed to fetch applicants for employer ID %d: %v", userIdInt, err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch applicants", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Infof("Successfully retrieved applicants for employer ID: %d", userIdInt)
	response := response.ClientResponse(http.StatusOK, "Applicants retrieved successfully", applicants, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) SaveAJob(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "SaveAJob")
	logEntry.Info("Processing save job request")

	jobIDStr := c.Query("job_id")
	jobIdInt, err := strconv.ParseInt(jobIDStr, 10, 32)
	if err != nil {
		logEntry.WithError(err).Error("Invalid or missing job ID")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid or missing job ID", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry = logEntry.WithField("job_id", jobIdInt)

	userID, userIDExists := c.Get("id")
	if !userIDExists {
		logEntry.Error("User ID not found")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, userIDOk := userID.(int32)
	if !userIDOk {
		logEntry.Error("Invalid user ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry = logEntry.WithField("user_id", userIdInt)

	Data, err := jh.GRPC_Client.SaveAJob(userIdInt, int32(jobIdInt))
	if err != nil {
		logEntry.WithError(err).Error("Failed to save job")
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to save job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Job saved successfully")
	response := response.ClientResponse(http.StatusOK, "Job saved successfully", Data, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) DeleteSavedJob(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "DeleteSavedJob")
	logEntry.Info("Processing delete saved job request")

	jobIDStr := c.Query("job_id")
	jobIdInt, err := strconv.ParseInt(jobIDStr, 10, 32)
	if err != nil {
		logEntry.WithError(err).Error("Invalid job ID format")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry = logEntry.WithField("job_id", jobIdInt)

	userID, userIDExists := c.Get("id")
	if !userIDExists {
		logEntry.Error("User ID not found in context")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid or missing user ID", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, userIDOk := userID.(int32)
	if !userIDOk {
		logEntry.Error("Invalid user ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	logEntry = logEntry.WithField("user_id", userIdInt)

	logEntry.Info("Calling GRPC client to delete saved job")
	err = jh.GRPC_Client.DeleteSavedJob(int32(jobIdInt), userIdInt)
	if err != nil {
		logEntry.WithError(err).Error("Failed to delete saved job")
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Job deleted successfully")
	response := response.ClientResponse(http.StatusOK, "Job deleted successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) GetASavedJob(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetASavedJob")
	logEntry.Info("Processing get saved job request")

	userID, userIDExists := c.Get("id")
	if !userIDExists {
		logEntry.Warn("User ID not found in context")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdInt, userIDOk := userID.(int32)
	if !userIDOk {
		logEntry.Warn("Invalid user ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry = logEntry.WithField("user_id", userIdInt)
	logEntry.Info("Fetching saved job for user")

	job, err := jh.GRPC_Client.GetASavedJob(userIdInt)
	if err != nil {
		logEntry.WithError(err).Error("Failed to get saved job")
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to get job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.WithField("job", job).Info("Saved job fetched successfully")
	response := response.ClientResponse(http.StatusOK, "Job fetched successfully", job, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobHandler) ApplySavedJob(c *gin.Context) {
	log.Println("ApplySavedJob: Handler started")

	userID, userIDExists := c.Get("id")
	userIdInt, userIDOk := userID.(int32)
	if !userIDExists || !userIDOk {
		errMsg := "Invalid or missing user ID"
		log.Println("ApplySavedJob: ", errMsg)
		errs := response.ClientResponse(http.StatusBadRequest, errMsg, nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobIDStr := c.PostForm("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		errMsg := "Invalid job ID"
		log.Println("ApplySavedJob: ", errMsg, " - Error: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, errMsg, nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	log.Printf("ApplySavedJob: userID: %d, jobID: %d", userIdInt, jobID)

	savedJobs, err := jh.GRPC_Client.GetASavedJob(userIdInt)
	if err != nil {
		errMsg := "Failed to check saved jobs"
		log.Println("ApplySavedJob: ", errMsg, " - Error: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, errMsg, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	jobIsSaved := false
	for _, savedJob := range savedJobs {
		if savedJob.JobID == jobID {
			jobIsSaved = true
			break
		}
	}

	if !jobIsSaved {
		errMsg := "No such saved job found"
		log.Println("ApplySavedJob: ", errMsg)
		errs := response.ClientResponse(http.StatusNotFound, errMsg, nil, nil)
		c.JSON(http.StatusNotFound, errs)
		return
	}

	var jobApplication models.ApplyJob
	jobApplication.JobID = jobID
	jobApplication.CoverLetter = c.PostForm("cover_letter")
	jobApplication.JobseekerID = int64(userIdInt)

	log.Println("ApplySavedJob: Preparing to receive resume file")

	file, err := c.FormFile("resume")
	if err != nil {
		errMsg := "Error in getting resume file"
		log.Println("ApplySavedJob: ", errMsg, " - Error: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, errMsg, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	filePath := fmt.Sprintf("uploads/resumes/%d_%s", jobApplication.JobID, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		errMsg := "Failed to save resume file"
		log.Println("ApplySavedJob: ", errMsg, " - Error: ", err)
		errorRes := response.ClientResponse(http.StatusInternalServerError, errMsg, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	log.Println("ApplySavedJob: Resume file saved at ", filePath)

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		errMsg := "Failed to read resume file"
		log.Println("ApplySavedJob: ", errMsg, " - Error: ", err)
		errorRes := response.ClientResponse(http.StatusInternalServerError, errMsg, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	jobApplication.Resume = fileBytes
	jobApplication.ResumeURL = filePath

	log.Println("ApplySavedJob: Applying for job with GRPC Client")

	res, err := jh.GRPC_Client.ApplyJob(jobApplication, file)
	if err != nil {
		errMsg := "Failed to apply for job"
		log.Println("ApplySavedJob: ", errMsg, " - Error: ", err)
		errorRes := response.ClientResponse(http.StatusInternalServerError, errMsg, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	log.Println("ApplySavedJob: Job applied successfully")

	successRes := response.ClientResponse(http.StatusOK, "Job applied successfully", res, nil)
	c.JSON(http.StatusOK, successRes)
}

func (jh *JobHandler) ScheduleInterview(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "ScheduleInterview")
	logEntry.Info("Processing schedule interview request")

	userID, userIDExists := c.Get("id")
	employerIDInt, userIDOk := userID.(int32)
	if !userIDExists || !userIDOk {
		logEntry.Warn("Invalid or missing user ID")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid or missing user ID", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobID, err := strconv.ParseInt(c.Query("job_id"), 10, 64)
	if err != nil {
		logEntry.WithError(err).Warn("Invalid job ID")
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	jobseekerID, err := strconv.ParseInt(c.Query("jobseeker_id"), 10, 64)
	if err != nil {
		logEntry.WithError(err).Warn("Invalid jobseeker ID")
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid jobseeker ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	interviewDate, err := time.Parse("2006-01-02", c.Query("interview_date"))
	if err != nil {
		logEntry.WithError(err).Warn("Invalid interview date")
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid interview date", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	interviewTime, err := time.Parse("15:04", c.Query("interview_time"))
	if err != nil {
		logEntry.WithError(err).Warn("Invalid interview time")
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid interview time", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	interviewLink := c.Query("link")
	interviewType := c.Query("interview_type")
	if interviewType != "ONLINE" && interviewType != "OFFLINE" {
		logEntry.Warn("Invalid interview type")
		errorRes := response.ClientResponse(http.StatusBadRequest, "Invalid interview type", nil, "Interview type must be ONLINE or OFFLINE")
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	scheduledTime := time.Date(
		interviewDate.Year(), interviewDate.Month(), interviewDate.Day(),
		interviewTime.Hour(), interviewTime.Minute(), 0, 0, time.UTC,
	)

	interview := models.Interview{
		JobID:         jobID,
		JobseekerID:   jobseekerID,
		EmployerID:    employerIDInt,
		ScheduledTime: scheduledTime,
		Mode:          interviewType,
		Link:          interviewLink,
		Status:        "SCHEDULED",
	}

	scheduledInterview, err := jh.GRPC_Client.ScheduleInterview(interview)
	if err != nil {
		logEntry.WithError(err).Error("Failed to schedule interview")
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to schedule interview", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	logEntry.Info("Interview scheduled successfully")
	successRes := response.ClientResponse(http.StatusOK, "Interview scheduled successfully", scheduledInterview, nil)
	c.JSON(http.StatusOK, successRes)
}

func (jh *JobHandler) GetInterviews(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetInterviews")
	logEntry.Info("Processing get interviews request")

	userID, userIDExists := c.Get("id")
	if !userIDExists {
		logEntry.Error("User ID not found in context")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid or missing user ID", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	employerID, ok := userID.(int32)
	if !ok {
		logEntry.Error("Invalid user ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid user ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobIDStr := c.Query("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 32)
	if err != nil {
		logEntry.WithField("job_id", jobIDStr).Error("Invalid job ID")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	getInterview, err := jh.GRPC_Client.GetInterview(int32(jobID), employerID)
	if err != nil {
		logEntry.WithError(err).Error("Failed to fetch interview details")
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch interview details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	logEntry.Info("Interview details fetched successfully")
	successRes := response.ClientResponse(http.StatusOK, "Interview details fetched successfully", getInterview, nil)
	c.JSON(http.StatusOK, successRes)
}
