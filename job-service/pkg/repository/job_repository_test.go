package repository

import (
	"Auth/pkg/utils/models"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetAllJobs(t *testing.T) {
	type args struct {
		employerID int32
	}

	tests := []struct {
		name        string
		args        args
		mockData    func(mockSQL sqlmock.Sqlmock)
		expected    []models.AllJob
		expectedErr error
	}{
		{
			name: "Success: Get all jobs for employer",
			args: args{
				employerID: 1,
			},
			mockData: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `SELECT id, title, application_deadline, employer_id FROM job_opening_responses WHERE employer_id = \?`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "application_deadline", "employer_id"}).
						AddRow(1, "Job 1", time.Now(), 1).
						AddRow(2, "Job 2", time.Now(), 1))
			},

			expected: []models.AllJob{
				{ID: 1, Title: "Job 1", ApplicationDeadline: time.Now().Truncate(time.Second), EmployerID: 1},
				{ID: 2, Title: "Job 2", ApplicationDeadline: time.Now().Truncate(time.Second), EmployerID: 1},
			},
			expectedErr: nil,
		},
		{
			name: "Error: Unable to get jobs for employer",
			args: args{
				employerID: 2,
			},
			mockData: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `SELECT id, title, application_deadline, employer_id FROM job_opening_responses WHERE employer_id = \?`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs(2).
					WillReturnError(errors.New("error getting jobs"))
			},
			expected:    nil,
			expectedErr: errors.New("error getting jobs"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.mockData(mockSQL)

			jr := jobRepository{DB: gormDB}
			result, err := jr.GetAllJobs(tt.args.employerID)

			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
func TestGetAJob(t *testing.T) {
	type args struct {
		employerID int32
		jobID      int32
	}

	tests := []struct {
		name        string
		args        args
		mockData    func(mockSQL sqlmock.Sqlmock)
		expected    models.JobOpeningResponse
		expectedErr error
	}{
		{
			name: "Success: Get a job for employer",
			args: args{
				employerID: 1,
				jobID:      1,
			},
			mockData: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `SELECT id, title, application_deadline, employer_id FROM job_opening_responses WHERE id = ? AND employer_id = ?`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "application_deadline", "employer_id"}).
						AddRow(1, "Job 1", time.Now(), 1))
			},
			expected: models.JobOpeningResponse{
				ID:                  1,
				Title:               "Job 1",
				ApplicationDeadline: time.Now(),
				EmployerID:          1,
			},
			expectedErr: nil,
		},
		{
			name: "Error: Unable to get job for employer",
			args: args{
				employerID: 2,
				jobID:      1,
			},
			mockData: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `SELECT id, title, application_deadline, employer_id FROM job_opening_responses WHERE id = ? AND employer_id = ?`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs(1, 2).
					WillReturnError(errors.New("error getting job"))
			},
			expected:    models.JobOpeningResponse{},
			expectedErr: errors.New("error getting job"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.mockData(mockSQL)

			jr := jobRepository{DB: gormDB}
			result, err := jr.GetAJob(tt.args.employerID, tt.args.jobID)

			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
func TestDeleteAJob(t *testing.T) {
	mockErr := errors.New("error deleting job")

	tests := []struct {
		name         string
		employerID   int32
		jobID        int32
		mockData     func(mockSQL sqlmock.Sqlmock)
		expectedErr  error
		expectedStmt string
	}{
		{
			name:       "Success: Delete a job",
			employerID: 1,
			jobID:      1,
			mockData: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec(`DELETE FROM job_opening_responses WHERE id = ? AND employer_id = ?`).
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr:  nil,
			expectedStmt: "DELETE FROM job_opening_responses WHERE id = ? AND employer_id = ?",
		},
		{
			name:       "Error: Failed to delete job",
			employerID: 2,
			jobID:      1,
			mockData: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec(`DELETE FROM job_opening_responses WHERE id = ? AND employer_id = ?`).
					WithArgs(1, 2).
					WillReturnError(mockErr)
			},
			expectedErr:  fmt.Errorf("failed to delete job: %w", mockErr),
			expectedStmt: "DELETE FROM job_opening_responses WHERE id = ? AND employer_id = ?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.mockData(mockSQL)

			jr := jobRepository{DB: gormDB}
			err := jr.DeleteAJob(tt.employerID, tt.jobID)

			assert.Equal(t, tt.expectedErr, err)

			// Verify that all mock expectations were met
			if err := mockSQL.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
