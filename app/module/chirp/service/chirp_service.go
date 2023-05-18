package service

import (
	"chirp-api/app/module/chirp/repository"
	"chirp-api/app/module/chirp/request"
	"chirp-api/internal/ent"

	"github.com/google/uuid"
)

type ChirpService struct {
	Repo *repository.ChirpRepository
}

type IChirpService interface {
	GetChirps() ([]*ent.Post, error)
	GetChiprsByID(id uuid.UUID) (*ent.Post, error)
	CreateChirp(request request.PostRequest) (*ent.Post, error)
	UpdateChirp(id uuid.UUID, request request.PostRequest) (*ent.Post, error)
	DeleteChirp(id uuid.UUID) error
}

func NewChirpService(repo *repository.ChirpRepository) *ChirpService {
	return &ChirpService{
		Repo: repo,
	}
}

func (s *ChirpService) GetChirps() ([]*ent.Post, error) {
	return s.Repo.GetChirps()
}

func (s *ChirpService) GetChiprsByID(id uuid.UUID) (*ent.Post, error) {
	return s.Repo.GetChirpByID(id)
}

func (s *ChirpService) CreateChirp(request request.PostRequest) (*ent.Post, error) {
	return s.Repo.CreatePost(request)
}

func (s *ChirpService) UpdateChirp(id uuid.UUID, request request.PostRequest) (*ent.Post, error) {
	return s.Repo.UpdatePost(id, request)
}

func (s *ChirpService) DeleteChirp(id uuid.UUID) error {
	return s.Repo.DeletePost(id)
}
