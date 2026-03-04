package repository

import (
	"database/sql"

	"web3-recruitment-admin/internal/model"
)

type ApplicationRepository struct {
	db *sql.DB
}

func NewApplicationRepository(db *sql.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

func (r *ApplicationRepository) FindAll() ([]model.Application, error) {
	rows, err := r.db.Query(`
		SELECT id, job_id, name, email, resume, cover_letter, status, created_at
		FROM applications 
		ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []model.Application
	for rows.Next() {
		var app model.Application
		err := rows.Scan(
			&app.ID, &app.JobID, &app.Name, &app.Email, &app.ResumeText,
			&app.CoverLetter, &app.Status, &app.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	return apps, nil
}

func (r *ApplicationRepository) FindByJobID(jobID int) ([]model.Application, error) {
	rows, err := r.db.Query(`
		SELECT id, job_id, name, email, resume, cover_letter, status, created_at
		FROM applications 
		WHERE job_id = $1
		ORDER BY created_at DESC`, jobID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []model.Application
	for rows.Next() {
		var app model.Application
		err := rows.Scan(
			&app.ID, &app.JobID, &app.Name, &app.Email, &app.ResumeText,
			&app.CoverLetter, &app.Status, &app.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	return apps, nil
}

func (r *ApplicationRepository) CountByJobID(jobID int) (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM applications WHERE job_id = $1", jobID).Scan(&count)
	return count, err
}
