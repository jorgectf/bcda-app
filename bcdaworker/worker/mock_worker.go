// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package worker

import (
	context "context"

	models "github.com/CMSgov/bcda-app/bcda/models"
	mock "github.com/stretchr/testify/mock"
)

// MockWorker is an autogenerated mock type for the Worker type
type MockWorker struct {
	mock.Mock
}

// ProcessJob provides a mock function with given fields: ctx, job, jobArgs
func (_m *MockWorker) ProcessJob(ctx context.Context, job models.Job, jobArgs models.JobEnqueueArgs) error {
	ret := _m.Called(ctx, job, jobArgs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Job, models.JobEnqueueArgs) error); ok {
		r0 = rf(ctx, job, jobArgs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateJob provides a mock function with given fields: ctx, jobArgs
func (_m *MockWorker) ValidateJob(ctx context.Context, jobArgs models.JobEnqueueArgs) (*models.Job, error) {
	ret := _m.Called(ctx, jobArgs)

	var r0 *models.Job
	if rf, ok := ret.Get(0).(func(context.Context, models.JobEnqueueArgs) *models.Job); ok {
		r0 = rf(ctx, jobArgs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.JobEnqueueArgs) error); ok {
		r1 = rf(ctx, jobArgs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
