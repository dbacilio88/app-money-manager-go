# API Documentation - Control Financiero

**Base URL**: `http://localhost:8080/api/v1`

**Stack**: Go 1.25.3 + MongoDB + Gin Framework

## Autenticación

Todas las rutas protegidas requieren un header de autenticación:

```
Authorization: Bearer <access_token>
```

Los tokens JWT tienen una duración de 15 minutos. Usa el refresh token para obtener nuevos access tokens.

---

## Endpoints Públicos

### 1. Registro de Usuario

**POST** `/auth/register`

Registra un nuevo usuario en el sistema. El usuario se crea con estado `pending` y debe ser aprobado por un administrador.

**Request Body**:
```json
{
  "nombre": "Juan Pérez",
  "email": "juan@example.com",
  "password": "password123"
}
```

**Response** (201 Created):
```json
{
  "success": true,
  "message": "Usuario registrado. Espera aprobación del administrador",
  "data": {
    "usuario": {
      "id": "67890abcdef1234567890abc",
      "nombre": "Juan Pérez",
      "email": "juan@example.com",
      "rol": "user",
      "estado": "pending"
    }
  }
}
```

---

### 2. Login

**POST** `/auth/login`

Inicia sesión con email y contraseña.

**Request Body**:
```json
{
  "email": "juan@example.com",
  "password": "password123"
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Login exitoso",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "usuario": {
      "id": "67890abcdef1234567890abc",
      "nombre": "Juan Pérez",
      "email": "juan@example.com",
      "rol": "user",
      "estado": "active"
    }
  }
}
```

**Errores**:
- `401`: Usuario o contraseña incorrectos
- `403`: Usuario pendiente de aprobación o suspendido

---

### 3. Google OAuth - Obtener URL

**GET** `/auth/google`

Obtiene la URL para iniciar el flujo de autenticación con Google.

**Response** (200 OK):
```json
{
  "success": true,
  "data": {
    "url": "https://accounts.google.com/o/oauth2/v2/auth?..."
  }
}
```

---

### 4. Google OAuth - Callback

**GET** `/auth/google/callback?code=<authorization_code>`

Callback de Google OAuth. Redirige al frontend con los tokens.

**Query Parameters**:
- `code`: Authorization code de Google

**Response**: Redirección a `APP_URL/?access_token=...&refresh_token=...`

---

### 5. Refresh Token

**POST** `/auth/refresh`

Obtiene un nuevo access token usando un refresh token válido.

**Request Body**:
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

---

## Endpoints Protegidos (Requieren Autenticación)

### 6. Obtener Perfil

**GET** `/perfil`

Obtiene la información del usuario autenticado.

**Response** (200 OK):
```json
{
  "success": true,
  "data": {
    "id": "67890abcdef1234567890abc",
    "nombre": "Juan Pérez",
    "email": "juan@example.com",
    "foto": "https://...",
    "rol": "user",
    "estado": "active",
    "createdAt": "2025-10-27T10:00:00Z"
  }
}
```

---

### 7. Actualizar Perfil

**PUT** `/perfil`

Actualiza la información del perfil del usuario.

**Request Body**:
```json
{
  "nombre": "Juan Carlos Pérez",
  "foto": "https://..."
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Perfil actualizado",
  "data": {
    "id": "67890abcdef1234567890abc",
    "nombre": "Juan Carlos Pérez",
    "email": "juan@example.com",
    "foto": "https://...",
    "rol": "user"
  }
}
```

---

### 8. Cambiar Contraseña

**POST** `/auth/change-password`

Cambia la contraseña del usuario autenticado.

**Request Body**:
```json
{
  "password_actual": "password123",
  "password_nueva": "newpassword456"
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Contraseña actualizada exitosamente"
}
```

---

## Categorías

### 9. Listar Categorías

**GET** `/categorias`

Lista todas las categorías del usuario (propias + globales).

**Response** (200 OK):
```json
{
  "success": true,
  "data": [
    {
      "id": "67890abcdef1234567890abc",
      "nombre": "Salario",
      "tipo": "ingreso",
      "color": "#4CAF50",
      "usuarioId": null,
      "createdAt": "2025-10-27T10:00:00Z"
    }
  ]
}
```

---

### 10. Crear Categoría

**POST** `/categorias`

Crea una nueva categoría personal.

**Request Body**:
```json
{
  "nombre": "Freelance",
  "tipo": "ingreso",
  "color": "#2196F3"
}
```

**Response** (201 Created):
```json
{
  "success": true,
  "message": "Categoría creada",
  "data": {
    "id": "67890abcdef1234567890abc",
    "nombre": "Freelance",
    "tipo": "ingreso",
    "color": "#2196F3",
    "usuarioId": "67890abcdef1234567890abc"
  }
}
```

---

### 11. Actualizar Categoría

**PUT** `/categorias/:id`

Actualiza una categoría propia (no se pueden modificar las globales).

**Request Body**:
```json
{
  "nombre": "Proyectos Freelance",
  "color": "#1976D2"
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Categoría actualizada",
  "data": {
    "id": "67890abcdef1234567890abc",
    "nombre": "Proyectos Freelance",
    "tipo": "ingreso",
    "color": "#1976D2"
  }
}
```

---

### 12. Eliminar Categoría

**DELETE** `/categorias/:id`

