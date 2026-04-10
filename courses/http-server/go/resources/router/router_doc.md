# Routing: 404 vs 405

Un router HTTP mapea `(method, path)` a un handler. Los dos códigos de error
que tu router debe distinguir son:

- **404 Not Found** — ningún handler está registrado para ese path bajo
  ningún method.
- **405 Method Not Allowed** — el path existe pero bajo otros methods. Por
  ejemplo, el cliente hace `POST /health` y tú solo tienes `GET /health`.

La diferencia importa porque un 404 le dice al cliente "ese recurso no
existe en este servidor" y un 405 le dice "existe, pero no con ese verbo —
revisa tu method". Los clientes REST bien escritos distinguen ambas.

Estrategia típica:

1. Primero buscas match exacto por `(method, path)`. Si existe, despachas.
2. Si no, recorres las rutas registradas; si al menos una comparte el mismo
   `path` con otro method, responde 405.
3. Si tampoco hay coincidencia por path, delega al `Fallback` (archivos
   estáticos, etc.) o responde 404.
