# Constitución del Proyecto — Control Financiero

Fecha: 2025-10-27
Versión: 0.1 (Borrador)
Autor: Equipo de Desarrollo

Resumen
-------
Este documento define la visión, objetivos, alcance, restricciones y criterios de éxito del sistema "Control Financiero". El propósito es alinear a los interesados y recursos técnicos antes de comenzar la implementación.

Visión
------
Proveer a individuos y pequeñas empresas una aplicación segura, privada y fácil de usar para registrar y gestionar actividades financieras personales y de negocio: ingresos, egresos, préstamos, alquileres y otras transacciones recurrentes. El sistema debe ser accesible mediante una interfaz web moderna, segura y responsiva.

Objetivos Principales
---------------------
- Registrar y categorizar transacciones financieras (ingresos/egresos).
- Manejar préstamos y flujos relacionados (préstamos recibidos, pagados, intereses).
- Soportar CRUD de usuarios, transacciones y categorías.
- Autenticación segura con JWT y refresh tokens; expiración controlada y renovación desde UI con modal.
- Aprobación administrativa de nuevos usuarios registrados.
- Dashboard con métricas y gráficas (mensuales, por categoría, por cuenta).
- Despliegue mediante Docker y orquestación con docker-compose; integración CI/CD en GitHub Actions para despliegue en AWS (fase posterior).

Alcance
-------
Incluye:
- Backend en Go con Gin.
- Persistencia en MongoDB.
- Frontend en HTML/CSS/JS (sin frameworks SPA pesados) para máxima portabilidad.
- API REST con endpoints versionados (/api/v1).
- JWT + refresh token.
- Docker + docker-compose.

Exclusiones (por ahora)
-----------------------
- Integración completa con pasarelas de pago externas.
- Escalado multi-región en AWS (se planifica en fase posterior).

Usuarios y Roles
----------------
- Administrador: aprueba usuarios, gestiona roles y puede ver reportes globales.
- Usuario autenticado: gestiona sus propias transacciones, categorías y perfiles.
- Invitado/Visitante: puede registrarse (requiere aprobación por admin).

Criterios de Éxito
------------------
- API documentada y testeable.
- Interfaz usable, responsive y accesible en móvil y escritorio.
- JWT seguro con expiración y refresh correcto.
- Docker y docker-compose para entorno local reproducible.

Restricciones Técnicas
---------------------
- Lenguaje: Go (>= 1.25.3)
- Framework web: Gin
- BD: MongoDB (imagen oficial)
- Frontend ligero: HTML/CSS/vanilla JS + Chart.js
- CI/CD: GitHub Actions (esqueleto en esta versión)

Políticas de Seguridad (básicas)
-------------------------------
- No exponer secretos en el repo (usar variables de entorno).
- TLS en entornos de producción.
- Validación y escape de datos de entrada.

Gobernanza y Entrega
--------------------
El proyecto seguirá entregas iterativas: MVP (autenticación, CRUD básico, dashboard simple), seguida por funcionalidades avanzadas (préstamos, notificaciones, reportes automatizados).

Participación y Colaboración
-----------------------------
Los siguientes colaboradores y herramientas forman parte activa del ciclo de desarrollo y mantenimiento:

- 🤖 **Renovate** (https://github.com/apps/renovate):
  Automatiza la actualización de dependencias y mantiene el proyecto alineado con las versiones más recientes y seguras.

- 🤖 **CodeRabbit AI** (https://github.com/apps/coderabbitai):
  Asiste en la revisión automatizada de código y en la generación de sugerencias inteligentes para mejorar la calidad técnica.

- 🧠 **GitHub Copilot (Code Review Integration)**:
  Participa en la revisión de código mediante sugerencias contextuales, asegurando adherencia a los principios definidos en esta constitución.

- 🔒 **GitHub Advanced Security**:
  Supervisa vulnerabilidades, análisis de secretos y dependencias inseguras en tiempo real.

Repositorio y Control de Versiones
-----------------------------------
- **Repositorio GitHub**: https://github.com/dbacilio88/app-money-manager-go.git
- **Branch principal**: main
- **Estrategia de branching**: GitFlow (feature/* para nuevas funcionalidades, hotfix/* para correcciones urgentes)
- **Pull Requests**: Obligatorios con revisión de código antes de merge

Cambios futuros
---------------
Este documento podrá actualizarse con aprobaciones del equipo. Los cambios mayores requerirán versión nueva.
