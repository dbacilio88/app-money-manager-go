# Sistema de Control Financiero Personal

## Propósito
El propósito de este sistema es ofrecer una herramienta moderna y confiable que permita al usuario gestionar sus finanzas personales de forma detallada y visual, mediante formularios.  
El sistema registrará ingresos, egresos y categorías, generando automáticamente reportes mensuales con gráficos interactivos, ofreciendo una visión clara del estado financiero.

## Objetivos
- Registrar categorias.
- Registrar ingresos y egresos por categoría.
- Definir un monto inicial como base del balance (el monto debe ingresarse por formulario).
- Emitir un reporte mensual con resumen de ingresos, egresos y tres gráficas.
- Permitir inicio de sesión seguro con roles de **Administrador** y **Usuario**, debe usar la api de google para el inicio de sesión.
- Proveer una interfaz web moderna, responsiva y agradable, con soporte para modo claro/oscuro.

## Actores
| Actor | Descripción | Permisos |
|--------|--------------|-----------|
| **Administrador** | Control total del sistema. | Crear, editar y eliminar registros, gestionar usuarios, ver reportes, generar backups mensuales, cambiar el rol del usuario, visualizar perfiles, listar usuarios. |
| **Usuario** | Usuario regular. | Registrar ingresos y egresos, ver reportes, visualizar su perfil,actualizar perfil (sin eliminar registros). |

## Principales flujos de usuario

### Flujo 1: Autenticación
1. El usuario ingresa correo y contraseña.
2. El sistema valida credenciales y asigna rol.
3. Redirige al panel principal.

### Flujo 2: Registro de movimientos
1. El usuario selecciona “Registrar ingreso” o “Registrar egreso”.
2. Completa monto, categoría, descripción y fecha.
3. El sistema guarda y actualiza balance.

### Flujo 3: Reporte mensual
1. Al cierre de mes, el sistema genera reporte con totales y gráficas:
   - Distribución de gastos por categoría.
   - Evolución de ingresos/egresos.
   - Balance acumulado.
2. El usuario o administrador visualiza el reporte en “Reportes”.

### Flujo 4: Gestión de usuarios (Administrador)
1. El administrador crea,lista,edita o elimina usuarios.
2. Asigna roles y restablece contraseñas.
3. Visualiza perfiles de usuario.

### Flujo 5: Personalización de tema
1. El usuario selecciona tema claro, oscuro o automático.
2. Preferencia guardada en el perfil o sesión.
3. Visualizar su perfil de usuario.

## Reglas del sistema
1. Los usuarios no pueden eliminar registros financieros.
2. Cada usuario solo visualiza sus propios datos.
3. Los reportes se generan automáticamente al cierre de mes.
4. La base de datos mantiene integridad referencial.
5. El diseño visual sigue paleta pastel y modo oscuro/claro.
6. El Sistema debera usar los recursos que son proporcionados en la carpeta `./assets`.
7. Los recursos proporcionados como videos, fuentes, imagenes, deberan ser usandos en el sistema.
8. De requerir imagenes como `.ico`, debera convertir al formato deseado.
9. El login al sistema debe ser por el api de google.