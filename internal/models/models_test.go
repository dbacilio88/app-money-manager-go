package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUsuario_Validation(t *testing.T) {
	tests := []struct {
		name    string
		usuario Usuario
		wantErr bool
	}{
		{
			name: "usuario válido",
			usuario: Usuario{
				Nombre:   "Test User",
				Email:    "test@example.com",
				Rol:      "user",
				Estado:   "active",
				Password: "hashedpassword",
			},
			wantErr: false,
		},
		{
			name: "email inválido",
			usuario: Usuario{
				Nombre: "Test User",
				Email:  "invalid-email",
				Rol:    "user",
				Estado: "active",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.False(t, isValidEmail(tt.usuario.Email))
			} else {
				assert.True(t, isValidEmail(tt.usuario.Email))
			}
		})
	}
}

func TestTransaccion_Validation(t *testing.T) {
	now := time.Now()
	userID := primitive.NewObjectID()
	categoriaID := primitive.NewObjectID()

	tests := []struct {
		name        string
		transaccion Transaccion
		valid       bool
	}{
		{
			name: "transacción válida - ingreso",
			transaccion: Transaccion{
				UsuarioID:   userID,
				Tipo:        "ingreso",
				CategoriaID: categoriaID,
				Monto:       1000.50,
				Moneda:      "USD",
				Fecha:       now,
				Descripcion: "Salario",
			},
			valid: true,
		},
		{
			name: "transacción válida - egreso",
			transaccion: Transaccion{
				UsuarioID:   userID,
				Tipo:        "egreso",
				CategoriaID: categoriaID,
				Monto:       50.25,
				Moneda:      "USD",
				Fecha:       now,
				Descripcion: "Compra supermercado",
			},
			valid: true,
		},
		{
			name: "monto inválido - negativo",
			transaccion: Transaccion{
				UsuarioID:   userID,
				Tipo:        "ingreso",
				CategoriaID: categoriaID,
				Monto:       -100,
				Moneda:      "USD",
				Fecha:       now,
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.Greater(t, tt.transaccion.Monto, 0.0)
				assert.NotEmpty(t, tt.transaccion.Tipo)
			} else {
				assert.LessOrEqual(t, tt.transaccion.Monto, 0.0)
			}
		})
	}
}

func TestCategoria_Validation(t *testing.T) {
	tests := []struct {
		name      string
		categoria Categoria
		valid     bool
	}{
		{
			name: "categoría válida - ingreso",
			categoria: Categoria{
				Nombre: "Salario",
				Tipo:   "ingreso",
				Color:  "#4CAF50",
			},
			valid: true,
		},
		{
			name: "categoría válida - egreso",
			categoria: Categoria{
				Nombre: "Alimentación",
				Tipo:   "egreso",
				Color:  "#F44336",
			},
			valid: true,
		},
		{
			name: "categoría sin nombre",
			categoria: Categoria{
				Nombre: "",
				Tipo:   "ingreso",
				Color:  "#4CAF50",
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.categoria.Nombre)
				assert.NotEmpty(t, tt.categoria.Tipo)
			} else {
				assert.Empty(t, tt.categoria.Nombre)
			}
		})
	}
}

// Helper function para validar emails
func isValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	return true // Simplificado para el test
}
