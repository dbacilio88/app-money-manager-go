package services

import (
	"context"

	"control-financiero/internal/models"
	"control-financiero/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoriaService struct {
	categoriaRepo *repositories.CategoriaRepository
}

func NewCategoriaService(db *mongo.Database) *CategoriaService {
	return &CategoriaService{
		categoriaRepo: repositories.NewCategoriaRepository(db),
	}
}

func (s *CategoriaService) Create(ctx context.Context, categoria *models.Categoria) error {
	return s.categoriaRepo.Create(ctx, categoria)
}

func (s *CategoriaService) GetAll(ctx context.Context, usuarioID *primitive.ObjectID) ([]*models.Categoria, error) {
	return s.categoriaRepo.FindAll(ctx, usuarioID)
}

func (s *CategoriaService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Categoria, error) {
	return s.categoriaRepo.FindByID(ctx, id)
}

func (s *CategoriaService) Update(ctx context.Context, categoria *models.Categoria) error {
	return s.categoriaRepo.Update(ctx, categoria)
}

func (s *CategoriaService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.categoriaRepo.Delete(ctx, id)
}
