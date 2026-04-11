# Hint: Implementar Write

```go
func (r *response) Write(b []byte) (int, error) {
    if !r.wroteHeader {
        r.SetHeader("Content-Length", fmt.Sprintf("%d", len(b)))
        r.WriteHeader(200)
    }
    return r.conn.Write(b)
}
```

Si el usuario llama `Write` sin haber llamado `WriteHeader`, asumimos HTTP 200 OK y calculamos el Content-Length automáticamente. Esto simplifica el uso de la API — similar a como funcionan Express o Gin.
