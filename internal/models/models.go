package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre       string             `bson:"nombre" json:"nombre" binding:"required"`
	Email        string             `bson:"email" json:"email" binding:"required,email"`
	PasswordHash string             `bson:"passwordHash,omitempty" json:"-"`
	Foto         string             `bson:"foto,omitempty" json:"foto"`
	GoogleID     string             `bson:"googleId,omitempty" json:"googleId"`
	Rol          string             `bson:"rol" json:"rol"`
	Estado       string             `bson:"estado" json:"estado"` // pending, active, suspended
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Categoria struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Nombre    string              `bson:"nombre" json:"nombre" binding:"required"`
	Tipo      string              `bson:"tipo" json:"tipo" binding:"required"` // ingreso, egreso, ambos
	Color     string              `bson:"color" json:"color"`
	UsuarioID *primitive.ObjectID `bson:"usuarioId,omitempty" json:"usuarioId"` // null para categorías globales
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
}

type Transaccion struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID   primitive.ObjectID `bson:"usuarioId" json:"usuarioId" binding:"required"`
	Tipo        string             `bson:"tipo" json:"tipo" binding:"required"` // ingreso, egreso, prestamo, alquiler, otro
	CategoriaID primitive.ObjectID `bson:"categoriaId" json:"categoriaId" binding:"required"`
	Monto       float64            `bson:"monto" json:"monto" binding:"required,gt=0"`
	Moneda      string             `bson:"moneda" json:"moneda"`
	Fecha       time.Time          `bson:"fecha" json:"fecha" binding:"required"`
	Descripcion string             `bson:"descripcion" json:"descripcion"`
	Cuenta      *Cuenta            `bson:"cuenta,omitempty" json:"cuenta"`
	MetodoPago  string             `bson:"metodoPago,omitempty" json:"metodoPago"`
	Tags        []string           `bson:"tags,omitempty" json:"tags"`
	Referencia  string             `bson:"referencia,omitempty" json:"referencia"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Cuenta struct {
	ID     primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Nombre string             `bson:"nombre" json:"nombre"`
}

type RefreshToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID primitive.ObjectID `bson:"usuarioId" json:"usuarioId"`
	Token     string             `bson:"token" json:"token"`
	ExpiresAt time.Time          `bson:"expiresAt" json:"expiresAt"`
	Revoked   bool               `bson:"revoked" json:"revoked"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type Rol struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre   string             `bson:"nombre" json:"nombre"`
	Permisos []string           `bson:"permisos" json:"permisos"`
}

type AuditLog struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	UsuarioID *primitive.ObjectID `bson:"usuarioId,omitempty" json:"usuarioId"`
	Accion    string              `bson:"accion" json:"accion"`
	Detalle   interface{}         `bson:"detalle,omitempty" json:"detalle"`
	IP        string              `bson:"ip,omitempty" json:"ip"`
	UserAgent string              `bson:"userAgent,omitempty" json:"userAgent"`
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
}

// DTOs para requests/responses
type RegisterRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Usuario      *Usuario `json:"usuario"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=6"`
}

type ReporteRequest struct {
	Year  int `form:"year" binding:"required"`
	Month int `form:"month" binding:"required,min=1,max=12"`
}

type EstadisticasResponse struct {
	TotalIngresos float64            `json:"totalIngresos"`
	TotalEgresos  float64            `json:"totalEgresos"`
	Balance       float64            `json:"balance"`
	PorCategoria  map[string]float64 `json:"porCategoria"`
}
