package repository

import (
	"chirp-api/app/module/chirp/request"
	"chirp-api/internal/bootstrap/database"
	"chirp-api/internal/ent"
	"chirp-api/internal/ent/post"
	"context"

	"github.com/google/uuid"
)

type ChirpRepository struct {
	DB *database.Database
}

type IArticleRepository interface {
	GetChirps() ([]*ent.Post, error)
	GetChirpByID(id uuid.UUID) (*ent.Post, error)
	CreatePost(request request.PostRequest) (*ent.Post, error)
	UpdatePost(id uuid.UUID, request request.PostRequest) (*ent.Post, error)
	DeletePost(id uuid.UUID) error
}

func NewChirpRepository(database *database.Database) *ChirpRepository {
	return &ChirpRepository{
		DB: database,
	}
}

func (s *ChirpRepository) GetChirps() ([]*ent.Post, error) {
	return s.DB.Ent.Post.Query().Order(ent.Asc(post.FieldID)).All(context.Background())
}

func (s *ChirpRepository) GetChirpByID(id uuid.UUID) (*ent.Post, error) {
	return s.DB.Ent.Post.Query().Where(post.IDEQ(id)).First(context.Background())
}

func (s *ChirpRepository) CreatePost(request request.PostRequest) (*ent.Post, error) {
	return s.DB.Ent.Post.Create().SetContent(request.Content).SetAuthorId(request.AuthorId).Save(context.Background())
}

func (s *ChirpRepository) UpdatePost(id uuid.UUID, request request.PostRequest) (*ent.Post, error) {
	return s.DB.Ent.Post.UpdateOneID(id).SetContent(request.Content).Save(context.Background())
}

func (s *ChirpRepository) DeletePost(id uuid.UUID) error {
	return s.DB.Ent.Post.DeleteOneID(id).Exec(context.Background())
}
