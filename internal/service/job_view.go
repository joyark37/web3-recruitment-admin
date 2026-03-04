package service

import (
	"web3-recruitment-admin/internal/model"
	"web3-recruitment-admin/internal/repository"
)

type JobViewService struct {
	repo *repository.JobViewRepository
}

func NewJobViewService(repo *repository.JobViewRepository) *JobViewService {
	return &JobViewService{repo: repo}
}

func (s *JobViewService) TrackView(jobID int, ipAddress, userAgent string) error {
	view := &model.JobView{
		JobID:     jobID,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
	return s.repo.Create(view)
}

func (s *JobViewService) GetStatsByJobID(jobID int) (*model.JobViewStats, error) {
	return s.repo.GetStatsByJobID(jobID)
}

func (s *JobViewService) GetAllStats() ([]model.JobViewStats, error) {
	return s.repo.GetAllStats()
}
