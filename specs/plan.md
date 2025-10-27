# Plan de Proyecto — Control Financiero

Versión: 0.1
Resumen: Plan de trabajo por fases y sprints para entregar un MVP funcional y preparar despliegue.

**Stack tecnológico**: **Go 1.25.3** (framework Gin) + **MongoDB** como base de datos.

Fase 0 — Preparación (1 semana)
- Definir requisitos finales y aprobar especificaciones (docs en `specs/`).
- Scaffolding del repo con **Go 1.25.3**, configuración CI básica (lint/tests), entorno Docker local (**MongoDB** + app).

Fase 1 — MVP (3-4 semanas)
Sprint 1 (1 semana):
- Autenticación básica (register/login) con JWT y refresh token.
- Esqueleto de usuarios y roles, endpoint de aprobación de usuarios por admin.
- Conexión a **MongoDB** y pruebas locales.

Sprint 2 (1 semana):
- CRUD de categorías.
- CRUD de transacciones: ingreso/egreso simple.
- Dashboard básico (balance, últimos movimientos).

Sprint 3 (1-2 semanas):
- Login con Google (OAuth) y flujo de callback.
- Modal de refresh token en frontend y expiración visual.
- Tests básicos y documentación de endpoints.

Fase 2 — Funcionalidades Avanzadas (3 semanas)
- Préstamos: manejar cronograma de pagos y amortización.
- Alquileres y pagos recurrentes: crear y gestionar recurrencias.
- Reportes enriquecidos y export CSV/PDF.

Fase 3 — Harden & Deploy (2 semanas)
- Docker production optimizado, Compose y Makefile.
- Preparar GitHub Actions CI/CD para despliegue en AWS (esqueleto).
- Revisión de seguridad, auditoría y tests de integración.

Entregables por sprint
- Código en rama feature/* con PR y revisiones.
- Tests unitarios e integración mínima.
- Documentación actualizada en `specs/` y `docs/`.

Métricas de éxito por sprint
- Cobertura básica de pruebas (objetivo inicial >= 40%).
- Pipeline CI ejecutado y limpiado con éxito.
- Demo funcional con auth y CRUD básicos.

Colaboración y Automatización
------------------------------
El proyecto utiliza herramientas modernas para mantener calidad y seguridad:

- 🤖 **Renovate**: Actualización automática de dependencias (Go modules, Docker, GitHub Actions)
- 🤖 **CodeRabbit AI**: Revisión automatizada de código en Pull Requests
- 🧠 **GitHub Copilot**: Asistencia en revisión de código y desarrollo
- 🔒 **GitHub Advanced Security**: Análisis de vulnerabilidades y secretos

Repositorio y Flujo de Trabajo
-------------------------------
- **Repositorio**: https://github.com/dbacilio88/app-money-manager-go.git
- **Estrategia**: GitFlow (feature/*, hotfix/*, main)
- **Pull Requests**: Obligatorios con revisión automatizada y manual
- **CI/CD**: GitHub Actions ejecuta tests, linting y security scans en cada PR
