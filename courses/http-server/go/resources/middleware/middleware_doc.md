# Middlewares como Decorator Chain

Un _middleware_ envuelve un `Handler` para agregarle comportamiento sin
modificar el handler original. En Go se expresa como:

```go
type Middleware func(Handler) Handler
```

Un middleware típico:

```go
func LoggerMiddleware(next Handler) Handler {
    return HandlerFunc(func(w ResponseWriter, r *Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.Path, time.Since(start))
    })
}
```

## Encadenamiento

Si aplicas varios middlewares:

```go
chain := ApplyMiddlewares(router, LoggerMiddleware, CORSMiddleware)
```

El orden importa: el **primero** de la lista se ejecuta primero al entrar y
último al salir (es el más externo). Para lograrlo, iteras la lista al revés
y envuelves progresivamente:

```go
for i := len(mws) - 1; i >= 0; i-- {
    base = mws[i](base)
}
return base
```

## CORS preflight

Un browser, antes de enviar una request "no simple" (POST con JSON, por
ejemplo), envía primero un `OPTIONS` con headers `Access-Control-Request-*`.
Tu servidor debe responder `200` con los headers `Access-Control-Allow-*`
**sin** llamar al handler downstream. Es un _early return_ en el middleware.
