# Signature: Write

```go
func (r *response) Write(b []byte) (int, error)
```

### Comportamiento esperado

1. Si `r.wroteHeader` es false (WriteHeader no fue llamado aún):
   - Auto-setear `Content-Length` al tamaño del body: `r.SetHeader("Content-Length", fmt.Sprintf("%d", len(b)))`
   - Llamar `r.WriteHeader(200)` — asume 200 OK por defecto
2. Escribir los bytes del body al socket: `r.conn.Write(b)`
3. Retornar el resultado de `r.conn.Write(b)`
