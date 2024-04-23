package service

import (
	"context"
	"simplejobportal/auth"
	"simplejobportal/data/request"
	"simplejobportal/data/response"
	"simplejobportal/helper"
	"simplejobportal/model"
	"simplejobportal/repository"
	"strconv"
)

type TalentServiceImpl struct {
	TalentRepository repository.TalentRepository
}

func NewTalentServiceImpl(talentRepository repository.TalentRepository) TalentService {
	return &TalentServiceImpl{TalentRepository: talentRepository}
}

// WithdrawApplication implements TalentService.
func (b *TalentServiceImpl) WithdrawApplication(ctx context.Context, id int, jwtToken string) error {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	talentId := claims["userId"].(string)
	userType := claims["userType"].(string)
	auth.IsTalent(userType)

	talentIdInt, err := strconv.Atoi(talentId)

	helper.PanicIfError(err)
	err = b.TalentRepository.DeleteApplication(ctx, id, talentIdInt)

	return err
}

// ApplicationDetail implements TalentService.
func (b *TalentServiceImpl) SeeApplicationDetail(ctx context.Context, applicationId int, jwtToken string) response.ApplicationDetailResponse {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	userType := claims["userType"].(string)
	auth.IsTalent(userType)

	applicationDetail, err := b.TalentRepository.ViewApplicationDetail(ctx, applicationId)
	helper.PanicIfError(err)

	applicationDetailResponse := response.ApplicationDetailResponse{Id: applicationDetail.Id, JobId: applicationDetail.JobId, JobTitle: applicationDetail.JobTitle, Status: applicationDetail.Status}

	return applicationDetailResponse
}

// SeeJob implements TalentService.
func (b *TalentServiceImpl) SeeJob(ctx context.Context, jwtToken string) []response.JobResponse {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	userType := claims["userType"].(string)
	auth.IsTalent(userType)

	jobs := b.TalentRepository.ViewJob(ctx)

	var jobResponse []response.JobResponse

	for _, job := range jobs {
		jobResponse = append(jobResponse, response.JobResponse{
			Id:          job.Id,
			Title:       job.Title,
			Detail:      job.Detail,
			Requirement: job.Requirement,
			PostedBy:    job.PostedBy,
		})
	}
	return jobResponse
}

func (b *TalentServiceImpl) ApplyJob(ctx context.Context, request request.ApplyJobRequest, jwtToken string) {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	applicantId := claims["userId"].(string)
	userType := claims["userType"].(string)
	auth.IsTalent(userType)

	applicantIdInt, err := strconv.Atoi(applicantId)

	helper.PanicIfError(err)

	application := model.Application{
		JobId:       request.JobId,
		ApplicantId: applicantIdInt,
		Status:      1,
	}
	b.TalentRepository.ApplyJob(ctx, application)
}
