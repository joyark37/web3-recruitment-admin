package repository

import (
	"database/sql"

	"web3-recruitment-admin/internal/model"
)

type JobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (r *JobRepository) FindAll() ([]model.JobWithStats, error) {
	rows, err := r.db.Query(`
		SELECT 
			j.id, j.title, j.company, j.location, j.job_type, j.category, j.status, j.created_at,
			COALESCE(v.view_count, 0) as view_count,
			COALESCE(a.app_count, 0) as app_count
		FROM jobs j
		LEFT JOIN (SELECT job_id, COUNT(*) as view_count FROM job_views GROUP BY job_id) v ON j.id = v.job_id
		LEFT JOIN (SELECT job_id, COUNT(*) as app_count FROM applications GROUP BY job_id) a ON j.id = a.job_id
		ORDER BY j.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []model.JobWithStats
	for rows.Next() {
		var job model.JobWithStats
		err := rows.Scan(
			&job.ID, &job.Title, &job.Company, &job.Location, &job.Type, 
			&job.Category, &job.Status, &job.CreatedAt, &job.ViewCount, &job.ApplicationCount,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (r *JobRepository) FindByID(id int) (*model.Job, error) {
	var job model.Job
	err := r.db.QueryRow(`
		SELECT id, title, company, location, job_type, category, status, created_at
		FROM jobs WHERE id = $1`, id).Scan(
		&job.ID, &job.Title, &job.Company, &job.Location, &job.Type, 
		&job.Category, &job.Status, &job.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &job, nil
}
