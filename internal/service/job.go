package service

import (
	"web3-recruitment-admin/internal/model"
	"web3-recruitment-admin/internal/repository"
)

type JobService struct {
	repo *repository.JobRepository
}

func NewJobService(repo *repository.JobRepository) *JobService {
	return &JobService{repo: repo}
}

func (s *JobService) GetAllWithStats() ([]model.JobWithStats, error) {
	return s.repo.FindAll()
}

func (s *JobService) GetByID(id int) (*model.Job, error) {
	return s.repo.FindByID(id)
}
