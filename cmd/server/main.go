package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"control-financiero/internal/config"
	"control-financiero/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	cfg := config.Load()

	// Conectar a MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := config.ConnectMongoDB(ctx, cfg)
	if err != nil {
		log.Fatal("❌ Error conectando a MongoDB:", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Println("Error desconectando MongoDB:", err)
		}
	}()

	log.Println("✅ Conectado a MongoDB")

	// Inicializar colecciones y datos por defecto
	if err := config.InitializeCollections(mongoClient, cfg); err != nil {
		log.Fatal("❌ Error inicializando colecciones:", err)
	}

	// Configurar Gin
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Configurar rutas
	routes.Setup(router, mongoClient, cfg)

	// Servidor con graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Canal para señales del sistema
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("🚀 Servidor iniciando en el puerto %s\n", cfg.Port)
		log.Printf("🌐 API disponible en: http://localhost:%s/api/v1\n", cfg.Port)
		log.Printf("🖥️  Aplicación web en: http://localhost:%s\n", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar servidor: %v\n", err)
		}
	}()

	// Esperar señal de terminación
	<-quit
	log.Println("🛑 Apagando servidor...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("❌ Error en shutdown:", err)
	}

	log.Println("✅ Servidor detenido correctamente")
}
