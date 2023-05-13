package service

import (
	"chirp-api/app/module/chirp/repository"
	"chirp-api/internal/ent"
)

type ChirpService struct {
	Repo *repository.ChirpRepository
}

type IChirpService interface {
	GetChirps() ([]*ent.Post, error)
}

func NewChirpService(repo *repository.ChirpRepository) *ChirpService {
	return &ChirpService{
		Repo: repo,
	}
}

func (s *ChirpService) GetChirps() ([]*ent.Post, error) {
	return s.Repo.GetChirps()
}
