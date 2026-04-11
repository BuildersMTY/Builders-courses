# Signature: parseHeaders

```go
func parseHeaders(reader *bufio.Reader) (map[string]string, error)
```

### Comportamiento esperado

1. Crea un `map[string]string` vacío para almacenar los headers
2. Lee líneas del reader en un loop con `reader.ReadString('\n')`
3. Recorta `\r\n` de cada línea
4. Si la línea está vacía (`""`), termina el loop — eso indica el fin de los headers
5. Divide cada línea por `":"` usando `strings.SplitN(line, ":", 2)`
6. Si no tiene 2 partes, ignora el header (continue) — no panic
7. Aplica `strings.TrimSpace` a key y value antes de guardarlos en el mapa

### Ejemplo

Entrada: `"Host: localhost\r\nContent-Type: text/plain\r\n\r\n"`
Salida: `map[string]string{"Host": "localhost", "Content-Type": "text/plain"}, nil`
