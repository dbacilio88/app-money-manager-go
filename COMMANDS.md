# 🚀 Guía Rápida de Comandos - Control Financiero

## Desarrollo Local (Sin Docker)

### Prerequisitos
- Go 1.25.3 instalado
- MongoDB 7.0 corriendo en localhost:27017

### Comandos Básicos

```bash
# 1. Instalar dependencias
go mod download

# 2. Crear archivo .env
cp .env.example .env
# Editar .env con tus configuraciones

# 3. Ejecutar aplicación
go run cmd/server/main.go

# O usar el Makefile
make run
```

---

## Desarrollo con Docker (Recomendado)

### Inicio Rápido

```bash
# Levantar todos los servicios (app + mongo + mongo-express)
docker-compose up -d

# Ver logs en tiempo real
docker-compose logs -f app

# Acceder a:
# - App: http://localhost:8080
# - Mongo Express: http://localhost:8081 (admin/admin123)
```

### Con Makefile

```bash
# Ver todos los comandos disponibles
make help

# Construir imagen Docker
make docker-build

# Levantar servicios
make docker-up

# Ver logs
make docker-logs

# Detener servicios
make docker-down

# Reiniciar servicios
make docker-restart
```

---

## Testing

```bash
# Ejecutar todos los tests
go test ./...

# Tests con verbosidad
go test -v ./...

# Tests con cobertura
go test -v -race -coverprofile=coverage.out ./...

# Ver reporte HTML de cobertura
go tool cover -html=coverage.out

# O usar Makefile
make test
```

---

## Linting y Formateo

```bash
# Formatear código
go fmt ./...

# Verificar con go vet
go vet ./...

# Instalar golangci-lint (si no está instalado)
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2

# Ejecutar linter completo
golangci-lint run

# Fix automático (cuando sea posible)
golangci-lint run --fix
```

---

## Compilación

```bash
# Compilar para tu sistema
go build -o bin/server cmd/server/main.go

# Compilar para Linux (desde Windows/Mac)
GOOS=linux GOARCH=amd64 go build -o bin/server cmd/server/main.go

# Compilar para Windows (desde Linux/Mac)
GOOS=windows GOARCH=amd64 go build -o bin/server.exe cmd/server/main.go

# O usar Makefile
make build
```

---

## Base de Datos

### Mongo Shell

```bash
# Conectar a MongoDB (con Docker)
docker exec -it control-financiero-mongo mongosh

# Dentro de mongosh:
use control_financiero

# Ver usuarios
db.usuarios.find().pretty()

# Ver categorías
db.categorias.find().pretty()

# Ver transacciones
db.transacciones.find().pretty()
```

### Mongo Express Web UI

1. Abrir http://localhost:8081
2. Usuario: admin
3. Password: admin123

---

## Endpoints de API

### Autenticación

```bash
# Registrar usuario
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@controlfinanciero.com",
    "password": "admin123"
  }'

# Guardar el access_token de la respuesta
export TOKEN="tu_access_token_aqui"
```

### Perfil

```bash
# Obtener perfil
curl http://localhost:8080/api/v1/perfil \
  -H "Authorization: Bearer $TOKEN"
```

### Categorías

```bash
# Listar categorías
curl http://localhost:8080/api/v1/categorias \
  -H "Authorization: Bearer $TOKEN"

# Crear categoría
curl -X POST http://localhost:8080/api/v1/categorias \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Ahorros",
    "tipo": "ingreso",
    "color": "#4CAF50"
  }'
```

### Transacciones

```bash
# Listar transacciones
curl http://localhost:8080/api/v1/transacciones \
  -H "Authorization: Bearer $TOKEN"

# Crear transacción
curl -X POST http://localhost:8080/api/v1/transacciones \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "tipo": "ingreso",
    "categoriaId": "67890abcdef1234567890abc",
    "monto": 1000.00,
    "moneda": "USD",
    "fecha": "2025-10-27T10:00:00Z",
    "descripcion": "Pago freelance"
  }'
```

