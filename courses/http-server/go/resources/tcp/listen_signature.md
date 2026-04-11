# Signature: Server.Start

## Tipo: Server

```go
type Server struct {
    Addr    string  // dirección de escucha, ej: ":8080"
    Handler Handler // componente que maneja cada request parseado
}
```

## Método a implementar

```go
func (s *Server) Start() error
```

### Comportamiento esperado

1. Abre un listener TCP en `s.Addr` con `net.Listen("tcp", s.Addr)`
2. Si el listener no puede abrirse, retorna el error
3. Defer `listener.Close()`
4. Entra en un loop infinito llamando `listener.Accept()`
5. Por cada conexión aceptada, lanza una goroutine: `go s.handleConnection(conn)`
6. Si `Accept()` falla en una conexión individual, loggea con `log.Printf` y hace `continue` — NO mata el servidor

### Paquetes necesarios

- `net` — `net.Listen`, `net.Conn`
- `log` — `log.Printf`
- `fmt` — `fmt.Errorf` para wrapping de errores
