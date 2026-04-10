# Hint: Logger, CORS y ApplyMiddlewares

**loggingResponseWriter** usa _embedding_ de interfaz para heredar gratis
`Write` y `SetHeader` del writer original, y solo redefine `WriteHeader`:

```go
type loggingResponseWriter struct {
    ResponseWriter
    statusCode int
}

func (w *loggingResponseWriter) WriteHeader(code int) {
    w.statusCode = code
    w.ResponseWriter.WriteHeader(code)
}
```

**LoggerMiddleware**:

```go
return HandlerFunc(func(w ResponseWriter, r *Request) {
    start := time.Now()
    lw := &loggingResponseWriter{ResponseWriter: w}
    next.ServeHTTP(lw, r)
    log.Printf("%s %s %d %v", r.Method, r.Path, lw.statusCode, time.Since(start))
})
```

**CORSMiddleware**:

```go
return HandlerFunc(func(w ResponseWriter, r *Request) {
    w.SetHeader("Access-Control-Allow-Origin",  "*")
    w.SetHeader("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.SetHeader("Access-Control-Allow-Headers", "Content-Type")
    if r.Method == "OPTIONS" {
        w.WriteHeader(200)
        return // no descender al handler
    }
    next.ServeHTTP(w, r)
})
```

**ApplyMiddlewares** encadena al revés para que el primer middleware sea el más
externo:

```go
for i := len(middlewares) - 1; i >= 0; i-- {
    base = middlewares[i](base)
}
return base
```
