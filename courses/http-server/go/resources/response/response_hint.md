# Hint: ResponseWriter paso a paso

**SetHeader** es trivial: `r.headers[key] = value`.

**WriteHeader(code)**:

1. Si `r.wroteHeader` es `true`, retorna inmediato.
2. Marca `r.wroteHeader = true` y guarda `r.statusCode = code`.
3. Escribe la status line:
   `r.conn.Write([]byte(fmt.Sprintf("HTTP/1.1 %d %s\r\n", code, StatusText(code))))`
4. Itera `r.headers` y escribe cada una como `"Key: Value\r\n"`.
5. Escribe la línea vacía terminadora: `r.conn.Write([]byte("\r\n"))`.

**Write(b)**:

1. Si `!r.wroteHeader`, setea `Content-Length` con `fmt.Sprintf("%d", len(b))`
   y llama a `r.WriteHeader(200)` **antes** de escribir los bytes.
2. `return r.conn.Write(b)`.

El orden importa: si haces `Write` del body antes de emitir la línea de
status + headers, el cliente recibirá basura.
