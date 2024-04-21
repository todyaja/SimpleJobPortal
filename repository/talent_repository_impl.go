package repository

import (
	"context"
	"database/sql"
	"errors"
	"simplejobportal/helper"
	"simplejobportal/model"
)

type TalentRepositoryImpl struct {
	Db *sql.DB
}

func NewTalentRepository(Db *sql.DB) TalentRepository {
	return &TalentRepositoryImpl{Db: Db}
}

// ViewApplicationDetail implements TalentRepository.
func (b *TalentRepositoryImpl) ViewApplicationDetail(ctx context.Context, applicationId int) (model.ApplicationDetail, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := `SELECT
			a.id,
			b.id,
			b.title,
			c.description
			FROM public.application a
			JOIN public.job b ON a.job_id = b.id
			JOIN public.application_status c ON a.status = c.id
			WHERE a.id = $1`

	result, errQuery := tx.QueryContext(ctx, SQL, applicationId)
	helper.PanicIfError(errQuery)

	defer result.Close()

	applicationDetail := model.ApplicationDetail{}

	if result.Next() {
		err := result.Scan(&applicationDetail.Id, &applicationDetail.JobId, &applicationDetail.JobTitle, &applicationDetail.Status)
		helper.PanicIfError(err)
		return applicationDetail, err
	} else {
		return applicationDetail, errors.New("application not found")
	}
}

// ApplyJob implements TalentRepository.
func (b *TalentRepositoryImpl) ApplyJob(ctx context.Context, application model.Application) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	SQL := "INSERT INTO public.application(job_id, applicant_id, status) VALUES ($1, $2, $3)"
	_, err = tx.ExecContext(ctx, SQL, &application.JobId, &application.ApplicantId, &application.Status)
	helper.PanicIfError(err)
}

// Save implements TalentRepository.
func (b *TalentRepositoryImpl) ViewJob(ctx context.Context) []model.Job {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, title, detail, requirement, posted_by FROM job"

	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)

	defer result.Close()

	var jobs []model.Job

	for result.Next() {
		job := model.Job{}
		err := result.Scan(&job.Id, &job.Title, &job.Detail, &job.Requirement, &job.PostedBy)
		helper.PanicIfError(err)

		jobs = append(jobs, job)
	}

	return jobs
}
