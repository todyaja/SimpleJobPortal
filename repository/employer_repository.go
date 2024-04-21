package repository

import (
	"context"
	"simplejobportal/data/request"
	"simplejobportal/model"
)

type EmployerRepository interface {
	PostJob(ctx context.Context, job model.Job)
	ViewJobApplicant(ctx context.Context, jobId int, postedBy int) model.JobApplicant
	ProcessApplicant(ctx context.Context, request request.ProcessApplicantRequest, processedBy int) error
}
