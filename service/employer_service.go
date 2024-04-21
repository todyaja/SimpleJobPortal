package service

import (
	"context"
	"simplejobportal/data/request"
	"simplejobportal/data/response"
)

type EmployerService interface {
	CreateJob(ctx context.Context, request request.JobCreateRequest, jwtToken string)
	SeeJobApplicant(ctx context.Context, request request.ViewApplicantRequest, jwtToken string) response.JobApplicantResponse
	ProcessApplicant(ctx context.Context, request request.ProcessApplicantRequest, jwtToken string) error
}
