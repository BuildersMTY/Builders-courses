# Hint: Implementar LoggerMiddleware

## Paso 1: loggingResponseWriter.WriteHeader

```go
func (w *loggingResponseWriter) WriteHeader(code int) {
    w.statusCode = code
    w.ResponseWriter.WriteHeader(code)
}
```

## Paso 2: LoggerMiddleware

```go
func LoggerMiddleware(next Handler) Handler {
    return HandlerFunc(func(w ResponseWriter, r *Request) {
        start := time.Now()

        lw := &loggingResponseWriter{
            ResponseWriter: w,
            statusCode:     0,
        }

        next.ServeHTTP(lw, r)

        duration := time.Since(start)
        log.Printf("%s %s %d %v", r.Method, r.Path, lw.statusCode, duration)
    })
}
```

Envolvemos el `ResponseWriter` original con nuestro `loggingResponseWriter` para interceptar el status code sin modificar el flujo.