Elimina una categoría propia.

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Categoría eliminada"
}
```

---

## Transacciones

### 13. Listar Transacciones

**GET** `/transacciones?fecha_inicio=2025-10-01&fecha_fin=2025-10-31&tipo=ingreso`

Lista las transacciones del usuario autenticado con filtros opcionales.

**Query Parameters**:
- `fecha_inicio` (opcional): Fecha desde (formato: YYYY-MM-DD)
- `fecha_fin` (opcional): Fecha hasta (formato: YYYY-MM-DD)
- `tipo` (opcional): Tipo de transacción (ingreso, egreso, prestamo, alquiler, otro)
- `categoria_id` (opcional): ID de categoría

**Response** (200 OK):
```json
{
  "success": true,
  "data": [
    {
      "id": "67890abcdef1234567890abc",
      "tipo": "ingreso",
      "categoriaId": "67890abcdef1234567890abc",
      "categoria": {
        "nombre": "Salario",
        "color": "#4CAF50"
      },
      "monto": 3000.00,
      "moneda": "USD",
      "fecha": "2025-10-25T10:00:00Z",
      "descripcion": "Salario mensual",
      "createdAt": "2025-10-25T10:00:00Z"
    }
  ]
}
```

---

### 14. Crear Transacción

**POST** `/transacciones`

Crea una nueva transacción.

**Request Body**:
```json
{
  "tipo": "ingreso",
  "categoriaId": "67890abcdef1234567890abc",
  "monto": 3000.00,
  "moneda": "USD",
  "fecha": "2025-10-25T10:00:00Z",
  "descripcion": "Salario mensual"
}
```

**Response** (201 Created):
```json
{
  "success": true,
  "message": "Transacción creada",
  "data": {
    "id": "67890abcdef1234567890abc",
    "tipo": "ingreso",
    "categoriaId": "67890abcdef1234567890abc",
    "monto": 3000.00,
    "moneda": "USD",
    "fecha": "2025-10-25T10:00:00Z",
    "descripcion": "Salario mensual"
  }
}
```

---

### 15. Actualizar Transacción

**PUT** `/transacciones/:id`

Actualiza una transacción existente.

**Request Body**:
```json
{
  "monto": 3200.00,
  "descripcion": "Salario mensual + bono"
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Transacción actualizada",
  "data": {
    "id": "67890abcdef1234567890abc",
    "monto": 3200.00,
    "descripcion": "Salario mensual + bono"
  }
}
```

---

### 16. Eliminar Transacción

**DELETE** `/transacciones/:id`

Elimina una transacción.

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Transacción eliminada"
}
```

---

## Reportes

### 17. Estadísticas Generales

**GET** `/reportes/estadisticas?mes=10&anio=2025`

Obtiene estadísticas de ingresos, egresos y balance.

**Query Parameters**:
- `mes` (opcional): Mes (1-12)
- `anio` (opcional): Año (YYYY)

**Response** (200 OK):
```json
{
  "success": true,
  "data": {
    "totalIngresos": 5000.00,
    "totalEgresos": 3200.00,
    "balance": 1800.00,
    "transacciones": 25,
    "porCategoria": [
      {
        "categoria": "Salario",
        "tipo": "ingreso",
        "total": 3000.00
      },
      {
        "categoria": "Alimentación",
        "tipo": "egreso",
        "total": 800.00
      }
    ]
  }
}
```

---

### 18. Balance Actual

**GET** `/reportes/balance`

Obtiene el balance actual total del usuario.

**Response** (200 OK):
```json
{
  "success": true,
  "data": {
    "balance": 12500.00,
    "moneda": "USD"
  }
}
```

---

## Endpoints de Administrador (Requieren rol: admin)

### 19. Listar Todos los Usuarios

**GET** `/admin/usuarios`

Lista todos los usuarios del sistema.

**Response** (200 OK):
```json
{
  "success": true,
  "data": [
    {
      "id": "67890abcdef1234567890abc",
      "nombre": "Juan Pérez",
      "email": "juan@example.com",
      "rol": "user",
      "estado": "active",
      "createdAt": "2025-10-27T10:00:00Z"
    }
  ]
}
```

---

### 20. Aprobar Usuario

**POST** `/admin/usuarios/:id/aprobar`

Aprueba un usuario pendiente (cambia estado de `pending` a `active`).

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Usuario aprobado"
}
```

---

### 21. Activar Usuario

**POST** `/admin/usuarios/:id/activar`

Activa un usuario suspendido.

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Usuario activado"
}
```

---

### 22. Desactivar Usuario

**POST** `/admin/usuarios/:id/desactivar`

Suspende un usuario activo.

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Usuario desactivado"
}
```

---

### 23. Cambiar Rol de Usuario

**POST** `/admin/usuarios/:id/cambiar-rol`

Cambia el rol de un usuario (user ↔ admin).

**Request Body**:
```json
{
  "rol": "admin"
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Rol actualizado"
}
```

---

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200    | OK - Operación exitosa |
| 201    | Created - Recurso creado |
| 400    | Bad Request - Datos inválidos |
| 401    | Unauthorized - Token inválido o expirado |
| 403    | Forbidden - Sin permisos |
| 404    | Not Found - Recurso no encontrado |
| 409    | Conflict - Conflicto (ej: email duplicado) |
| 500    | Internal Server Error - Error del servidor |

---

## Formato de Respuesta Estándar

### Éxito
```json
{
  "success": true,
  "message": "Mensaje descriptivo",
  "data": { /* datos */ }
}
```

### Error
```json
{
  "success": false,
  "error": "Mensaje de error descriptivo"
}
```

---

## Notas Adicionales

- Todos los timestamps están en formato ISO 8601 (UTC)
- Los montos son números decimales con 2 decimales de precisión
- Los IDs son ObjectIDs de MongoDB (24 caracteres hexadecimales)
- Las fechas en query parameters usan formato YYYY-MM-DD
- Las respuestas siempre son JSON con Content-Type: application/json
