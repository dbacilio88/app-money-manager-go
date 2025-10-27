# C4 Model — Componentes (Web Application)

Componentes principales dentro del contenedor Web Application (Go/Gin):

1. Router & Handlers (internal/routes)
   - Ruteo de endpoints /api/v1
   - Middleware para logging, CORS y autenticación
2. Controllers
   - AuthController: login, register, callback, refresh
   - UsuarioController: CRUD usuarios (admin aprovisionamiento)
   - TransaccionController: CRUD transacciones
   - CategoriaController: CRUD categorias
   - ReporteController: endpoints para métricas
3. Services
   - UsuarioService, TransaccionService, CategoriaService, AuthService
   - Contienen la lógica de negocio y validaciones
4. Repositories / Data Access
   - Repositorios que encapsulan operaciones con MongoDB (colecciones y queries)
5. Auth Layer
   - JWT manager (generar/verificar tokens)
   - Refresh token manager (guardar/revocar)
6. Static Assets
   - HTML, CSS, JS y assets servidos por el servidor

Interacción entre componentes (texto):
- Handlers -> Controllers -> Services -> Repositories -> MongoDB
- Auth middleware intercepta requests y valida JWTs antes de Controller

Notas técnicas:
- Separación clara de responsabilidades para facilitar testing.
- Services deben ser agnósticos a HTTP para permitir tests unitarios.
