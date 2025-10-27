package services

import (
	"context"
	"errors"
	"time"

	"control-financiero/internal/auth"
	"control-financiero/internal/config"
	"control-financiero/internal/models"
	"control-financiero/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	userRepo         *repositories.UsuarioRepository
	refreshTokenRepo *repositories.RefreshTokenRepository
	cfg              *config.Config
}

func NewAuthService(db *mongo.Database, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo:         repositories.NewUsuarioRepository(db),
		refreshTokenRepo: repositories.NewRefreshTokenRepository(db),
		cfg:              cfg,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.Usuario, error) {
	// Verificar si el usuario ya existe
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("el correo ya está registrado")
	}

	// Hash de la contraseña
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Crear usuario
	usuario := &models.Usuario{
		Nombre:       req.Nombre,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Rol:          "user",
		Estado:       "pending",
	}

	if err := s.userRepo.Create(ctx, usuario); err != nil {
		return nil, err
	}

	return usuario, nil
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	// Buscar usuario
	usuario, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Verificar estado del usuario
	if usuario.Estado != "active" {
		return nil, errors.New("usuario no activo o pendiente de aprobación")
	}

	// Verificar contraseña
	if !auth.CheckPassword(req.Password, usuario.PasswordHash) {
		return nil, errors.New("credenciales inválidas")
	}

	// Generar tokens
	accessToken, err := auth.GenerateJWT(usuario, s.cfg.JWTSecret, s.cfg.JWTExpiration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Guardar refresh token
	refreshDuration, _ := time.ParseDuration(s.cfg.RefreshExpiration)
	rt := &models.RefreshToken{
		UsuarioID: usuario.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(refreshDuration),
		Revoked:   false,
	}

	if err := s.refreshTokenRepo.Create(ctx, rt); err != nil {
		return nil, err
	}

	// Ocultar password hash en la respuesta
	usuario.PasswordHash = ""

	return &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Usuario:      usuario,
	}, nil
}

func (s *AuthService) LoginWithGoogle(ctx context.Context, googleID, email, nombre, foto string) (*models.LoginResponse, error) {
	// Buscar usuario por Google ID
	usuario, err := s.userRepo.FindByGoogleID(ctx, googleID)
	if err != nil {
		// Si no existe, buscar por email
		usuario, err = s.userRepo.FindByEmail(ctx, email)
		if err != nil {
			// Crear nuevo usuario
			usuario = &models.Usuario{
				Nombre:   nombre,
				Email:    email,
				Foto:     foto,
				GoogleID: googleID,
				Rol:      "user",
				Estado:   "pending",
			}

			if err := s.userRepo.Create(ctx, usuario); err != nil {
				return nil, err
			}
		} else {
			// Actualizar Google ID
			usuario.GoogleID = googleID
			if err := s.userRepo.Update(ctx, usuario); err != nil {
				return nil, err
			}
		}
	}

	// Verificar estado
	if usuario.Estado != "active" {
		return nil, errors.New("usuario no activo o pendiente de aprobación")
	}

	// Generar tokens
	accessToken, err := auth.GenerateJWT(usuario, s.cfg.JWTSecret, s.cfg.JWTExpiration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Guardar refresh token
	refreshDuration, _ := time.ParseDuration(s.cfg.RefreshExpiration)
	rt := &models.RefreshToken{
		UsuarioID: usuario.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(refreshDuration),
		Revoked:   false,
	}

	if err := s.refreshTokenRepo.Create(ctx, rt); err != nil {
		return nil, err
	}

	usuario.PasswordHash = ""

	return &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Usuario:      usuario,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	// Buscar refresh token
	rt, err := s.refreshTokenRepo.FindByToken(ctx, tokenString)
	if err != nil {
		return "", errors.New("refresh token inválido")
	}

	// Verificar expiración
	if time.Now().After(rt.ExpiresAt) {
		return "", errors.New("refresh token expirado")
	}

	// Buscar usuario
	usuario, err := s.userRepo.FindByID(ctx, rt.UsuarioID)
	if err != nil {
		return "", err
	}

	// Verificar estado
	if usuario.Estado != "active" {
		return "", errors.New("usuario no activo")
	}

	// Generar nuevo access token
	accessToken, err := auth.GenerateJWT(usuario, s.cfg.JWTSecret, s.cfg.JWTExpiration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, usuarioID primitive.ObjectID, req *models.ChangePasswordRequest) error {
	usuario, err := s.userRepo.FindByID(ctx, usuarioID)
	if err != nil {
		return err
	}

	// Verificar contraseña actual
	if !auth.CheckPassword(req.CurrentPassword, usuario.PasswordHash) {
		return errors.New("contraseña actual incorrecta")
	}

	// Hash nueva contraseña
	hashedPassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	usuario.PasswordHash = hashedPassword
	return s.userRepo.Update(ctx, usuario)
}
