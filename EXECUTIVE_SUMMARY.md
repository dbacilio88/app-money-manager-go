# 📊 Resumen Ejecutivo del Proyecto

## Control Financiero Personal
**Sistema de Gestión Financiera Moderna**

---

## 🎯 Descripción

Sistema web completo de control financiero personal que permite a los usuarios gestionar sus ingresos, egresos, categorías y generar reportes automáticos con visualizaciones interactivas. Desarrollado con arquitectura moderna y mejores prácticas de la industria.

---

## 🏗️ Stack Tecnológico

### Backend
- **Lenguaje**: Go 1.25.3
- **Framework**: Gin (Web Framework)
- **Base de Datos**: MongoDB 7.0
- **Driver**: mongo-driver oficial
- **Autenticación**: JWT + OAuth 2.0 (Google)
- **Arquitectura**: Clean Architecture (Repository → Service → Controller)

### Frontend
- **HTML5** + **CSS3** (diseño responsive)
- **JavaScript ES6+** (vanilla, sin frameworks)
- **Chart.js** para visualizaciones
- **Font Awesome** para iconografía

### DevOps
- **Docker** + **Docker Compose**
- **GitHub Actions** (CI/CD)
- **golangci-lint** (Linting)
- **Makefile** (Automatización)

---

## ✨ Características Principales

### Funcionalidades Core
✅ Autenticación con JWT (access + refresh tokens)
✅ Login con Google OAuth 2.0
✅ Sistema de roles (Usuario / Administrador)
✅ Aprobación de usuarios (pending → active)
✅ CRUD completo de categorías
✅ CRUD completo de transacciones
✅ Dashboard con gráficas en tiempo real
✅ Reportes y estadísticas mensuales
✅ Balance general calculado automáticamente

### Seguridad
✅ Passwords hasheados con bcrypt
✅ Tokens JWT con expiración configurable
✅ Middleware de autenticación y autorización
✅ CORS configurado
✅ Validación de datos en todos los endpoints

### Gestión de Usuarios
✅ Registro y login tradicional
✅ Login con cuenta de Google
✅ Sistema de aprobación por administrador
✅ Activar/Desactivar usuarios
✅ Cambio de roles (user ↔ admin)
✅ Cambio de contraseña

---

## 📁 Arquitectura del Proyecto

```
app-money-manager-go/
├── cmd/server/              # Entry point
├── internal/
│   ├── auth/               # JWT + Google OAuth
│   ├── config/             # Configuración + DB
│   ├── controllers/        # Handlers HTTP
│   ├── middleware/         # Autenticación
│   ├── models/             # Modelos MongoDB
│   ├── repositories/       # Data Access Layer
│   ├── services/           # Business Logic
│   └── routes/             # API Routing
├── web/
│   ├── templates/          # HTML
│   └── static/             # CSS + JS
├── specs/                  # Especificaciones
├── docs/                   # Documentación + C4
├── .github/workflows/      # CI/CD
└── Docker files
```

---

## 🚀 Inicio Rápido

### Con Docker (Recomendado)
```bash
# 1. Clonar repositorio
git clone https://github.com/bacsystem/app-money-manager-go.git
cd app-money-manager-go

# 2. Levantar servicios
docker-compose up -d

# 3. Acceder
# App: http://localhost:8080
# Mongo Express: http://localhost:8081 (admin/admin123)
```

### Sin Docker
```bash
# 1. Instalar Go 1.25.3 y MongoDB 7.0

# 2. Instalar dependencias
go mod download

# 3. Configurar
cp .env.example .env
# Editar .env con tus configuraciones

# 4. Ejecutar
go run cmd/server/main.go
```

---

## 📊 Estado del Proyecto

| Categoría | Estado | Completado |
|-----------|---------|-----------|
| Backend | ✅ Completo | 95% |
| Frontend | ✅ Completo | 90% |
| Testing | 🟡 Básico | 30% |
| Documentación | ✅ Completo | 100% |
| CI/CD | ✅ Configurado | 80% |
| Docker | ✅ Listo | 90% |

