# C4 Model — Despliegue

Entorno local (docker-compose):
- app: imagen construida desde Dockerfile (Go binary)
- mongo: imagen oficial mongodb (volumen para persistencia)
- mongo-express (opcional) para inspección

Entorno producción (alto nivel):
- Contenedor app ejecutando binario Go en ECS / EKS o EC2 Auto Scaling
- MongoDB en clúster administrado (MongoDB Atlas o replica set en EC2)
- Load balancer (ALB), certificado TLS (ACM), CloudFront opcional
- Secret Manager para claves (Google client secret, JWT secret)

Consideraciones de red:
- Backend stateless; sesiones manejadas por JWT.
- Mínimo: exponer puerto 443 y no exponer MongoDB al Internet público.

Backup y recuperación:
- Snapshots regulares de MongoDB y pruebas de restauración periódicas.
