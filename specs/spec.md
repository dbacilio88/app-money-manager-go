# Especificación del Sistema — Control Financiero (Detallada)

## 1. Resumen ejecutivo
Aplicación web para registro y gestión de actividades financieras personales y de negocio (ingresos, egresos, préstamos, alquileres, pagos recurrentes). Construida con **Go 1.25.3** (framework Gin) y **MongoDB**, frontend HTML/CSS/JS, autenticación vía JWT/refresh tokens y login con Google (OAuth).

## 2. Requisitos funcionales (priorizados)
1. Autenticación y gestión de sesión
   - Registro de usuario (requiere aprobación del administrador)
   - Login con Google OAuth y credenciales tradicionales (correo + contraseña)
   - JWT de acceso corto (p. ej. 15 min) + refresh token (p. ej. 7 días)
   - Modal en UI para renovar token antes de expiración
2. Gestión de usuarios (CRUD) — admin
   - Listar, aprobar, suspender, cambiar roles
   - Activar y desactivar usuarios
   - Cambiar contraseña desde perfil
3. Gestión de categorías (CRUD)
4. Gestión de transacciones (CRUD)
   - Tipos: ingreso, egreso, préstamo (subtipo: recibido/pagado), alquiler, otros
   - Campos: monto, moneda, categoría, fecha, descripción, cuenta (opcional), etiquetas
5. Dashboard y reportes
   - Gráficas: gasto por categoría, evolución temporal, balance acumulado
   - Filtro por rango, categoría, cuenta
6. Auditoría y logs
   - Log de acciones críticas (creación/edición/eliminación/approvals)
7. API REST versionada (/api/v1)

## 3. Requisitos no funcionales
- Seguridad: cifrado de contraseñas (bcrypt), HTTPS en producción, protección CSRF en formularios importantes
- Rendimiento: respuesta API < 200ms para operaciones comunes en cargas normales
- Escalabilidad: diseño para escalado horizontal (stateless backend + MongoDB replicaset)
- Disponibilidad y resiliencia: run local via Docker; backups planificados (externo)

## 4. Contratos API (ejemplos)
- POST /api/v1/auth/register -> {email, nombre, password} 201 creado (estado: pending)
- GET /api/v1/auth/google -> {url_autorizacion}
- GET /api/v1/auth/google/callback?code=... -> redirige a frontend con token
- POST /api/v1/auth/login -> {email,password} -> {access_token, refresh_token}
- POST /api/v1/auth/refresh -> {refresh_token} -> {access_token}
- GET /api/v1/usuarios (admin)
- PATCH /api/v1/usuarios/:id/activar (admin) -> activa un usuario
- PATCH /api/v1/usuarios/:id/desactivar (admin) -> desactiva un usuario
- POST /api/v1/transacciones -> crear transacción
- GET /api/v1/reportes/mes?year=YYYY&month=MM -> devuelve métricas y series para gráficas

## 5. Criterios de aceptación (ejemplos)
- Registro crea usuario con estado `pending` y notifica al admin
- Admin puede aprobar y activar usuarios
- Admin puede desactivar usuarios (cambiando estado a `suspended` o `inactive`)
- Un usuario desactivado no puede iniciar sesión
- Un usuario autenticado solo puede acceder a sus transacciones
- Refresh token renueva access token sin requerir re-login si es válido

## 6. Riesgos y mitigaciones
- Riesgo: exposición de tokens → Mitigación: store refresh tokens en httpOnly cookies o en DB y revocación
- Riesgo: inyección/validación → Mitigación: validaciones estrictas, uso de drivers oficiales de MongoDB

## 7. Roadmap (alto nivel)
- MVP (Sprints 1-3): auth, CRUD usuarios/transacciones/categorías, dashboard básico
- Fase 2 (Sprints 4-6): préstamos, pagos recurrentes, reportes avanzados, export CSV/PDF
- Fase 3: CI/CD, despliegue en AWS, monitorización y alertas

## 8. Entregables iniciales
- Código backend (Go 1.25.3 + Gin) con tests básicos
- Base de datos MongoDB con colecciones e índices optimizados
- Frontend responsive y accesible con login y dashboard
- Docker + docker-compose + Makefile
- Documentación técnica en `specs/` y diagramas C4 en `docs/`

## 9. Colaboración y Herramientas de Desarrollo

### Repositorio
- **GitHub**: https://github.com/dbacilio88/app-money-manager-go.git
- **Branch principal**: main
- **Estrategia**: GitFlow con feature branches y pull requests obligatorios

### Herramientas de Automatización y Calidad

- 🤖 **Renovate** (https://github.com/apps/renovate):
  Automatiza la actualización de dependencias (Go modules, Docker images, GitHub Actions) manteniendo el proyecto actualizado y seguro.

- 🤖 **CodeRabbit AI** (https://github.com/apps/coderabbitai):
  Realiza revisión automatizada de código en Pull Requests, sugiere mejoras de calidad, detecta code smells y valida adherencia a best practices.

- 🧠 **GitHub Copilot (Code Review Integration)**:
  Asiste en la revisión de código con sugerencias contextuales inteligentes, ayudando a mantener los estándares definidos en la constitución del proyecto.

- 🔒 **GitHub Advanced Security**:
  Proporciona análisis continuo de vulnerabilidades, escaneo de secretos expuestos, revisión de dependencias inseguras y alertas de seguridad en tiempo real.

### Flujo de Trabajo
1. Crear feature branch desde `main`
2. Desarrollar con tests (go test)
3. Commit siguiendo Conventional Commits
4. Push y crear Pull Request
5. Revisión automatizada (CodeRabbit AI, GitHub Actions CI)
6. Revisión manual del equipo
7. Merge a `main` tras aprobación
8. Renovate mantiene dependencias actualizadas automáticamente
