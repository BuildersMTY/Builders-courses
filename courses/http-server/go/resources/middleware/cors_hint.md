# Hint: Implementar CORSMiddleware

```go
func CORSMiddleware(next Handler) Handler {
    return HandlerFunc(func(w ResponseWriter, r *Request) {
        w.SetHeader("Access-Control-Allow-Origin", "*")
        w.SetHeader("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.SetHeader("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(200)
            return  // Early exit — no llamar al handler
        }

        next.ServeHTTP(w, r)
    })
}
```

El preflight (OPTIONS) es una consulta del browser antes de enviar el request real. Respondemos 200 directo sin procesar la lógica del handler.
