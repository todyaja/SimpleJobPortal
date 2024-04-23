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
	DeleteJob(ctx context.Context, id int, employerId int) error
	UpdateJob(ctx context.Context, id int, employerId int, job model.Job) error
}
