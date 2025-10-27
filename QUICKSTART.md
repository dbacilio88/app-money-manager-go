# 🚀 Guía de Inicio Rápido - Control Financiero Personal

Esta guía te permitirá tener el sistema funcionando en menos de 5 minutos.

**Stack**: Go 1.25.3 + MongoDB + Gin Framework

## ⚡ Opción 1: Inicio Súper Rápido (Recomendado)

### 1. Prerrequisitos Mínimos
```bash
# Verificar que tengas Docker instalado
docker --version
docker-compose --version
```

### 2. Clonar y Ejecutar
```bash
# Clonar repositorio
git clone https://github.com/tu-usuario/control-financiero.git
cd control-financiero

# Iniciar todos los servicios (incluye MongoDB automática)
docker-compose up -d

# Ver logs en tiempo real
docker-compose logs -f app
```

### 3. ¡Listo! 🎉
- **Aplicación**: http://localhost:8080
- **MongoDB Web UI**: http://localhost:8081 (usuario: admin, password: admin123)

## 🔧 Opción 2: Configuración Personalizada

### 1. Configurar Variables de Google OAuth (Opcional pero Recomendado)

```bash
# Copiar archivo de configuración
cp .env.example .env

# Editar con tus credenciales de Google Cloud Console
nano .env
```

**Obtener credenciales de Google:**
1. Ir a [Google Cloud Console](https://console.cloud.google.com/)
2. Crear proyecto → APIs & Services → Credentials
3. Crear OAuth 2.0 Client ID
4. Agregar `http://localhost:8080/api/v1/auth/google/callback` como URI autorizada
5. Copiar Client ID y Client Secret al `.env`

### 2. Ejecutar con Configuración Personalizada
```bash
docker-compose up --build -d
```

## 🧪 Opción 3: Desarrollo Local (Sin Docker)

### 1. Instalar Dependencias
```bash
# Go 1.22+
go version

# PostgreSQL
# Ubuntu/Debian: sudo apt install postgresql postgresql-contrib
# macOS: brew install postgresql
# Windows: Descargar desde postgresql.org
```

### 2. Configurar Base de Datos Local
```bash
# Crear base de datos
sudo -u postgres createdb control_financiero

# Configurar variables de entorno
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=tu_password
export DB_NAME=control_financiero
export JWT_SECRET=tu_jwt_secret_seguro
```

### 3. Ejecutar Aplicación
```bash
# Instalar dependencias
go mod download

# Ejecutar
go run cmd/server/main.go
```

## 🎯 Primer Uso del Sistema

### 1. Acceder a la Aplicación
1. Abrir navegador en http://localhost:8080
2. Hacer clic en "Continuar con Google"
3. Completar autenticación (o usar modo demo)

### 2. Configuración Inicial
1. **Establecer Balance Inicial**: Ve a tu perfil y establece tu balance actual
2. **Crear Categorías**: Alimentación, Transporte, Entretenimiento, etc.
3. **Registrar Primera Transacción**: Ingreso o gasto para probar el sistema

### 3. Explorar Funcionalidades
- **Dashboard**: Vista general con gráficas
- **Transacciones**: Historial y registro de movimientos
- **Categorías**: Organización de tus gastos e ingresos
- **Reportes**: Análisis mensual automático

## 🛠️ Comandos Útiles

### Docker
```bash
# Ver estado de servicios
docker-compose ps

# Parar servicios
docker-compose down

# Reiniciar solo la app
docker-compose restart app

# Ver logs de BD
docker-compose logs postgres

# Backup de base de datos
docker-compose exec postgres pg_dump -U postgres control_financiero > backup_$(date +%Y%m%d).sql
```

### Desarrollo
```bash
# Tests
go test ./...

# Formatear código
go fmt ./...

# Construir binario
go build -o bin/server cmd/server/main.go
```

## 🚨 Solución de Problemas Rápidos

### Error: "Puerto 8080 en uso"
```bash
# Cambiar puerto en docker-compose.yml
ports:
  - "8081:8080"  # Cambiar primer número
```

### Error: "Base de datos no conecta"
```bash
# Verificar logs
docker-compose logs postgres

# Reiniciar servicios
docker-compose down && docker-compose up -d
```

### Error: "Google OAuth no funciona"
- **Solución temporal**: La aplicación incluye un modo demo que simula el login
- **Solución completa**: Configurar credenciales reales en `.env`

### Frontend no carga estilos
```bash
# Verificar que los assets estén montados
docker-compose logs app | grep assets

# Reiniciar contenedor
docker-compose restart app
```

## 🎮 Modo Demo

Si no quieres configurar Google OAuth inmediatamente, la aplicación incluye un **modo demo** que:

- ✅ Simula un usuario autenticado
- ✅ Permite probar todas las funcionalidades
- ✅ Usa datos de ejemplo para las gráficas
- ✅ Funciona sin configuración externa

**Para activar modo demo**: Simplemente haz clic en "Continuar con Google" y el sistema simulará automáticamente el login.

## 📱 Acceso Móvil

La aplicación es completamente responsiva:

- **Desde tu móvil**: http://tu-ip-local:8080
- **Desde otra computadora en tu red**: http://tu-ip-local:8080

Para encontrar tu IP local:
```bash
# Linux/macOS
ip addr show | grep inet

# Windows
ipconfig
```

## ⚙️ Variables de Entorno Importantes

```env
# Básicas para funcionar
DB_HOST=postgres
DB_NAME=control_financiero
JWT_SECRET=cambiar-por-algo-seguro

# Para producción
GOOGLE_CLIENT_ID=tu_client_id_real
GOOGLE_CLIENT_SECRET=tu_client_secret_real
GIN_MODE=release

# Personalización
PORT=8080
APP_URL=http://localhost:8080
```

## 🎯 Próximos Pasos

Una vez que tengas el sistema funcionando:

1. **Personalizar**: Agregar tus categorías reales
2. **Importar datos**: Si tienes transacciones previas
3. **Configurar OAuth**: Para usar login real de Google
4. **Backup**: Configurar respaldos automáticos
5. **Despliegue**: Mover a un servidor de producción

## 📞 Ayuda Rápida

**¿No funciona algo?**
1. Verificar logs: `docker-compose logs`
2. Reiniciar: `docker-compose restart`
3. Verificar puertos: `netstat -tulpn | grep :8080`
4. Limpiar y reiniciar: `docker-compose down -v && docker-compose up -d`

**¿Necesitas más ayuda?**
- Revisar README.md completo
- Crear issue en GitHub
- Verificar documentación en `/specs`

---

🚀 **¡En menos de 5 minutos deberías tener tu sistema financiero personal funcionando!**

¿Todo funcionando? ⭐ ¡No olvides darle una estrella al proyecto en GitHub!
