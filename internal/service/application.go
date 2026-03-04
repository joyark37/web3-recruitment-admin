package service

import (
	"web3-recruitment-admin/internal/model"
	"web3-recruitment-admin/internal/repository"
)

type ApplicationService struct {
	repo *repository.ApplicationRepository
}

func NewApplicationService(repo *repository.ApplicationRepository) *ApplicationService {
	return &ApplicationService{repo: repo}
}

func (s *ApplicationService) GetAll() ([]model.Application, error) {
	return s.repo.FindAll()
}

func (s *ApplicationService) GetByJobID(jobID int) ([]model.Application, error) {
	return s.repo.FindByJobID(jobID)
}

func (s *ApplicationService) CountByJobID(jobID int) (int, error) {
	return s.repo.CountByJobID(jobID)
}
