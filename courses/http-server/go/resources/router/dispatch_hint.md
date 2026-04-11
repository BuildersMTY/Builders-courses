# Hint: Agregar 405 y Fallback

Después del match exacto pero antes del 404 final:

## Paso 1: Detectar 405

```go
for k := range router.routes {
    if k.path == r.Path {
        WriteError(w, 405, "Method Not Allowed")
        return
    }
}
```

Si la ruta existe pero con otro método, es un 405 (no un 404).

## Paso 2: Fallback

```go
if router.Fallback != nil {
    router.Fallback.ServeHTTP(w, r)
    return
}
```

El fallback se usa para servir archivos estáticos cuando no hay ruta registrada.
