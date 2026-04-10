# Firmas del Módulo 3

```go
type ResponseWriter interface {
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
    SetHeader(key, value string)
}

type response struct {
    conn        net.Conn
    headers     map[string]string
    statusCode  int
    wroteHeader bool
}

func NewResponseWriter(conn net.Conn) ResponseWriter
func (r *response) SetHeader(key, value string)
func (r *response) WriteHeader(statusCode int)
func (r *response) Write(b []byte) (int, error)

func StatusText(code int) string // devuelve "OK", "Not Found", etc.
func WriteError(w ResponseWriter, statusCode int, message string)
```

Útiles para construir las líneas:

```go
fmt.Sprintf("HTTP/1.1 %d %s\r\n", code, reason)
fmt.Sprintf("%s: %s\r\n", key, val)
```
