package service

import (
	"context"
	"simplejobportal/data/request"
	"simplejobportal/data/response"
)

type TalentService interface {
	SeeJob(ctx context.Context, jwtToken string) []response.JobResponse
	ApplyJob(ctx context.Context, request request.ApplyJobRequest, jwtToken string)
	SeeApplicationDetail(ctx context.Context, applicationId int, jwtToken string) response.ApplicationDetailResponse
}
