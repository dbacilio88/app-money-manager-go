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

type TransaccionController struct {
	transaccionService *services.TransaccionService
}

func NewTransaccionController(db *mongo.Database) *TransaccionController {
	return &TransaccionController{
		transaccionService: services.NewTransaccionService(db),
	}
}

func (c *TransaccionController) Create(ctx *gin.Context) {
	userID, _ := middleware.GetUserID(ctx)

	var transaccion models.Transaccion
	if err := ctx.ShouldBindJSON(&transaccion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaccion.UsuarioID = userID

	if err := c.transaccionService.Create(context.Background(), &transaccion); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaccion)
}

func (c *TransaccionController) GetAll(ctx *gin.Context) {
	userID, _ := middleware.GetUserID(ctx)

	transacciones, err := c.transaccionService.GetByUsuario(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transacciones)
}

func (c *TransaccionController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	transaccion, err := c.transaccionService.GetByID(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transacción no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, transaccion)
}

func (c *TransaccionController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var transaccion models.Transaccion
	if err := ctx.ShouldBindJSON(&transaccion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaccion.ID = id

	if err := c.transaccionService.Update(context.Background(), &transaccion); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaccion)
}

func (c *TransaccionController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.transaccionService.Delete(context.Background(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Transacción eliminada correctamente"})
}

func (c *TransaccionController) GetEstadisticas(ctx *gin.Context) {
	userID, _ := middleware.GetUserID(ctx)

	var req models.ReporteRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	estadisticas, err := c.transaccionService.GetEstadisticas(context.Background(), userID, req.Year, req.Month)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, estadisticas)
}
