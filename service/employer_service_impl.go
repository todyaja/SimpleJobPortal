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

type EmployerServiceImpl struct {
	EmployerRepository repository.EmployerRepository
}

// DeleteJob implements EmployerService.
func (b *EmployerServiceImpl) DeleteJob(ctx context.Context, id int, jwtToken string) error {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	employerId := claims["userId"].(string)
	userType := claims["userType"].(string)
	auth.IsEmployer(userType)

	employerIdInt, err := strconv.Atoi(employerId)

	helper.PanicIfError(err)
	err = b.EmployerRepository.DeleteJob(ctx, id, employerIdInt)

	return err
}

// UpdateJob implements EmployerService.
func (b *EmployerServiceImpl) UpdateJob(ctx context.Context, id int, request request.JobUpdateRequest, jwtToken string) error {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	employerId := claims["userId"].(string)
	userType := claims["userType"].(string)
	auth.IsEmployer(userType)

	employerIdInt, err := strconv.Atoi(employerId)

	helper.PanicIfError(err)

	job := model.Job{
		Title:       request.Title,
		Detail:      request.Detail,
		Requirement: request.Requirement,
	}
	err = b.EmployerRepository.UpdateJob(ctx, id, employerIdInt, job)

	return err
}

func NewEmployerServiceImpl(employerRepository repository.EmployerRepository) EmployerService {
	return &EmployerServiceImpl{EmployerRepository: employerRepository}
}

// ProcessApplicant implements EmployerService.
func (b *EmployerServiceImpl) ProcessApplicant(ctx context.Context, request request.ProcessApplicantRequest, jwtToken string) error {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	userType := claims["userType"].(string)
	userId := claims["userId"].(string)
	auth.IsEmployer(userType)

	userIdInt, _ := strconv.Atoi(userId)

	//To Do : Make sure that only job creator can change applicant status

	helper.PanicIfError(err)

	err = b.EmployerRepository.ProcessApplicant(ctx, request, userIdInt)
	return err

}

// SeeJobApplicant implements EmployerService.
func (b *EmployerServiceImpl) SeeJobApplicant(ctx context.Context, request request.ViewApplicantRequest, jwtToken string) response.JobApplicantResponse {
	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	userType := claims["userType"].(string)
	userId := claims["userId"].(string)
	userIdInt, _ := strconv.Atoi(userId)
	auth.IsEmployer(userType)

	jobApplicantResponse := response.JobApplicantResponse{}

	jobApplicants := b.EmployerRepository.ViewJobApplicant(ctx, request.JobId, userIdInt)

	if jobApplicants.JobId != 0 && jobApplicants.JobTitle != "" {
		jobApplicantResponse.JobId = jobApplicants.JobId
		jobApplicantResponse.JobTitle = jobApplicants.JobTitle

		// Map each field to array of applicants
		for _, applicant := range jobApplicants.Applicants {
			jobApplicantResponse.Applicants = append(jobApplicantResponse.Applicants, response.Applicant{
				UserId: applicant.UserId,
				Email:  applicant.Email,
				Status: applicant.Status,
			})
		}
	}

	return jobApplicantResponse

}

// Create implements EmployerService.
func (b *EmployerServiceImpl) CreateJob(ctx context.Context, request request.JobCreateRequest, jwtToken string) {

	claims, err := auth.ParseJWT(jwtToken)
	helper.PanicIfError(err)

	postedBy := claims["userId"].(string)
	userType := claims["userType"].(string)
	auth.IsEmployer(userType)

	postedByInt, err := strconv.Atoi(postedBy)

	helper.PanicIfError(err)

	job := model.Job{
		Title:       request.Title,
		Detail:      request.Detail,
		Requirement: request.Requirement,
		PostedBy:    postedByInt,
	}
	b.EmployerRepository.PostJob(ctx, job)
}
