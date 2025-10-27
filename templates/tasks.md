description: "Tareas del Sistema de Control Financiero Personal"
---

# Tasks: Sistema de Control Financiero Personal

**Input**: plan.md (requerido), spec.md (historias de usuario), data-model.md  
**Prerequisites**: plan.md (stack técnico y arquitectura)  
**Tests**: Pruebas unitarias y de integración básicas en la Fase 7  
**Project Type**: Web app (Backend en Go + Frontend HTML/JS)

---

## Phase 1: Setup (Configuración base)

**Purpose**: Inicializar el entorno de desarrollo y estructura del proyecto.

- [ ] T001 Crear estructura de carpetas `/cmd`, `/internal`, `/pkg`, `/assets`
- [ ] T002 Inicializar módulo Go (`go mod init`)
- [ ] T003 [ ] Instalar dependencias base: Gin, GORM, JWT, dotenv
- [ ] T004 Configurar conexión inicial a PostgreSQL
- [ ] T005 Crear archivo `.env` con variables de entorno básicas

---

## Phase 2: Foundational (Diseño de base de datos)

**Purpose**: Definir entidades, relaciones y datos iniciales.

- [ ] T006 Crear entidades principales:
  - Users  
  - Roles  
  - Categories  
  - Transactions  
  - Reports
- [ ] T007 Definir migraciones automáticas con GORM
- [ ] T008 [P] Insertar datos iniciales (roles, admin por defecto)
- [ ] T009 Crear servicios de acceso a datos (`internal/db/`)

**Checkpoint**: Base de datos funcional y conectada al backend.

---

## Phase 3: User Story 1 - Autenticación y Roles (Priority: P1) 🎯

**Goal**: Permitir autenticación segura y control por roles.  
**Independent Test**: Login, JWT y acceso a rutas protegidas.

- [ ] T010 Implementar login con JWT (`internal/auth/`)
- [ ] T011 [P] Crear middleware para control de roles (`internal/middleware/roles.go`)
- [ ] T012 Proteger rutas según rol (`internal/routes/`)
- [ ] T013 Endpoint para perfil del usuario actual (`/api/v1/profile`)
- [ ] T014 [P] Pruebas unitarias de autenticación y roles

**Checkpoint**: Autenticación y roles operativos.

---

## Phase 4: User Story 2 - Gestión de Finanzas (Priority: P2)

**Goal**: Permitir CRUD de categorías y registro de transacciones.  
**Independent Test**: Crear, editar, listar y eliminar operaciones financieras.

- [ ] T015 CRUD de categorías (solo admin)
- [ ] T016 Registrar ingresos y egresos
- [ ] T017 Validar montos positivos
- [ ] T018 Calcular y almacenar balance mensual
- [ ] T019 [P] Pruebas de integración de operaciones financieras

---

## Phase 5: User Story 3 - Reportes (Priority: P3)

**Goal**: Generar reportes visuales y automáticos.  
**Independent Test**: Consultar y visualizar reportes mensuales.

- [ ] T020 Generar reporte mensual automático (cron interno o servicio)
- [ ] T021 API para consulta de reportes (`/api/v1/reports`)
- [ ] T022 Generar gráficas (Chart.js o ApexCharts)
- [ ] T023 Guardar historial de reportes por usuario

**Checkpoint**: Reportes funcionales y exportables.

---

## Phase 6: User Story 4 - Interfaz Web (Priority: P4)

**Goal**: Crear interfaz moderna, responsiva y accesible.  
**Independent Test**: Navegación completa desde navegador, con modo oscuro/claro.

- [ ] T024 Implementar frontend responsivo (HTML5, JS moderno)
- [ ] T025 Integrar frontend con API REST
- [ ] T026 Diseñar menú horizontal y tema pastel
- [ ] T027 Implementar modo oscuro/claro/sistema
- [ ] T028 Sección de perfil y actualización de datos
- [ ] T006 Usar modades de confirmacion para acciones como eliminar, registrar, etc. [https://sweetalert.js.org/guides/]
---

## Phase 7: Polish & Cross-Cutting Concerns (Despliegue y Pruebas Finales)

**Purpose**: Empaquetar, probar y desplegar.

- [ ] T029 Crear `docker-compose.yml` (app + db)
- [ ] T030 Agregar `.env.example`
- [ ] T031 [P] Pruebas unitarias y de integración
- [ ] T032 Documentar API con Swagger o Redoc
- [ ] T033 [P] Crear manual de instalación y uso
- [ ] T034 [P] Validar `quickstart.md` (inicio rápido del proyecto)

---

## Dependencies & Execution Order

- **Setup (Fase 1)**: No depende de otras fases.  
- **Foundational (Fase 2)**: Depende de la Fase 1.  
- **User Stories (Fases 3–6)**: Dependen de la Fase 2.  
- **Polish (Fase 7)**: Depende de la finalización de las anteriores.

---

## Parallel Opportunities

- [P] Tareas que pueden ejecutarse en paralelo:
  - Instalación de dependencias  
  - Migraciones de base de datos  
  - Pruebas y documentación  
  - Integración frontend-backend

---

## Implementation Strategy

1. Completar Fase 1 y Fase 2 (bloqueantes).  
2. Implementar Autenticación (Fase 3) como MVP.  
3. Agregar Finanzas (Fase 4), Reportes (Fase 5) y UI (Fase 6).  
4. Desplegar y probar (Fase 7).

---

**Checkpoint final:**  
Sistema operativo con autenticación, gestión financiera, reportes y UI funcional.  
Listo para despliegue con Docker, GitHub actions, Aws y documentación.
