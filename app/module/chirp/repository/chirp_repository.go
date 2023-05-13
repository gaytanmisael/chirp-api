package repository

import (
	"chirp-api/internal/bootstrap/database"
	"chirp-api/internal/ent"
	"chirp-api/internal/ent/post"
	"context"
)

type ChirpRepository struct {
	DB *database.Database
}

type IArticleRepository interface {
	GetChirps() ([]*ent.Post, error)
	GetChirpByID(id int) (*ent.Post, error)
}

func NewChirpRepository(database *database.Database) *ChirpRepository {
	return &ChirpRepository{
		DB: database,
	}
}

func (s *ChirpRepository) GetChirps() ([]*ent.Post, error) {
	return s.DB.Ent.Post.Query().Order(ent.Asc(post.FieldID)).All(context.Background())
}
