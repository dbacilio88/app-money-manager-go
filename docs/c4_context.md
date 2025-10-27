# C4 Model — Context

Resumen: visión de alto nivel de los sistemas y actores.

Actores:
- Usuario final (web)
- Administrador
- Sistema externo: Google (OAuth)

Sistema: Control Financiero (aplicación web)
- Interactúa con usuarios vía navegador.
- Provee API REST para operaciones y UI estática para frontend.
- Se integra con Google OAuth para autenticación social.
- Persiste datos en MongoDB.

Relaciones (alto nivel):
- Usuario → Navegador → Sistema (HTTP/HTTPS)
- Sistema → MongoDB (driver oficial)
- Sistema ↔ Google OAuth (autorización)

Notas:
- El sistema será desplegado en contenedores Docker; en producción se recomienda behind a load balancer con TLS.
