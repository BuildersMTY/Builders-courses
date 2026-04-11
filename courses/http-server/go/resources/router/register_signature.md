# Signature: Handle y dispatch exacto

## Métodos a implementar

```go
func (router *Router) Handle(method, path string, handler HandlerFunc)
func (router *Router) ServeHTTP(w ResponseWriter, r *Request)  // primera parte
```

### Handle
Almacena el handler en el mapa de rutas usando `routeKey{method, path}` como clave.

### ServeHTTP (primera parte)
1. Construye `routeKey{method: r.Method, path: r.Path}`
2. Busca en `router.routes` — si hay match exacto, llama `handler.ServeHTTP(w, r)` y retorna
3. Si no hay match, responde 404 con `WriteError(w, 404, "Not Found")`
