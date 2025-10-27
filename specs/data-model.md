# Modelo de Datos — Control Financiero

Versión: 0.1

Nota: Este documento describe el modelo de datos pensado para **MongoDB**. Se usa una combinación de colecciones normalizadas (usuarios, roles, categorias) y documentos embebidos cuando resulta práctico (movimientos asociados a cuentas específicas).

**Tecnologías**: Backend en **Go 1.25.3** con base de datos **MongoDB**.

Colecciones principales
----------------------
1. usuarios
```json
{
  "_id": ObjectId,
  "nombre": "string",
  "email": "string",
  "passwordHash": "string", // bcrypt
  "foto": "string",
  "googleId": "string | null",
  "rol": "string", // 'user' | 'admin'
  "estado": "string", // 'pending' | 'active' | 'suspended'
  "createdAt": ISODate,
  "updatedAt": ISODate
}
```

2. categorias
```json
{
  "_id": ObjectId,
  "usuarioId": ObjectId, // usuario propietario (null si global)
  "nombre": "string",
  "tipo": "string", // 'ingreso' | 'egreso' | 'ambos'
  "color": "string",
  "createdAt": ISODate,
  "updatedAt": ISODate
}
```

3. transacciones
```json
{
  "_id": ObjectId,
  "usuarioId": ObjectId,
  "tipo": "string", // 'ingreso' | 'egreso' | 'prestamo' | 'alquiler' | 'otro'
  "categoriaId": ObjectId,
  "monto": NumberDecimal,
  "moneda": "string",
  "fecha": ISODate,
  "descripcion": "string",
  "cuenta": {
    "id": ObjectId,
    "nombre": "string"
  },
  "metodoPago": "string",
  "tags": ["string"],
  "referencia": "string",
  "createdAt": ISODate,
  "updatedAt": ISODate
}
```

4. refresh_tokens
```json
{
  "_id": ObjectId,
  "usuarioId": ObjectId,
  "token": "string",
  "expiresAt": ISODate,
  "revoked": boolean,
  "createdAt": ISODate
}
```

5. roles (simple)
```json
{
  "_id": ObjectId,
  "nombre": "string",
  "permisos": ["string"]
}
```

6. logs / audit
```json
{
  "_id": ObjectId,
  "usuarioId": ObjectId | null,
  "accion": "string",
  "detalle": "object",
  "ip": "string",
  "userAgent": "string",
  "createdAt": ISODate
}
```

Índices sugeridos
-----------------
- usuarios: unique(email)
- transacciones: {usuarioId:1, fecha:-1}
- categorias: {usuarioId:1, nombre:1}
- refresh_tokens: {token:1}

Notas de diseño
----------------
- Se recomienda usar ObjectId para relaciones y usar `$lookup` cuando se requiera información combinada para reportes.
- Para reportes de alta performance, se evaluará creación de colecciones agregadas o pipelines periódicos (ETL) que precomputen métricas por mes/categoría.

Repositorio y Versionado
-------------------------
- **Repositorio GitHub**: https://github.com/dbacilio88/app-money-manager-go.git
- **Migraciones**: Las migraciones de esquema se gestionan mediante scripts Go en `internal/config/database.go`
- **Seed Data**: Datos iniciales (usuarios admin, categorías globales) se crean automáticamente al iniciar la aplicación

Calidad y Mantenimiento
------------------------
El modelo de datos es monitoreado y mantenido con:
- 🤖 **Renovate**: Actualiza dependencias del driver MongoDB
- 🤖 **CodeRabbit AI**: Revisa cambios en modelos durante PRs
- 🧠 **GitHub Copilot**: Asiste en la definición de nuevos modelos
- 🔒 **GitHub Advanced Security**: Escanea vulnerabilidades en queries y datos sensibles
