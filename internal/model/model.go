package model

import "time"

type JobView struct {
	ID        int       `json:"id"`
	JobID     int       `json:"jobId"`
	IPAddress string    `json:"ipAddress,omitempty"`
	UserAgent string    `json:"userAgent,omitempty"`
	ViewedAt  time.Time `json:"viewedAt"`
}

type TrackViewRequest struct {
	JobID int `json:"jobId" binding:"required"`
}

type JobViewStats struct {
	JobID        int    `json:"jobId"`
	JobTitle     string `json:"jobTitle"`
	Company      string `json:"company"`
	TotalViews   int    `json:"totalViews"`
	UniqueIPs    int    `json:"uniqueIPs"`
	ViewsToday   int    `json:"viewsToday"`
	ViewsThisWeek int   `json:"viewsThisWeek"`
}

type Application struct {
	ID              int       `json:"id"`
	JobID          int       `json:"jobId"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	ResumeText     string    `json:"resumeText,omitempty"`
	ResumeFilename string    `json:"resumeFilename,omitempty"`
	CoverLetter   string    `json:"coverLetter,omitempty"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
}

type Job struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type JobWithStats struct {
	Job
	ViewCount       int `json:"viewCount"`
	ApplicationCount int `json:"applicationCount"`
}
