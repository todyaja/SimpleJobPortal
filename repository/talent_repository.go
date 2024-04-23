package repository

import (
	"context"
	"simplejobportal/model"
)

type TalentRepository interface {
	ViewJob(ctx context.Context) []model.Job
	ApplyJob(ctx context.Context, application model.Application)
	ViewApplicationDetail(ctx context.Context, applicationId int) (model.ApplicationDetail, error)
	DeleteApplication(ctx context.Context, id int, talentId int) error
}
