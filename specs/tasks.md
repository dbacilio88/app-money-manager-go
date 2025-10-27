# Tasks — Lista detallada de tareas (nivel técnico)

Este archivo prioriza las tareas técnicas que guiarán las implementaciones por sprint.

**Stack**: Backend **Go 1.25.3** + Base de datos **MongoDB**.

MVP (Sprint 1)
- [ ] Scaffold del repo con **Go 1.25.3** (estructura de carpetas, go.mod, main.go)
- [ ] Conexión a **MongoDB** (pool, variables de entorno)
- [ ] Implementar modelos (usuarios, roles, categorias, transacciones)
- [ ] Auth: register/login, password hashing (bcrypt), JWT issuance, refresh token storage
- [ ] Endpoint admin para aprobar usuarios (estado pending → active)
- [ ] Tests unitarios iniciales para autenticación

MVP (Sprint 2)
- [ ] CRUD categorías (endpoints + UI)
- [ ] CRUD transacciones (endpoints + UI)
- [ ] Dashboard básico: balance y últimos movimientos
- [ ] Validaciones y manejo de errores centralizado

MVP (Sprint 3)
- [ ] Google OAuth: endpoint para obtener URL, callback handler
- [ ] Frontend: redirección a Google y captura del token
- [ ] Modal para refresh token y renovación desde UI
- [ ] Endpoints de reportes (mes, categoría)

Fase 2 — Avanzado
- [ ] Préstamos: modelo y CRUD, cron para pagos
- [ ] Alquileres y pagos recurrentes
- [ ] Export CSV/PDF
- [ ] Optimización de consultas y agregaciones para reportes

Infra y DevOps
- [ ] Dockerfile y docker-compose (app + mongo)
- [ ] Makefile para tareas comunes (build/run/test)
- [ ] GitHub Actions: CI (tests, lint, build)
- [ ] Plantilla de despliegue en AWS (CloudFormation / Terraform) — Fase posterior
- [ ] Configurar Renovate para actualización automática de dependencias
- [ ] Configurar CodeRabbit AI para revisión de código automatizada
- [ ] Habilitar GitHub Advanced Security (Dependabot, secret scanning, code scanning)

Repositorio y Colaboración
---------------------------
- **Repositorio**: https://github.com/dbacilio88/app-money-manager-go.git
- **Herramientas integradas**:
  - 🤖 Renovate: Actualización automática de dependencias
  - 🤖 CodeRabbit AI: Revisión automatizada de PRs
  - 🧠 GitHub Copilot: Asistencia en desarrollo y revisión
  - 🔒 GitHub Advanced Security: Análisis de seguridad continuo

Notas
-----
- Cada ítem debe tener una issue asociada en GitHub y una PR con revisión.
- Las PRs requieren aprobación de al menos un revisor humano además de las verificaciones automatizadas.
- Seguir Conventional Commits para mensajes de commit (feat:, fix:, docs:, etc.).
