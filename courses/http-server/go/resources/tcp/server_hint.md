# Hint: Server.Start paso a paso

1. Abre el listener con `net.Listen("tcp", s.Addr)`. Si falla, retorna el error
   envuelto con `fmt.Errorf`.
2. `defer listener.Close()`.
3. Dentro de `for {}`, llama a `listener.Accept()`. Ante error, `log.Printf`
   y `continue`.
4. Dispara `go s.handleConnection(conn)` por cada conexión aceptada.

Para `handleConnection`:

1. `defer conn.Close()` al principio.
2. Envuelve la conexión en `bufio.NewReader(conn)`.
3. Llama a `ParseRequest(reader)`. Si retorna error, crea el
   `ResponseWriter` con `NewResponseWriter(conn)`, llama a
   `WriteError(w, 400, "Bad Request")` y sal.
4. Si el parseo fue exitoso, crea el `ResponseWriter` y llama a
   `s.Handler.ServeHTTP(w, req)`.

No necesitas timeouts ni deadlines en este módulo — son optimizaciones para
más adelante.
