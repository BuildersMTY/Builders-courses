# Solución de Referencia: Servidor HTTP/1.1 (Go Stdlib)

Este proyecto contiene la implementación de referencia completa y desde cero del servidor HTTP/1.1, utilizando únicamente sockets TCP a través del paquete `net`. Se prioriza la claridad, modularidad, y uso de la Standard Library de Go.

Para ejecutar el servidor integrado final en este nivel, usar:

```bash
go run .
```

---

## Módulos y Análisis de Edge Cases Manejados

El objetivo didáctico no es hacer un reemplazo de producción de `net/http`, sino implementar de forma correcta los fundamentos del estándar de forma concisa. Cada componente interactúa en cadena y maneja _edge cases_ estandarizados:

### Módulo 1: TCP Foundation (`server.go`)

- **Aceptación concurrente:** Bloquear un hilo en cada `Accept()` rompería el procesamiento. Se soluciona generando una _goroutine_ despachada inmediatamente por cada conexión.
- **Edge Case (Fallo de Aceptación TCP):** Si una sola conexión falla en `listener.Accept()`, no provocamos el pánico o `log.Fatal` matando servidor entero. Retornamos un log y aplicamos un `continue` aislando la falla para mantener la resiliencia en clientes restantes sin colgar el main-thread.

### Módulo 2: HTTP Parsing (`request.go`)

- **Edge Case (Headers sin Body / Ausencia de Content-Length):** El Parseo asume extraer bytes del body de red _exclusivamente_ si el parser recibe el valor de `Content-Length`. Como mecanismo evita estamparse y quedarse bloqueado perennemente tratando de encontrar más bits a la espera, limitándose a `io.ReadFull(reader, bytesLim)`.
- **Edge Case (Formato Header Roto):** Ocurrirá seguido manual testing con `netcat`. Si viene un string en CRLF que no respete la inyección de llaves estilo `Key: Value` separados por colon temporal (mal formato del cliente), optamos por atraparlo, aplicar `continue` en el bucle ignorándolo de la lectura en vez de paniquear la rutina de buffer actual permitiendo que avance limpiamente.

### Módulo 3: HTTP Response (`response.go`)

- **Edge Case (Auto-inyección Content-Length & Estatus 200 Automático):** Con el objetivo de simplificar y ser _user-friendly_ como APIs populares (Express, Gin), si el developer o el estudiante llama directamente a `w.Write([]byte("contenido"))` y olvidó ejecutar _`w.WriteHeader(200)`_, capturamos la llamada inicial para inyectar status 200 y precargamos `w.SetHeader("Content-Length", len(body))`.
- **Edge Case (Dual WriteHeader):** Si imprudencias hacen que `WriteHeader` se escriba de nuevo, rompemos el protocolo TCP reenviando bytes. Empleamos el tag bandera de memoria booleana interna `wroteHeader` neutralizando la segunda llamada para garantizar la integridad serial del socket.

### Módulo 4: Router (`router.go`)

- **Edge Case (405 Method Not Allowed):** Un cliente solicita a una ruta específica como `POST /health` que solo admite `GET /health`. De cara a Rest API correctas, la tabla de hash iterará y si se ubica al menos otro endpoint ajeno con el mismo _Path_, abortará el _404_ estándar preyectando de forma preferente código _405 Method Not Allowed_.

### Módulo 5: Static Files (`static.go`)

- **Edge Case (Mime Type Headers):** Leer la extensión y traducir `strings.ToLower(filepath.Ext())` impidiendo que Chrome malinterprete el renderizado un JS o CSS pasándolo neutro y bloqueando render engines modernos.
- **Edge Case (Index Fallback File):** Llamadas a directorios en el root arrojarían error o requerirían enumeración de path. Si la evaluación `info.IsDir()` devuelve `true`, acoplamos silenciosamente en ruta el `index.html` bajo el capó sin emitir redirects.
- **Edge Case Crítico (Seguridad Path Traversal):** Un intruso en Curl solicita subidas por `../../../../etc/passwd` tratando de leer Linux files internos. Unificamos la ruta base impuesta en el servidor con el parámetro entrante después de una sanitización de strings `filepath.Clean()`. Como etapa paranoica, comparamos el AbsPath resultante forzándolo a que deba empezar y contener un substring originado estrictamente dentro del Path referenciado (Directory Jailing). Con eso la barrera devuelve `403 Forbidden` al primer intento.

### Módulo 6: Middleware (`handler.go`)

- **Edge Case (Interceptación en Decorator Chain):** Para métricas y logs post-ejecución, `LoggerMiddleware` requiere conocer el código status resuelto. Hacemos bypass e inyectamos la mini estructura abstracta nativa `loggingResponseWriter` mediante técnica _Monkey Patch_ incrustada dentro de cada Request. Permitiendo sobre-escribir el flag de respuesta y derivarlo a `log.Printf` en una sola impresión de stdout limpia sin ensuciar implementaciones base.
- **Edge Case (CORS Options Preflight):** Al detectar `OPTIONS` abortamos en capa early-exit (`return` de inmediato inyectando HTTP status 200) y frenamos su descenso al controlador en el `router.go`, evitando malgastar procesamiento en bases de datos simuladas ante latencias simples.
