package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"simplejobportal/helper"
	"simplejobportal/model"
)

type TalentRepositoryImpl struct {
	Db *sql.DB
}

func NewTalentRepository(Db *sql.DB) TalentRepository {
	return &TalentRepositoryImpl{Db: Db}
}

// DeleteApplication implements TalentRepository.
func (b *TalentRepositoryImpl) DeleteApplication(ctx context.Context, id int, talentId int) error {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	//Validate talent id is same as applicant id
	SQL := `SELECT
			applicant_id
			FROM public.application a
			WHERE a.id = $1`

	result, errQuery := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(errQuery)

	defer result.Close()

	var applicantId int

	for result.Next() {
		err := result.Scan(&applicantId)
		helper.PanicIfError(err)
	}

	fmt.Println(id)
	fmt.Println(applicantId)
	fmt.Println(talentId)

	if applicantId != talentId {
		return errors.New("you are not authorized to withdraw this application")
	}

	//Remove from application table
	SQL = "DELETE FROM public.application WHERE ID = $1"

	_, errExec := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(errExec)

	return nil
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
