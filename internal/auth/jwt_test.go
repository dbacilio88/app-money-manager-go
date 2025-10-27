package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword123"
	
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)
}

func TestCheckPassword(t *testing.T) {
	password := "testPassword123"
	
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	
	// Password correcto
	err = CheckPassword(hash, password)
	assert.NoError(t, err)
	
	// Password incorrecto
	err = CheckPassword(hash, "wrongPassword")
	assert.Error(t, err)
}

func TestGenerateJWT(t *testing.T) {
	userID := "user123"
	email := "test@example.com"
	rol := "user"
	secret := "test-secret-key"
	
	token, err := GenerateJWT(userID, email, rol, secret, 15*time.Minute)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateJWT(t *testing.T) {
	userID := "user123"
	email := "test@example.com"
	rol := "user"
	secret := "test-secret-key"
	
	// Generar token válido
	token, err := GenerateJWT(userID, email, rol, secret, 15*time.Minute)
	assert.NoError(t, err)
	
	// Validar token
	claims, err := ValidateJWT(token, secret)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, rol, claims.Rol)
}

func TestValidateJWT_Expired(t *testing.T) {
	userID := "user123"
	email := "test@example.com"
	rol := "user"
	secret := "test-secret-key"
	
	// Generar token expirado
	token, err := GenerateJWT(userID, email, rol, secret, -1*time.Hour)
	assert.NoError(t, err)
	
	// Validar token expirado
	_, err = ValidateJWT(token, secret)
	assert.Error(t, err)
	assert.ErrorIs(t, err, jwt.ErrTokenExpired)
}

func TestValidateJWT_InvalidSecret(t *testing.T) {
	userID := "user123"
	email := "test@example.com"
	rol := "user"
	secret := "test-secret-key"
	wrongSecret := "wrong-secret-key"
	
	// Generar token
	token, err := GenerateJWT(userID, email, rol, secret, 15*time.Minute)
	assert.NoError(t, err)
	
	// Validar con secreto incorrecto
	_, err = ValidateJWT(token, wrongSecret)
	assert.Error(t, err)
}
