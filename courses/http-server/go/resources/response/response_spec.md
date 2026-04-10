# HTTP/1.1 Response Format

Una response HTTP/1.1 se escribe byte por byte en el mismo socket que recibió
la request:

```
HTTP/1.1 <status-code> <reason-phrase> CRLF
<Header-Name>: <value> CRLF
<Header-Name>: <value> CRLF
...
CRLF
<body bytes>
```

Ejemplo:

```
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 15

{"status":"ok"}
```

Reglas importantes:

- El primer byte que escribas **debe** ser la status line. Si llamas a
  `conn.Write(body)` antes del header, rompes el protocolo y el cliente lo
  interpretará como status line malformada.
- Los headers deben terminarse con una **línea en blanco** (otro CRLF). Sin
  esa línea el cliente queda esperando más headers para siempre.
- `Content-Length` debe coincidir exactamente con el número de bytes del
  body. Si mientes, el cliente cortará el body o quedará colgado.
- Status codes mínimos que vamos a usar: 200, 400, 403, 404, 405, 500.
- Llamar `WriteHeader` dos veces es un bug — la segunda llamada debería ser
  un no-op (guarded por un flag `wroteHeader`).
