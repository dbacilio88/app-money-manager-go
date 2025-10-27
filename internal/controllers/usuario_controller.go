package controllers

import (
	"context"
	"net/http"

	"control-financiero/internal/middleware"
	"control-financiero/internal/models"
	"control-financiero/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioController struct {
	usuarioService *services.UsuarioService
}

func NewUsuarioController(db *mongo.Database) *UsuarioController {
	return &UsuarioController{
		usuarioService: services.NewUsuarioService(db),
	}
}

func (c *UsuarioController) GetAll(ctx *gin.Context) {
	usuarios, err := c.usuarioService.GetAll(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usuarios)
}

func (c *UsuarioController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario, err := c.usuarioService.GetByID(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) GetProfile(ctx *gin.Context) {
	userID, _ := middleware.GetUserID(ctx)

	usuario, err := c.usuarioService.GetByID(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) UpdateProfile(ctx *gin.Context) {
	userID, _ := middleware.GetUserID(ctx)

	var usuario models.Usuario
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario.ID = userID

	if err := c.usuarioService.Update(context.Background(), &usuario); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Perfil actualizado correctamente"})
}

func (c *UsuarioController) Approve(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.usuarioService.Approve(context.Background(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Usuario aprobado correctamente"})
}

func (c *UsuarioController) Activate(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.usuarioService.Activate(context.Background(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Usuario activado correctamente"})
}

func (c *UsuarioController) Deactivate(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.usuarioService.Deactivate(context.Background(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Usuario desactivado correctamente"})
}

func (c *UsuarioController) ChangeRole(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req struct {
		Rol string `json:"rol" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.usuarioService.ChangeRole(context.Background(), id, req.Rol); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Rol actualizado correctamente"})
}

func (c *UsuarioController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.usuarioService.Delete(context.Background(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Usuario eliminado correctamente"})
}
