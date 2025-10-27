package routes

import (
	"net/http"

	"control-financiero/internal/auth"
	"control-financiero/internal/config"
	"control-financiero/internal/controllers"
	"control-financiero/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(router *gin.Engine, db *mongo.Client, cfg *config.Config) {
	// Inicializar Google OAuth
	auth.InitGoogleOAuth()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Controladores
	database := db.Database(cfg.MongoDB)
	authController := controllers.NewAuthController(database, cfg)
	usuarioController := controllers.NewUsuarioController(database)
	categoriaController := controllers.NewCategoriaController(database)
	transaccionController := controllers.NewTransaccionController(database)

	// Rutas públicas
	api := router.Group("/api/v1")
	{
		// Autenticación
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.POST("/refresh", authController.RefreshToken)
			auth.GET("/google", authController.GoogleAuthURL)
			auth.GET("/google/callback", authController.GoogleCallback)
		}

		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"service": "control-financiero",
			})
		})
	}

	// Rutas protegidas
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// Perfil
		protected.GET("/perfil", usuarioController.GetProfile)
		protected.PUT("/perfil", usuarioController.UpdateProfile)
		protected.POST("/cambiar-password", authController.ChangePassword)

		// Categorías
		categorias := protected.Group("/categorias")
		{
			categorias.POST("", categoriaController.Create)
			categorias.GET("", categoriaController.GetAll)
			categorias.GET("/:id", categoriaController.GetByID)
			categorias.PUT("/:id", categoriaController.Update)
			categorias.DELETE("/:id", categoriaController.Delete)
		}

		// Transacciones
		transacciones := protected.Group("/transacciones")
		{
			transacciones.POST("", transaccionController.Create)
			transacciones.GET("", transaccionController.GetAll)
			transacciones.GET("/:id", transaccionController.GetByID)
			transacciones.PUT("/:id", transaccionController.Update)
			transacciones.DELETE("/:id", transaccionController.Delete)
		}

		// Reportes
		reportes := protected.Group("/reportes")
		{
			reportes.GET("/estadisticas", transaccionController.GetEstadisticas)
		}

		// Rutas de administrador
		admin := protected.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			// Usuarios
			admin.GET("/usuarios", usuarioController.GetAll)
			admin.GET("/usuarios/:id", usuarioController.GetByID)
			admin.PATCH("/usuarios/:id/aprobar", usuarioController.Approve)
			admin.PATCH("/usuarios/:id/activar", usuarioController.Activate)
			admin.PATCH("/usuarios/:id/desactivar", usuarioController.Deactivate)
			admin.PATCH("/usuarios/:id/rol", usuarioController.ChangeRole)
			admin.DELETE("/usuarios/:id", usuarioController.Delete)
		}
	}

	// Servir archivos estáticos
	router.Static("/assets", "./assets")
	router.Static("/web/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")

	// Ruta principal
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Control Financiero",
		})
	})
}
