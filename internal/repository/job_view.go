package repository

import (
	"database/sql"

	"web3-recruitment-admin/internal/model"
)

type JobViewRepository struct {
	db *sql.DB
}

func NewJobViewRepository(db *sql.DB) *JobViewRepository {
	return &JobViewRepository{db: db}
}

func (r *JobViewRepository) Create(view *model.JobView) error {
	return r.db.QueryRow(`
		INSERT INTO job_views (job_id, ip_address, user_agent)
		VALUES ($1, $2, $3)
		RETURNING id, viewed_at`,
		view.JobID, view.IPAddress, view.UserAgent,
	).Scan(&view.ID, &view.ViewedAt)
}

func (r *JobViewRepository) GetStatsByJobID(jobID int) (*model.JobViewStats, error) {
	var stats model.JobViewStats
	err := r.db.QueryRow(`
		SELECT 
			j.id,
			j.title,
			j.company,
			COUNT(*) as total_views,
			COUNT(DISTINCT ip_address) as unique_ips,
			COUNT(*) FILTER (WHERE viewed_at >= CURRENT_DATE) as views_today,
			COUNT(*) FILTER (WHERE viewed_at >= CURRENT_DATE - INTERVAL '7 days') as views_this_week
		FROM job_views jv
		JOIN jobs j ON j.id = jv.job_id
		WHERE jv.job_id = $1
		GROUP BY j.id, j.title, j.company`,
		jobID,
	).Scan(&stats.JobID, &stats.JobTitle, &stats.Company, &stats.TotalViews, &stats.UniqueIPs, &stats.ViewsToday, &stats.ViewsThisWeek)
	
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

func (r *JobViewRepository) GetAllStats() ([]model.JobViewStats, error) {
	rows, err := r.db.Query(`
		SELECT 
			j.id,
			j.title,
			j.company,
			COUNT(*) as total_views,
			COUNT(DISTINCT ip_address) as unique_ips,
			COUNT(*) FILTER (WHERE viewed_at >= CURRENT_DATE) as views_today,
			COUNT(*) FILTER (WHERE viewed_at >= CURRENT_DATE - INTERVAL '7 days') as views_this_week
		FROM job_views jv
		JOIN jobs j ON j.id = jv.job_id
		GROUP BY j.id, j.title, j.company
		ORDER BY total_views DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []model.JobViewStats
	for rows.Next() {
		var s model.JobViewStats
		if err := rows.Scan(&s.JobID, &s.JobTitle, &s.Company, &s.TotalViews, &s.UniqueIPs, &s.ViewsToday, &s.ViewsThisWeek); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}
	return stats, nil
}

func (r *JobViewRepository) GetDailyViews(days int) ([]model.DailyView, error) {
	rows, err := r.db.Query(`
		SELECT 
			TO_CHAR(DATE(viewed_at), 'YYYY-MM-DD') as date,
			COUNT(*) as views
		FROM job_views
		WHERE viewed_at >= CURRENT_DATE - INTERVAL '7 days'
		GROUP BY DATE(viewed_at)
		ORDER BY date ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []model.DailyView
	for rows.Next() {
		var v model.DailyView
		if err := rows.Scan(&v.Date, &v.Views); err != nil {
			return nil, err
		}
		views = append(views, v)
	}
	if views == nil {
		views = []model.DailyView{}
	}
	return views, nil
}
