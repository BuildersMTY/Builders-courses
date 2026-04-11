# Hint: Implementar SetHeader y WriteHeader

## SetHeader — una línea

```go
func (r *response) SetHeader(key, value string) {
    r.headers[key] = value
}
```

## WriteHeader — el guard y la escritura

```go
func (r *response) WriteHeader(statusCode int) {
    if r.wroteHeader {
        return  // Protección contra doble escritura
    }
    r.statusCode = statusCode
    r.wroteHeader = true

    // Status line
    statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, StatusText(statusCode))
    r.conn.Write([]byte(statusLine))

    // Headers
    for k, v := range r.headers {
        r.conn.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
    }

    // Línea vacía = fin de headers
    r.conn.Write([]byte("\r\n"))
}
```

El flag `r.wroteHeader` es clave: previene que la status line se envíe dos veces, lo cual rompería el protocolo HTTP.
