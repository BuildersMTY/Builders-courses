# Signature: SetHeader y WriteHeader

## Métodos a implementar

```go
func (r *response) SetHeader(key, value string)
func (r *response) WriteHeader(statusCode int)
```

### SetHeader
Almacena un header en `r.headers[key] = value`. No tiene efecto si `r.wroteHeader` ya es true.

### WriteHeader
1. Si `r.wroteHeader` es true, retorna inmediatamente (no-op)
2. Marca `r.wroteHeader = true` y guarda `r.statusCode = statusCode`
3. Escribe la status line: `fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, StatusText(statusCode))`
4. Itera `r.headers` y escribe cada uno como `"Key: Value\r\n"`
5. Escribe `"\r\n"` (línea vacía que separa headers del body)

Todo se escribe directamente a `r.conn` con `r.conn.Write([]byte(...))`.
