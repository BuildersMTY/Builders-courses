# Firmas del Módulo 5

```go
type StaticHandler struct {
    baseDir string
}

func NewStaticHandler(baseDir string) *StaticHandler
func (h *StaticHandler) ServeHTTP(w ResponseWriter, r *Request)

func getMimeType(ext string) string // ".html" -> "text/html", etc.
```

Paquetes de stdlib que vas a usar:

```go
import (
    "io"
    "os"
    "path/filepath"
    "strings"
)

filepath.Clean(path)              // "/a/../b" -> "/b"
filepath.Join(base, rel)          // concatena con separator correcto
filepath.Abs(path) (string, error)
filepath.Ext(name) string         // "/x/y.css" -> ".css"
strings.ToLower(s)
strings.HasPrefix(s, prefix)

os.Stat(path) (os.FileInfo, error) // info.IsDir(), info.Size()
os.Open(path) (*os.File, error)

io.Copy(dst io.Writer, src io.Reader) (int64, error)
```

Recuerda: tu `ResponseWriter` implementa `io.Writer` vía su método `Write`, así
que `io.Copy(w, file)` funciona out-of-the-box.
