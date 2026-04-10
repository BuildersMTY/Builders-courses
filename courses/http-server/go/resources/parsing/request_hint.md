# Hint: ParseRequest paso a paso

1. Crea `req := &Request{Headers: make(map[string]string)}`.
2. Lee la request line con `reader.ReadString('\n')`, recórtala con
   `strings.TrimRight(line, "\r\n")`.
3. Split por espacio: `parts := strings.Split(line, " ")`. Debe tener
   exactamente 3 elementos, si no, retorna error.
4. Asigna `req.Method`, `req.Path`, `req.Version`.
5. Loop de headers: lee línea, recórtala, si es `""` rompe el loop.
6. Para cada header, `strings.SplitN(line, ":", 2)`. Si len != 2, `continue`
   (header malformado, no revientes).
7. Guarda `Headers[TrimSpace(key)] = TrimSpace(val)`.
8. Mira si existe `req.Headers["Content-Length"]`. Si sí y > 0, reserva
   `req.Body = make([]byte, cl)` y llama a `io.ReadFull(reader, req.Body)`.
9. Retorna `req, nil`.

Detalle crítico: **nunca** intentes leer el body sin `Content-Length`.
Tu `io.ReadFull` quedará colgado esperando bytes que no llegarán.
