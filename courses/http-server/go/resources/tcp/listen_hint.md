# Hint: Implementar Server.Start

## Paso 1: Abrir el socket TCP

```go
listener, err := net.Listen("tcp", s.Addr)
if err != nil {
    return fmt.Errorf("error abriendo puerto %s: %w", s.Addr, err)
}
defer listener.Close()
```

`net.Listen` abre un socket TCP en la dirección dada. Si el puerto ya está en uso, retorna error.

## Paso 2: Accept loop

```go
for {
    conn, err := listener.Accept()
    if err != nil {
        log.Printf("Error en Accept: %v", err)
        continue  // NO matar el servidor por un error de una conexión
    }
    // ...
}
```

`Accept()` bloquea hasta que un cliente se conecta. Si falla una conexión individual, loggeamos y seguimos aceptando.

## Paso 3: Goroutine por conexión

```go
go s.handleConnection(conn)
```

Cada conexión se maneja en su propia goroutine para no bloquear el accept loop. Esto permite atender múltiples clientes simultáneamente.
