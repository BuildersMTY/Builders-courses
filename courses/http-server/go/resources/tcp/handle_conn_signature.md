# Signature: handleConnection

## Método a implementar

```go
func (s *Server) handleConnection(conn net.Conn)
```

### Comportamiento esperado

1. `defer conn.Close()` — siempre cerrar la conexión al terminar
2. Crear un `bufio.NewReader(conn)` para leer de la conexión
3. Llamar `ParseRequest(reader)` para parsear el request HTTP
4. Si ParseRequest retorna error, crear un `NewResponseWriter(conn)` y responder con `WriteError(w, 400, "Bad Request")`, luego `return`
5. Si no hay error, crear un `NewResponseWriter(conn)`
6. Si `s.Handler != nil`, delegar con `s.Handler.ServeHTTP(w, req)`
7. Si no hay Handler, responder con `WriteError(w, 404, "Not Found")`

### Paquetes necesarios

- `bufio` — `bufio.NewReader`
- `net` — `net.Conn`
- `log` — `log.Printf` (para loggear errores)