### Reportes

```bash
# Balance actual
curl http://localhost:8080/api/v1/reportes/balance \
  -H "Authorization: Bearer $TOKEN"

# Estadísticas del mes
curl "http://localhost:8080/api/v1/reportes/estadisticas?mes=10&anio=2025" \
  -H "Authorization: Bearer $TOKEN"
```

---

## Git

```bash
# Ver estado
git status

# Agregar cambios
git add .

# Commit
git commit -m "feat: descripción del cambio"

# Push
git push origin main

# Crear rama para feature
git checkout -b feature/nueva-funcionalidad

# Merge de rama
git checkout main
git merge feature/nueva-funcionalidad
```

---

## Docker - Comandos Adicionales

```bash
# Ver contenedores corriendo
docker ps

# Ver todos los contenedores
docker ps -a

# Ver logs de un contenedor específico
docker logs -f control-financiero-app

# Entrar al contenedor de la app
docker exec -it control-financiero-app sh

# Entrar al contenedor de MongoDB
docker exec -it control-financiero-mongo mongosh

# Eliminar todo (contenedores, volúmenes, imágenes)
docker-compose down -v --rmi all

# Reconstruir sin caché
docker-compose build --no-cache
```

---

## Solución de Problemas

### MongoDB no conecta

```bash
# Verificar que MongoDB está corriendo
docker ps | grep mongo

# Ver logs de MongoDB
docker logs control-financiero-mongo

# Reiniciar MongoDB
docker restart control-financiero-mongo
```

### La app no inicia

```bash
# Ver logs de la aplicación
docker logs control-financiero-app

# Verificar variables de entorno
docker exec control-financiero-app env

# Reiniciar la aplicación
docker restart control-financiero-app
```

### Tests fallan

```bash
# Limpiar caché de Go
go clean -testcache

# Ejecutar un test específico
go test -v ./internal/auth -run TestHashPassword

# Ver más detalles
go test -v -race ./...
```

### Problemas con dependencias

```bash
# Limpiar módulos
go clean -modcache

# Re-descargar dependencias
go mod download

# Verificar integridad
go mod verify

# Actualizar go.sum
go mod tidy
```

---

## Usuarios por Defecto

### Administrador
- Email: admin@controlfinanciero.com
- Password: admin123

### Usuario Normal
- Email: user@example.com
- Password: user123

---

## Variables de Entorno Importantes

```bash
# MongoDB
MONGO_URI=mongodb://localhost:27017  # o mongodb://mongo:27017 en Docker
MONGO_DB=control_financiero

# JWT
JWT_SECRET=tu-secreto-super-seguro
JWT_EXPIRATION=15m
REFRESH_EXPIRATION=168h

# App
PORT=8080
ENV=development  # o production

# Google OAuth (opcional)
GOOGLE_CLIENT_ID=tu-client-id
GOOGLE_CLIENT_SECRET=tu-client-secret
GOOGLE_REDIRECT_URL=http://localhost:8080/api/v1/auth/google/callback
```

---

## Documentación Completa

- **API**: `docs/API.md`
- **Especificación**: `specs/spec.md`
- **Modelo de Datos**: `specs/data-model.md`
- **Plan de Proyecto**: `specs/plan.md`
- **README**: `README.md`
- **Quickstart**: `QUICKSTART.md`
- **Changelog**: `CHANGELOG.md`

---

## Scripts Útiles

### Backup de MongoDB

```bash
# Crear backup
docker exec control-financiero-mongo mongodump --out /backup

# Restaurar backup
docker exec control-financiero-mongo mongorestore /backup
```

### Limpiar Logs

```bash
# Limpiar logs de Docker
docker system prune -a

# Ver uso de disco
docker system df
```

---

**¡Todo listo para desarrollar!** 🎉

Para más información, consulta `README.md` o `QUICKSTART.md`
