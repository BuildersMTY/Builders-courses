# Hint: Implementar ApplyMiddlewares

```go
func ApplyMiddlewares(base Handler, middlewares ...Middleware) Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        base = middlewares[i](base)
    }
    return base
}
```

Iteramos en reversa para que `middlewares[0]` sea el más externo. Si tenemos `[Logger, CORS]`, el flujo será: Logger → CORS → base → CORS → Logger.
