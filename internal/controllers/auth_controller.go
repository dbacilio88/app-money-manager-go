package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"control-financiero/internal/auth"
	"control-financiero/internal/config"
	"control-financiero/internal/models"
	"control-financiero/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	authService *services.AuthService
	cfg         *config.Config
}

func NewAuthController(db *mongo.Database, cfg *config.Config) *AuthController {
	return &AuthController{
		authService: services.NewAuthService(db, cfg),
		cfg:         cfg,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req models.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario, err := c.authService.Register(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"mensaje": "Usuario registrado. Pendiente de aprobación por administrador",
		"usuario": usuario,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.authService.Login(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) GoogleAuthURL(ctx *gin.Context) {
	state := fmt.Sprintf("state-%d", time.Now().Unix())
	url := auth.GetGoogleAuthURL(state)

	ctx.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func (c *AuthController) GoogleCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.Redirect(http.StatusFound, c.cfg.AppURL+"?error=no_code")
		return
	}

	// Intercambiar código por token
	token, err := auth.ExchangeGoogleCode(code)
	if err != nil {
		ctx.Redirect(http.StatusFound, c.cfg.AppURL+"?error=exchange_failed")
		return
	}

	// Obtener información del usuario
	userInfo, err := auth.GetGoogleUserInfo(token)
	if err != nil {
		ctx.Redirect(http.StatusFound, c.cfg.AppURL+"?error=userinfo_failed")
		return
	}

	// Login o registro con Google
	response, err := c.authService.LoginWithGoogle(
		context.Background(),
		userInfo.ID,
		userInfo.Email,
		userInfo.Name,
		userInfo.Picture,
	)
	if err != nil {
		ctx.Redirect(http.StatusFound, c.cfg.AppURL+"?error="+err.Error())
		return
	}

	// Redirigir al frontend con tokens
	redirectURL := fmt.Sprintf("%s?access_token=%s&refresh_token=%s",
		c.cfg.AppURL,
		response.AccessToken,
		response.RefreshToken,
	)
	ctx.Redirect(http.StatusFound, redirectURL)
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var req models.RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := c.authService.RefreshToken(context.Background(), req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}

func (c *AuthController) ChangePassword(ctx *gin.Context) {
	userID := ctx.MustGet("userId").(primitive.ObjectID)

	var req models.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.authService.ChangePassword(context.Background(), userID, &req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Contraseña actualizada correctamente"})
}
