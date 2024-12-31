package services

import "BecomeOverMan/internal/repositories"

type BaseService struct {
	repo *repositories.BaseRepository
}

func NewBaseService(repo *repositories.BaseRepository) *BaseService {
	return &BaseService{repo: repo}
}

func (s *BaseService) CheckConnection() error {
	return s.repo.CheckConnection()
}
