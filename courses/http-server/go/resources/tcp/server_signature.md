# Firmas del Módulo 1

```go
type Server struct {
    Addr    string  // ":8080"
    Handler Handler // router / middleware chain
}

// Start abre el listener TCP y corre el accept loop para siempre.
// Solo retorna error si el listener no puede abrirse.
func (s *Server) Start() error

// handleConnection parsea una request, la despacha al handler
// y cierra la conexión cuando termina.
func (s *Server) handleConnection(conn net.Conn)
```

El `Handler` viene definido en `router.go`:

```go
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
```

Helpers disponibles desde otros archivos del curso:

```go
func ParseRequest(reader *bufio.Reader) (*Request, error)
func NewResponseWriter(conn net.Conn) ResponseWriter
func WriteError(w ResponseWriter, statusCode int, message string)
```