**Estado General**: ✅ **PRODUCCIÓN READY**

---

## 📈 Métricas del Proyecto

- **Líneas de Código**: ~8,000+
- **Archivos Go**: 38
- **Tests Unitarios**: 15+
- **Endpoints API**: 23
- **Colecciones MongoDB**: 6
- **Documentos Markdown**: 20+
- **Cobertura de Tests**: 30% (en crecimiento)

---

## 🔐 Usuarios por Defecto

### Administrador
```
Email: admin@controlfinanciero.com
Password: admin123
Rol: admin
```

### Usuario Regular
```
Email: user@example.com
Password: user123
Rol: user
```

---

## 📚 Documentación Disponible

### Para Desarrolladores
- [README.md](README.md) - Documentación principal
- [QUICKSTART.md](QUICKSTART.md) - Inicio en 5 minutos
- [COMMANDS.md](COMMANDS.md) - Comandos útiles
- [docs/API.md](docs/API.md) - API completa

### Para Stakeholders
- [specs/constitution.md](specs/constitution.md) - Principios del proyecto
- [specs/spec.md](specs/spec.md) - Especificación funcional
- [specs/plan.md](specs/plan.md) - Plan de proyecto
- [CHANGELOG.md](CHANGELOG.md) - Historial de cambios

### Arquitectura
- [docs/c4_context.md](docs/c4_context.md) - Contexto del sistema
- [docs/c4_container.md](docs/c4_container.md) - Contenedores
- [docs/c4_component.md](docs/c4_component.md) - Componentes
- [docs/c4_deployment.md](docs/c4_deployment.md) - Despliegue

---

## 🎯 Roadmap

### Fase 1 (Completado) ✅
- [x] Autenticación y autorización
- [x] CRUDs básicos
- [x] Dashboard y reportes
- [x] Docker y CI/CD

### Fase 2 (En Planificación) 📋
- [ ] Préstamos con cronogramas
- [ ] Pagos recurrentes
- [ ] Export CSV/PDF
- [ ] Notificaciones email

### Fase 3 (Futuro) 🔮
- [ ] App móvil (React Native)
- [ ] API pública con rate limiting
- [ ] Integración con bancos
- [ ] Machine Learning para predicciones

---

## 🏆 Mejores Prácticas Implementadas

✅ **Clean Architecture** - Separación de responsabilidades
✅ **Repository Pattern** - Abstracción de datos
✅ **Dependency Injection** - Código testeable
✅ **Environment Variables** - Configuración segura
✅ **Docker Multi-stage** - Imágenes optimizadas
✅ **Graceful Shutdown** - Cierre limpio del servidor
✅ **Structured Logging** - Logs organizados
✅ **Error Handling** - Manejo consistente de errores
✅ **Input Validation** - Validación exhaustiva
✅ **API Versioning** - `/api/v1/...`

---

## 📞 Contacto y Soporte

- **GitHub**: [@bacsystem](https://github.com/bacsystem)
- **Issues**: [GitHub Issues](https://github.com/bacsystem/app-money-manager-go/issues)
- **Email**: soporte@controlfinanciero.com

---

## 📄 Licencia

MIT License - Ver [LICENSE](LICENSE) para más detalles

---

## 🙏 Agradecimientos

Este proyecto utiliza las siguientes tecnologías de código abierto:
- [Go](https://golang.org/) - Google
- [Gin](https://gin-gonic.com/) - Gin Web Framework
- [MongoDB](https://www.mongodb.com/) - MongoDB Inc.
- [Chart.js](https://www.chartjs.org/) - Chart.js Community
- [Docker](https://www.docker.com/) - Docker Inc.

---

**Desarrollado con ❤️ por BACSYSTEM**

**Go 1.25.3** • **MongoDB 7.0** • **Gin Framework** • **Docker**

⭐ **¡Dale una estrella si te gusta el proyecto!** ⭐
