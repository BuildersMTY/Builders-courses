# Signature: CORSMiddleware

```go
func CORSMiddleware(next Handler) Handler
```

### Comportamiento esperado

1. Setear headers CORS en CADA response:
   - `Access-Control-Allow-Origin: *`
   - `Access-Control-Allow-Methods: GET, POST, OPTIONS`
   - `Access-Control-Allow-Headers: Content-Type`
2. Si el método es `OPTIONS` (preflight):
   - Responder `w.WriteHeader(200)` inmediatamente
   - `return` — NO llamar a `next.ServeHTTP`
3. Para otros métodos: llamar `next.ServeHTTP(w, r)` normalmente
