# Signature: parseRequestLine

```go
func parseRequestLine(reader *bufio.Reader) (method, path, version string, err error)
```

### Comportamiento esperado

1. Lee la primera línea del reader con `reader.ReadString('\n')`
2. Recorta `\r\n` del final con `strings.TrimRight(line, "\r\n")`
3. Divide por espacio con `strings.Split(line, " ")`
4. Si no tiene exactamente 3 partes, retorna error
5. Retorna las 3 partes: method (ej: "GET"), path (ej: "/health"), version (ej: "HTTP/1.1")

### Ejemplo

Entrada: `"GET /health HTTP/1.1\r\n"`
Salida: `"GET", "/health", "HTTP/1.1", nil`

### Paquetes necesarios

- `bufio` — `*bufio.Reader`
- `strings` — `TrimRight`, `Split`
- `fmt` — `fmt.Errorf`
