# C4 Model — Containers

Contenedores principales:
1. Web Application (Go/Gin)
   - Exposición: API REST / web static
   - Responsabilidad: lógica de negocio, autenticación, rutas, servicios
2. MongoDB
   - Persistencia de usuarios, transacciones, categorías y tokens
3. Frontend estático
   - HTML/CSS/JS servido desde Go (o CDN)
4. Nginx (opcional en entorno de producción)
   - Reverse proxy, TLS termination, caching

Diagrama (texto):
- User -> Browser -> Nginx -> Web Application
- Web Application -> MongoDB
- Web Application -> Google OAuth

Notas de despliegue:
- En docker-compose local, levantaremos `app` y `mongo` y, opcionalmente, `mongo-express` para inspección.
- En producción, usar replicaset de Mongo y AWS (ECS/EKS) para la app.
