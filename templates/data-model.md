# Modelo de Datos

## Entidades

### Users
- id (PK)
- nombre
- correo
- fecha_nacimiento
- telefono
- foto
- role_id (FK Roles)
- created_at, updated_at

### Roles
- id (PK)
- nombre (Admin, Usuario)

### Categories
- id (PK)
- nombre
- descripcion
- created_by (FK Users)

### Transactions
- id (PK)
- tipo (Ingreso/Egreso)
- monto
- descripcion
- fecha
- categoria_id (FK Categories)
- user_id (FK Users)

### Reports
- id (PK)
- user_id (FK Users)
- mes
- total_ingresos
- total_egresos
- fecha_generacion
