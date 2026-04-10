# TCP listeners en Go

En Go, un servidor TCP se monta con `net.Listen("tcp", addr)`. El listener
resultante expone `Accept() (net.Conn, error)` que **bloquea** hasta que
llega una conexión.

Un servidor concurrente sigue el patrón clásico _accept loop + goroutine por
conexión_:

```go
listener, err := net.Listen("tcp", ":8080")
if err != nil { return err }
defer listener.Close()

for {
    conn, err := listener.Accept()
    if err != nil {
        log.Printf("accept error: %v", err)
        continue // no mates el servidor por un error de una sola conexión
    }
    go handle(conn)
}
```

Puntos clave:

- `Accept` bloquea — sin goroutine, solo podrías servir una conexión a la vez.
- `continue` ante error de una conexión individual mantiene el servidor vivo.
- El listener debe cerrarse al salir (aunque en la práctica, el loop nunca
  termina en un servidor normal).
- Cada `net.Conn` implementa `io.Reader`/`io.Writer`, así que podemos envolverla
  en un `bufio.Reader` para parsear líneas cómodamente.
