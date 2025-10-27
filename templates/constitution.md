# Sistema de Control Financiero Personal Constitution

## Core Principles

### I. Seguridad
Todas las contraseñas deben cifrarse utilizando algoritmos modernos (bcrypt).  
El control de acceso se implementará mediante roles (administrador, usuario estándar, auditor).  
No se permitirá el acceso directo a datos sensibles sin autenticación ni autorización adecuadas.

### II. Integridad de Datos
Toda transacción debe validarse antes de almacenarse.  
No se permiten movimientos duplicados ni montos negativos.  
Cada registro financiero tendrá un identificador único (uuid) y trazabilidad de su origen.

### III. Usabilidad
La interfaz será responsiva, moderna y visualmente clara.  
El diseño priorizará la simplicidad: los usuarios deben poder realizar operaciones financieras sin capacitación previa.  
Las funciones críticas estarán accesibles en menos de tres clics.

### IV. Transparencia
Los reportes deben ser completos, precisos y fácilmente exportables (CSV, PDF).  
Cada cálculo debe poder rastrearse a su fuente de datos.  
El sistema debe ofrecer un historial auditable de movimientos y acciones del usuario.

### V. Flexibilidad y Responsabilidad
El sistema debe poder ampliarse con nuevas categorías, monedas o funcionalidades sin romper la compatibilidad existente.  
Solo el administrador podrá eliminar registros; los usuarios son responsables de los datos que ingresen.  
Toda eliminación o modificación significativa debe registrarse en un log seguro.

## Additional Constraints

- **Cumplimiento Legal:** Adherencia a normativas locales de protección de datos personales.  
- **Tecnología Base:** Aplicación web desarrollada en Golang (backend), html (frontend) y PostgreSQL (base de datos).  
- **Disponibilidad:** Objetivo de uptime ≥ 99.5 %.  
- **Backup:** Copias automáticas diarias y retención de 30 días, desde la web del administrador.
- **Seguridad de Sesión:** Tokens JWT con expiración de 24 h, renovación mediante refresh tokens, debe mostrarse en la pantalla el tiempo de expiracion en la pantalla de la web, parte superior derecha.

## Development Workflow

- **Test-First (obligatorio):** Todo nuevo módulo debe incluir pruebas unitarias y de integración antes del merge.  
- **Code Review:** Cada PR será revisado por al menos un miembro del equipo.  
- **Branching:** Flujo `main → develop → feature/*`.  
- **CI/CD:** Integración continua automatizada con ejecución de pruebas y análisis estático de código (Github actions).  
- **Documentación:** Cada feature deberá documentarse en el changelog y en el manual técnico.  

## Collaborators & Governance

Esta constitución define los principios rectores del proyecto y **tiene prioridad sobre cualquier práctica o convención previa**.  

### Participación y Colaboración
Los siguientes colaboradores y herramientas forman parte activa del ciclo de desarrollo y mantenimiento:

- 🤖 **Renovate** ([https://github.com/apps/renovate]):  
  Automatiza la actualización de dependencias y mantiene el proyecto alineado con las versiones más recientes y seguras.  

- 🤖 **CodeRabbit AI** ([https://github.com/apps/coderabbitai]):  
  Asiste en la revisión automatizada de código y en la generación de sugerencias inteligentes para mejorar la calidad técnica.  

- 🧠 **GitHub Copilot (Code Review Integration):**  
  Participa en la revisión de código mediante sugerencias contextuales, asegurando adherencia a los principios definidos en esta constitución.  

- 🔒 **GitHub Advanced Security:**  
  Supervisa vulnerabilidades, análisis de secretos y dependencias inseguras en tiempo real.  

### Aprobación de Cambios
1. Cualquier miembro del equipo puede proponer modificaciones o mejoras a esta constitución mediante Pull Request.  
2. Toda propuesta deberá ser revisada por al menos 1 revisor humano o una combinación de revisor humano + herramienta automatizada (Copilot o CodeRabbit).  
3. Los cambios aprobados se versionarán con una nota en el changelog y actualización de las fechas de ratificación.  

**Version**: 1.0.0 | **Ratified**: 2025-10-27 | **Last Amended**: 2025-10-27