package services

import (
	"context"
	"time"

	"control-financiero/internal/models"
	"control-financiero/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransaccionService struct {
	transaccionRepo *repositories.TransaccionRepository
}

func NewTransaccionService(db *mongo.Database) *TransaccionService {
	return &TransaccionService{
		transaccionRepo: repositories.NewTransaccionRepository(db),
	}
}

func (s *TransaccionService) Create(ctx context.Context, transaccion *models.Transaccion) error {
	return s.transaccionRepo.Create(ctx, transaccion)
}

func (s *TransaccionService) GetByUsuario(ctx context.Context, usuarioID primitive.ObjectID) ([]*models.Transaccion, error) {
	return s.transaccionRepo.FindByUsuario(ctx, usuarioID)
}

func (s *TransaccionService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Transaccion, error) {
	return s.transaccionRepo.FindByID(ctx, id)
}

func (s *TransaccionService) Update(ctx context.Context, transaccion *models.Transaccion) error {
	return s.transaccionRepo.Update(ctx, transaccion)
}

func (s *TransaccionService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.transaccionRepo.Delete(ctx, id)
}

func (s *TransaccionService) GetEstadisticas(ctx context.Context, usuarioID primitive.ObjectID, year, month int) (*models.EstadisticasResponse, error) {
	// Calcular rango de fechas
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0).Add(-time.Second)

	transacciones, err := s.transaccionRepo.FindByUsuarioAndRange(ctx, usuarioID, start, end)
	if err != nil {
		return nil, err
	}

	// Calcular estadísticas
	var totalIngresos, totalEgresos float64
	porCategoria := make(map[string]float64)

	for _, t := range transacciones {
		if t.Tipo == "ingreso" {
			totalIngresos += t.Monto
		} else if t.Tipo == "egreso" {
			totalEgresos += t.Monto
		}

		// Agregar por categoría (esto requeriría hacer lookup de la categoría)
		// Por simplicidad, usamos el ID de la categoría como string
		key := t.CategoriaID.Hex()
		porCategoria[key] += t.Monto
	}

	return &models.EstadisticasResponse{
		TotalIngresos: totalIngresos,
		TotalEgresos:  totalEgresos,
		Balance:       totalIngresos - totalEgresos,
		PorCategoria:  porCategoria,
	}, nil
}
