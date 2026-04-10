# Firmas del Módulo 2

```go
type Request struct {
    Method  string            // "GET", "POST", ...
    Path    string            // "/echo"
    Version string            // "HTTP/1.1"
    Headers map[string]string
    Body    []byte
}

func ParseRequest(reader *bufio.Reader) (*Request, error)
```

Funciones útiles de stdlib:

```go
// Lee hasta (e incluyendo) el siguiente '\n'.
func (b *bufio.Reader) ReadString(delim byte) (string, error)

// Lee EXACTAMENTE len(buf) bytes o retorna error.
func io.ReadFull(r io.Reader, buf []byte) (n int, err error)

// Parsea un entero base-10.
func strconv.Atoi(s string) (int, error)

// Divide por el primer separador (útil para "Key: Value").
func strings.SplitN(s, sep string, n int) []string
func strings.TrimSpace(s string) string
func strings.TrimRight(s, cutset string) string
```
