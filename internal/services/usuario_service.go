package services

import (
	"context"
	"errors"

	"control-financiero/internal/models"
	"control-financiero/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioService struct {
	userRepo *repositories.UsuarioRepository
}

func NewUsuarioService(db *mongo.Database) *UsuarioService {
	return &UsuarioService{
		userRepo: repositories.NewUsuarioRepository(db),
	}
}

func (s *UsuarioService) GetAll(ctx context.Context) ([]*models.Usuario, error) {
	usuarios, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// Ocultar password hash
	for _, u := range usuarios {
		u.PasswordHash = ""
	}

	return usuarios, nil
}

func (s *UsuarioService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Usuario, error) {
	usuario, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	usuario.PasswordHash = ""
	return usuario, nil
}

func (s *UsuarioService) Update(ctx context.Context, usuario *models.Usuario) error {
	existing, err := s.userRepo.FindByID(ctx, usuario.ID)
	if err != nil {
		return err
	}

	// Solo permitir actualizar ciertos campos
	existing.Nombre = usuario.Nombre
	existing.Foto = usuario.Foto

	return s.userRepo.Update(ctx, existing)
}

func (s *UsuarioService) Approve(ctx context.Context, id primitive.ObjectID) error {
	return s.userRepo.UpdateEstado(ctx, id, "active")
}

func (s *UsuarioService) Activate(ctx context.Context, id primitive.ObjectID) error {
	return s.userRepo.UpdateEstado(ctx, id, "active")
}

func (s *UsuarioService) Deactivate(ctx context.Context, id primitive.ObjectID) error {
	return s.userRepo.UpdateEstado(ctx, id, "suspended")
}

func (s *UsuarioService) ChangeRole(ctx context.Context, id primitive.ObjectID, rol string) error {
	if rol != "admin" && rol != "user" {
		return errors.New("rol inválido")
	}

	usuario, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	usuario.Rol = rol
	return s.userRepo.Update(ctx, usuario)
}

func (s *UsuarioService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.userRepo.Delete(ctx, id)
}
