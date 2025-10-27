# Control Financiero Personal 💰

![Go Version](https://img.shields.io/badge/Go-1.25.3-00ADD8?style=flat&logo=go)
![MongoDB](https://img.shields.io/badge/MongoDB-7.0-47A248?style=flat&logo=mongodb)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Status](https://img.shields.io/badge/Status-Production%20Ready-success)

Sistema moderno de gestión financiera personal desarrollado en **Go 1.25.3** con **MongoDB** e interfaz web responsiva. Permite el control de ingresos, egresos, categorías y genera reportes automáticos con visualizaciones interactivas.

## 🚀 Características Principales

- ✅ **Autenticación Google OAuth**: Login seguro con cuentas de Google
- ✅ **Gestión de Transacciones**: Registro de ingresos y egresos con categorías
- ✅ **Balance en Tiempo Real**: Cálculo automático del balance actual
- ✅ **Reportes Mensuales**: Generación automática de reportes con gráficas
- ✅ **Interfaz Moderna**: Diseño responsivo con modo claro/oscuro
- ✅ **Roles de Usuario**: Administrador y Usuario con permisos diferenciados
- ✅ **API RESTful**: Backend completamente documentado
- ✅ **Dockerizado**: Despliegue fácil con Docker Compose

## 🛠️ Tecnologías Utilizadas

### Backend
- **Go 1.25.3** - Lenguaje principal
- **Gin Framework** - Framework web HTTP
- **MongoDB Driver** - Driver oficial de MongoDB para Go
- **MongoDB** - Base de datos NoSQL
- **JWT** - Autenticación con tokens
- **Google OAuth 2.0** - Login con Google

### Frontend
- **HTML5 + CSS3** - Estructura y estilos
- **JavaScript ES6+** - Lógica del frontend
- **Chart.js** - Gráficas interactivas
- **Font Awesome** - Iconografía

### DevOps
- **Docker** - Contenedorización
- **Docker Compose** - Orquestación
- **MongoDB 7.0** - Base de datos NoSQL

## 📁 Estructura del Proyecto

```
control-financiero/
├── cmd/server/           # Punto de entrada de la aplicación
├── internal/            # Código interno de la aplicación
│   ├── auth/           # Sistema de autenticación JWT y OAuth
│   ├── config/         # Configuración de BD y variables
│   ├── controllers/    # Controladores HTTP
│   ├── middleware/     # Middlewares de autenticación
│   ├── models/         # Modelos de datos MongoDB
│   ├── repositories/   # Capa de acceso a datos
│   ├── services/       # Lógica de negocio
│   ├── routes/         # Definición de rutas
│   └── services/       # Lógica de negocio
├── pkg/utils/          # Utilidades compartidas
├── web/                # Frontend web
│   ├── static/         # CSS, JS, imágenes
│   └── templates/      # Plantillas HTML
├── assets/             # Recursos estáticos (fuentes, logos, videos)
├── scripts/            # Scripts de base de datos
├── specs/              # Documentación del proyecto
├── docker-compose.yml  # Configuración Docker
├── Dockerfile          # Imagen de la aplicación
└── README.md          # Este archivo
```

## 🚦 Inicio Rápido

### 1. Prerrequisitos

- **Docker** y **Docker Compose** instalados
- **Git** para clonar el repositorio
- (Opcional) **Go 1.22+** para desarrollo local

### 2. Clonar el Repositorio

```bash
git clone https://github.com/tu-usuario/control-financiero.git
cd control-financiero
```

### 3. Configurar Variables de Entorno

```bash
# Copiar archivo de ejemplo
cp .env.example .env

# Editar variables según tu entorno
nano .env
```

**Variables importantes a configurar:**

```env
# Google OAuth (Obtener en Google Cloud Console)
GOOGLE_CLIENT_ID=tu_google_client_id_real
GOOGLE_CLIENT_SECRET=tu_google_client_secret_real
GOOGLE_REDIRECT_URL=http://localhost:8080/auth/google/callback

# JWT Secret (Generar uno seguro)
JWT_SECRET=tu_jwt_secret_muy_seguro_aqui

# Base de Datos (Usar valores por defecto para desarrollo)
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=control_financiero
```

### 4. Ejecutar con Docker Compose

```bash
# Construir e iniciar todos los servicios
docker-compose up --build -d

# Ver logs en tiempo real
docker-compose logs -f app
```

### 5. Acceder a la Aplicación

- **Aplicación Web**: http://localhost:8080
- **API Health Check**: http://localhost:8080/api/v1/health
- **Adminer (BD)**: http://localhost:8081

## 🔧 Desarrollo Local

### Ejecutar sin Docker

1. **Instalar PostgreSQL localmente**
2. **Configurar variables de entorno**:
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   # ... otras variables
   ```
3. **Instalar dependencias de Go**:
   ```bash
   go mod download
   ```
4. **Ejecutar la aplicación**:
   ```bash
   go run cmd/server/main.go
   ```

### Comandos Útiles para Desarrollo

```bash
# Ejecutar tests
go test ./...

# Formatear código
go fmt ./...

# Verificar código
go vet ./...

# Construir binario
go build -o bin/server cmd/server/main.go

# Ejecutar con live reload (instalar air primero)
air
```

## 📚 Uso de la Aplicación

### 1. Primer Acceso

1. Abrir http://localhost:8080
2. Hacer clic en "Continuar con Google"
3. Completar autenticación OAuth
4. Establecer balance inicial (opcional)

### 2. Gestión de Categorías

- Crear categorías para organizar transacciones
- Ejemplos: Alimentación, Transporte, Entretenimiento, Salud

### 3. Registro de Transacciones

- **Ingresos**: Salario, bonos, inversiones
- **Egresos**: Gastos diarios categorizados
- El balance se actualiza automáticamente

### 4. Visualización de Reportes

- Dashboard con estadísticas del mes actual
- Gráficas de gastos por categoría
- Evolución mensual de ingresos/egresos
- Reportes históricos mensuales

## 🔐 API Endpoints

### Autenticación
- `GET /api/v1/auth/google` - Iniciar login con Google
- `GET /api/v1/auth/google/callback` - Callback de Google OAuth
- `GET /api/v1/perfil` - Obtener perfil del usuario
- `PUT /api/v1/perfil` - Actualizar perfil

### Categorías
- `GET /api/v1/categorias` - Listar categorías del usuario
- `POST /api/v1/categorias` - Crear nueva categoría
- `PUT /api/v1/categorias/{id}` - Actualizar categoría
- `DELETE /api/v1/categorias/{id}` - Eliminar categoría (Admin)

### Transacciones
- `GET /api/v1/transacciones` - Listar transacciones con paginación
- `POST /api/v1/transacciones` - Crear nueva transacción
- `PUT /api/v1/transacciones/{id}` - Actualizar transacción
- `DELETE /api/v1/transacciones/{id}` - Eliminar transacción (Admin)

### Balance
- `GET /api/v1/balance` - Obtener balance actual
- `POST /api/v1/balance/inicial` - Establecer balance inicial

### Reportes
- `GET /api/v1/reportes` - Listar todos los reportes
- `GET /api/v1/reportes/actual` - Reporte del mes actual
- `GET /api/v1/reportes/mes?mes=X&anio=Y` - Reporte de mes específico
- `GET /api/v1/reportes/estadisticas` - Estadísticas generales

## 👥 Roles y Permisos

### Usuario Regular
- ✅ Crear/editar sus propias categorías y transacciones
- ✅ Ver sus reportes y estadísticas
- ✅ Actualizar su perfil
- ❌ No puede eliminar registros
- ❌ No puede ver datos de otros usuarios

### Administrador
- ✅ Todos los permisos de usuario regular
- ✅ Eliminar categorías y transacciones
- ✅ Ver datos de todos los usuarios
- ✅ Gestionar roles de usuarios
- ✅ Generar reportes automáticos

## 🐳 Docker y Despliegue

### Servicios Incluidos

- **app**: Aplicación Go principal (puerto 8080)
- **postgres**: Base de datos PostgreSQL (puerto 5432)
- **adminer**: Interfaz web para BD (puerto 8081)

### Comandos Docker Útiles

```bash
# Ver logs de un servicio específico
docker-compose logs -f app

# Reconstruir solo la aplicación
docker-compose build app

# Ejecutar comando en contenedor
docker-compose exec app sh

# Hacer backup de la BD
docker-compose exec postgres pg_dump -U postgres control_financiero > backup.sql

# Restaurar backup
docker-compose exec -T postgres psql -U postgres control_financiero < backup.sql
```

## 🔧 Configuración de Google OAuth

### 1. Crear Proyecto en Google Cloud Console

1. Ir a [Google Cloud Console](https://console.cloud.google.com/)
2. Crear nuevo proyecto o seleccionar existente
3. Habilitar Google+ API

### 2. Configurar OAuth Consent Screen

1. Ir a "APIs & Services" > "OAuth consent screen"
2. Seleccionar "External" y completar información básica
3. Agregar scopes: `email`, `profile`

### 3. Crear Credenciales OAuth

1. Ir a "APIs & Services" > "Credentials"
2. Crear "OAuth 2.0 Client ID"
3. Tipo: "Web application"
4. URIs autorizadas: `http://localhost:8080/auth/google/callback`
5. Copiar Client ID y Client Secret al archivo `.env`

## 📊 Base de Datos

### Esquema Principal

```sql
-- Roles (Admin, Usuario)
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) UNIQUE NOT NULL
);

-- Usuarios del sistema
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    correo VARCHAR(255) UNIQUE NOT NULL,
    fecha_nacimiento DATE,
    telefono VARCHAR(20),
    foto TEXT,
    google_id VARCHAR(255) UNIQUE,
    rol_id INTEGER REFERENCES roles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categorías de transacciones
CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    descripcion TEXT,
    creado_por INTEGER REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Transacciones (ingresos y egresos)
CREATE TABLE transacciones (
    id SERIAL PRIMARY KEY,
    tipo VARCHAR(10) CHECK (tipo IN ('Ingreso', 'Egreso')),
    monto DECIMAL(15,2) NOT NULL CHECK (monto > 0),
    descripcion TEXT,
    fecha DATE NOT NULL,
    categoria_id INTEGER REFERENCES categorias(id),
    usuario_id INTEGER REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Balance de usuarios
CREATE TABLE balances (
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER UNIQUE REFERENCES usuarios(id),
    inicial DECIMAL(15,2) DEFAULT 0,
    actual DECIMAL(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reportes mensuales
CREATE TABLE reportes (
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER REFERENCES usuarios(id),
    mes INTEGER CHECK (mes BETWEEN 1 AND 12),
    anio INTEGER,
    total_ingresos DECIMAL(15,2) DEFAULT 0,
    total_egresos DECIMAL(15,2) DEFAULT 0,
    balance_total DECIMAL(15,2) DEFAULT 0,
    fecha_generacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🧪 Testing

### Ejecutar Tests

```bash
# Todos los tests
go test ./...

# Tests con cobertura
go test -cover ./...

# Tests de integración con BD de prueba
go test -tags=integration ./...
```

### Ejemplo de Test de API

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Login (requiere configuración OAuth)
curl -X POST http://localhost:8080/api/v1/auth/google

# Crear categoría (requiere JWT token)
curl -X POST http://localhost:8080/api/v1/categorias \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Alimentación","descripcion":"Gastos en comida"}'
```

## 🚨 Troubleshooting

### Problemas Comunes

1. **Error de conexión a BD**:
   ```bash
   # Verificar que PostgreSQL esté corriendo
   docker-compose ps

   # Ver logs de BD
   docker-compose logs postgres
   ```

2. **Google OAuth no funciona**:
   - Verificar `GOOGLE_CLIENT_ID` y `GOOGLE_CLIENT_SECRET`
   - Confirmar URL de redirect en Google Console
   - Verificar que la app esté en modo de testing/producción

3. **Frontend no carga**:
   - Verificar que los assets estén montados correctamente
   - Revisar logs del contenedor app
   - Confirmar permisos de archivos

4. **JWT inválido**:
   - Verificar `JWT_SECRET` en variables de entorno
   - Limpiar localStorage del navegador
   - Reiniciar sesión

### Logs y Debugging

```bash
# Ver todos los logs
docker-compose logs

# Logs en tiempo real de la app
docker-compose logs -f app

# Acceder al contenedor para debugging
docker-compose exec app sh

# Ver estado de servicios
docker-compose ps
```

## 🤝 Contribución

1. **Fork** el repositorio
2. Crear **rama de feature**: `git checkout -b feature/nueva-funcionalidad`
3. **Commit** cambios: `git commit -am 'Agregar nueva funcionalidad'`
4. **Push** a la rama: `git push origin feature/nueva-funcionalidad`
5. Crear **Pull Request**

### Estándares de Código

- Seguir convenciones de Go (gofmt, golint)
- Documentar funciones públicas
- Incluir tests para nuevas funcionalidades
- Mantener cobertura de tests >80%

## � Documentación Completa

Este proyecto cuenta con documentación exhaustiva:

### Guías de Usuario
- 📖 [README.md](README.md) - Documentación principal del proyecto
- ⚡ [QUICKSTART.md](QUICKSTART.md) - Guía de inicio rápido (5 minutos)
- 💻 [COMMANDS.md](COMMANDS.md) - Lista de comandos útiles para desarrollo
- 📋 [CHANGELOG.md](CHANGELOG.md) - Registro de cambios y actualizaciones

### Especificaciones Técnicas
- 📐 [Constitution](specs/constitution.md) - Principios y restricciones del proyecto
- 📝 [Specification](specs/spec.md) - Especificación funcional completa
- 🗄️ [Data Model](specs/data-model.md) - Modelo de datos MongoDB
- 📅 [Project Plan](specs/plan.md) - Plan de proyecto por fases
- ✅ [Tasks](specs/tasks.md) - Lista detallada de tareas técnicas

### Documentación de API
- 🌐 [API Documentation](docs/API.md) - 23 endpoints documentados con ejemplos

### Arquitectura C4
- 🏗️ [C4 Context](docs/c4_context.md) - Diagrama de contexto del sistema
- 📦 [C4 Container](docs/c4_container.md) - Contenedores y tecnologías
- ⚙️ [C4 Component](docs/c4_component.md) - Componentes internos
- 🚀 [C4 Deployment](docs/c4_deployment.md) - Arquitectura de despliegue

## �📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 📞 Soporte

- **Issues**: [GitHub Issues](https://github.com/tu-usuario/control-financiero/issues)
- **Documentación**: Ver carpeta `/specs` y `/docs`
- **Email**: soporte@controlfinanciero.com

---

**Desarrollado con ❤️ por el equipo de BACSYSTEM**

**Stack**: Go 1.25.3 • MongoDB 7.0 • Gin Framework • Docker

🌟 Si te gusta este proyecto, ¡no olvides darle una estrella en GitHub!
