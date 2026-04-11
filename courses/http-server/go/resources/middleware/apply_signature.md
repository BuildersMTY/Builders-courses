# Signature: ApplyMiddlewares

```go
func ApplyMiddlewares(base Handler, middlewares ...Middleware) Handler
```

### Comportamiento esperado

Itera el slice de middlewares en orden inverso (del último al primero) y envuelve progresivamente al handler base. El primer middleware de la lista será el más externo (se ejecuta primero al entrar, último al salir).

### Tipo Middleware

```go
type Middleware func(Handler) Handler
```

Un middleware recibe un Handler y retorna un Handler nuevo que lo envuelve.
