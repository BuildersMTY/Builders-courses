# Hint: Implementar handleConnection

## Paso 1: Cerrar la conexión al terminar

```go
defer conn.Close()
```

Siempre usar `defer` para garantizar que la conexión se cierre, incluso si hay un error.

## Paso 2: Crear un reader bufferizado

```go
reader := bufio.NewReader(conn)
```

`bufio.NewReader` envuelve la conexión TCP para leer línea por línea eficientemente.

## Paso 3: Parsear el request

```go
req, err := ParseRequest(reader)
if err != nil {
    w := NewResponseWriter(conn)
    WriteError(w, 400, "Bad Request")
    return
}
```

Si el parseo falla (request malformado), respondemos 400 y cerramos.

## Paso 4: Crear ResponseWriter y delegar

```go
w := NewResponseWriter(conn)
if s.Handler != nil {
    s.Handler.ServeHTTP(w, req)
} else {
    WriteError(w, 404, "Not Found")
}
```

Delegamos al Handler configurado (normalmente el Router con middlewares).
