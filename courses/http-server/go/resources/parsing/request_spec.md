# HTTP/1.1 Request Format (extracto RFC 7230)

Una request HTTP/1.1 tiene tres partes separadas por CRLF (`\r\n`):

```
<method> SP <request-target> SP HTTP/1.1 CRLF
<Header-Name>: <value> CRLF
<Header-Name>: <value> CRLF
...
CRLF
<body bytes>
```

Ejemplo real de `curl -X POST -d "hola" http://localhost:8080/echo`:

```
POST /echo HTTP/1.1
Host: localhost:8080
User-Agent: curl/8.4.0
Accept: */*
Content-Length: 4
Content-Type: application/x-www-form-urlencoded

hola
```

Reglas que te afectan:

- El separador entre línea es CRLF (`\r\n`), no solo `\n`.
- El fin de la sección de headers es una **línea vacía** (CRLF solo).
- El body **no** lleva terminador. El servidor sabe cuánto leer únicamente
  si el cliente envió `Content-Length: N` — en ese caso, lee exactamente N
  bytes. Si no hay `Content-Length`, no hay body.
- Los nombres de header son case-insensitive en la spec, pero en este curso
  los guardamos tal cual (case-sensitive).
