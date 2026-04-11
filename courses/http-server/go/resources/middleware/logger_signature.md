# Signature: LoggerMiddleware

```go
func LoggerMiddleware(next Handler) Handler
```

### loggingResponseWriter

```go
type loggingResponseWriter struct {
    ResponseWriter          // embebido
    statusCode     int      // captura el status code
}
```

Implementar `WriteHeader` en `loggingResponseWriter`: guardar `w.statusCode = code` y delegar `w.ResponseWriter.WriteHeader(code)`.

### LoggerMiddleware

1. Registrar `start := time.Now()`
2. Crear `lw := &loggingResponseWriter{ResponseWriter: w, statusCode: 0}`
3. Llamar `next.ServeHTTP(lw, r)` — usar `lw` en vez de `w`
4. Calcular `duration := time.Since(start)`
5. Loggear: `log.Printf("%s %s %d %v", r.Method, r.Path, lw.statusCode, duration)`
