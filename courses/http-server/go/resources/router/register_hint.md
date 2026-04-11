# Hint: Implementar Handle y dispatch exacto

## Handle — una línea

```go
func (router *Router) Handle(method, path string, handler HandlerFunc) {
    router.routes[routeKey{method: method, path: path}] = handler
}
```

## ServeHTTP — lookup en el mapa

```go
func (router *Router) ServeHTTP(w ResponseWriter, r *Request) {
    key := routeKey{method: r.Method, path: r.Path}
    if handler, ok := router.routes[key]; ok {
        handler.ServeHTTP(w, r)
        return
    }
    WriteError(w, 404, "Not Found")
}
```
