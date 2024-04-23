package repository

import (
	"context"
	"database/sql"
	"errors"
	"simplejobportal/data/request"
	"simplejobportal/helper"
	"simplejobportal/model"
)

type EmployerRepositoryImpl struct {
	Db *sql.DB
}

// UpdateJob implements EmployerRepository.
func (b *EmployerRepositoryImpl) UpdateJob(ctx context.Context, id int, employerId int, job model.Job) error {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	//Validate job creator same as the one who is going to delete job
	SQL := `SELECT posted_by FROM public.job a
			WHERE a.id = $1`

	result, errQuery := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(errQuery)

	defer result.Close()

	var jobCreatorId int
	for result.Next() {
		err := result.Scan(&jobCreatorId)
		helper.PanicIfError(err)
	}

	if jobCreatorId != employerId {
		return errors.New("you are not authorized to update this job")
	}
	SQL = "UPDATE public.job SET title = $1, detail = $2, requirement = $3 WHERE id  = $4"
	_, err = tx.ExecContext(ctx, SQL, job.Title, job.Detail, job.Requirement, id)
	helper.PanicIfError(err)

	return nil
}

// DeleteJob implements EmployerRepository.
func (b *EmployerRepositoryImpl) DeleteJob(ctx context.Context, id int, employerId int) error {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	//Validate job creator same as the one who is going to delete job
	SQL := `SELECT posted_by FROM public.job a
			WHERE a.id = $1`

	result, errQuery := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(errQuery)

	defer result.Close()

	var jobCreatorId int
	for result.Next() {
		err := result.Scan(&jobCreatorId)
		helper.PanicIfError(err)
	}

	if jobCreatorId != employerId {
		return errors.New("you are not authorized to process this applicant")
	}
	SQL = "DELETE FROM public.job WHERE ID = $1"

	_, errExec := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(errExec)
	return nil

}

func NewEmployerRepository(Db *sql.DB) EmployerRepository {
	return &EmployerRepositoryImpl{Db: Db}
}

// ProcessApplicant implements EmployerRepository.
func (b *EmployerRepositoryImpl) ProcessApplicant(ctx context.Context, request request.ProcessApplicantRequest, processedBy int) error {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	//Validate job creator same as the one who is going to process applicant
	SQL := `SELECT posted_by FROM public.application a
			JOIN public.job b ON a.job_id = b.id
			WHERE a.id = 1`

	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)

	defer result.Close()

	var jobCreatorId int
	for result.Next() {
		err := result.Scan(&jobCreatorId)
		helper.PanicIfError(err)
	}

	if jobCreatorId != processedBy {
		return errors.New("you are not authorized to process this applicant")
	}

	SQL = "UPDATE public.application SET status = $1 WHERE id = $2"
	_, err = tx.ExecContext(ctx, SQL, request.ChangeStatusTo, request.AppplicationId)
	helper.PanicIfError(err)
	return nil
}

// ViewJobApplicant implements EmployerRepository.
func (b *EmployerRepositoryImpl) ViewJobApplicant(ctx context.Context, jobId int, postedBy int) model.JobApplicant {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := `SELECT 
			b.id,
			b.title,
			c.id,
			c.email,
			d.description
			FROM public.application a
			JOIN public.job b ON a.job_id = b.id
			JOIN public.user c ON a.applicant_id = c.id
			JOIN public.application_status d ON a.status = d.id
			WHERE job_id = $1 
			AND b.posted_by = $2`

	result, errQuery := tx.QueryContext(ctx, SQL, jobId, postedBy)
	helper.PanicIfError(errQuery)

	defer result.Close()

	jobApplicant := model.JobApplicant{}

	applicants := []model.Applicant{}

	for result.Next() {
		applicant := model.Applicant{}
		err := result.Scan(&jobApplicant.JobId, &jobApplicant.JobTitle, &applicant.UserId, &applicant.Email, &applicant.Status)
		helper.PanicIfError(err)
		applicants = append(applicants, applicant)
	}

	jobApplicant.Applicants = applicants
	return jobApplicant
}

// Save implements EmployerRepository.
func (b *EmployerRepositoryImpl) PostJob(ctx context.Context, job model.Job) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	SQL := "INSERT INTO public.job(title, detail, requirement, posted_by) VALUES ($1, $2, $3, $4)"
	_, err = tx.ExecContext(ctx, SQL, job.Title, job.Detail, job.Requirement, job.PostedBy)
	helper.PanicIfError(err)

}
