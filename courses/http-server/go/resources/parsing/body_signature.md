# Signature: parseBody

```go
func parseBody(reader *bufio.Reader, headers map[string]string) ([]byte, error)
```

### Comportamiento esperado

1. Busca `"Content-Length"` en el mapa de headers
2. Si no existe o es `"0"`, retorna `nil, nil` (no hay body)
3. Convierte el valor a int con `strconv.Atoi`
4. Crea un slice de bytes del tamaño indicado: `make([]byte, cl)`
5. Lee exactamente esos bytes con `io.ReadFull(reader, body)`
6. Retorna el body leído

### Paquetes necesarios

- `bufio` — `*bufio.Reader`
- `strconv` — `Atoi`
- `io` — `io.ReadFull`
