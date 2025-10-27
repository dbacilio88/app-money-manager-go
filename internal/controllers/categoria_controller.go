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

type CategoriaController struct {
	categoriaService *services.CategoriaService
}

func NewCategoriaController(db *mongo.Database) *CategoriaController {
	return &CategoriaController{
		categoriaService: services.NewCategoriaService(db),
	}
}

func (c *CategoriaController) Create(ctx *gin.Context) {
	userID, _ := middleware.GetUserID(ctx)

	var categoria models.Categoria
	if err := ctx.ShouldBindJSON(&categoria); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoria.UsuarioID = &userID

	if err := c.categoriaService.Create(context.Background(), &categoria); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, categoria)
}

func (c *CategoriaController) GetAll(ctx *gin.Context) {
	userID, _ := middleware.GetUserID(ctx)

	categorias, err := c.categoriaService.GetAll(context.Background(), &userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categorias)
}

func (c *CategoriaController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	categoria, err := c.categoriaService.GetByID(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, categoria)
}

func (c *CategoriaController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var categoria models.Categoria
	if err := ctx.ShouldBindJSON(&categoria); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoria.ID = id

	if err := c.categoriaService.Update(context.Background(), &categoria); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categoria)
}

func (c *CategoriaController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.categoriaService.Delete(context.Background(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Categoría eliminada correctamente"})
}
