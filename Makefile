.PHONY: help build run test clean docker-build docker-up docker-down docker-logs

help: ## Mostrar esta ayuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Compilar la aplicación
	@echo "🔨 Compilando aplicación..."
	@go build -o bin/server cmd/server/main.go
	@echo "✅ Compilación exitosa"

run: ## Ejecutar la aplicación localmente
	@echo "🚀 Iniciando aplicación..."
	@go run cmd/server/main.go

test: ## Ejecutar pruebas
	@echo "🧪 Ejecutando pruebas..."
	@go test -v ./...

clean: ## Limpiar archivos generados
	@echo "🧹 Limpiando..."
	@rm -rf bin/
	@echo "✅ Limpieza completa"

docker-build: ## Construir imagen Docker
	@echo "🐳 Construyendo imagen Docker..."
	@docker-compose build
	@echo "✅ Imagen construida"

docker-up: ## Levantar servicios con Docker Compose
	@echo "🐳 Levantando servicios..."
	@docker-compose up -d
	@echo "✅ Servicios iniciados"
	@echo "📊 Mongo Express: http://localhost:8081"
	@echo "🌐 API: http://localhost:8080"

docker-down: ## Detener servicios Docker
	@echo "🛑 Deteniendo servicios..."
	@docker-compose down
	@echo "✅ Servicios detenidos"

docker-logs: ## Ver logs de los servicios
	@docker-compose logs -f app

docker-restart: ## Reiniciar servicios
	@make docker-down
	@make docker-up

install: ## Instalar dependencias
	@echo "📦 Instalando dependencias..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencias instaladas"

dev: ## Modo desarrollo con hot reload (requiere air)
	@echo "🔥 Iniciando en modo desarrollo..."
	@air

fmt: ## Formatear código
	@echo "✨ Formateando código..."
	@go fmt ./...
	@echo "✅ Código formateado"

lint: ## Ejecutar linter
	@echo "🔍 Analizando código..."
	@golangci-lint run
	@echo "✅ Análisis completo"
