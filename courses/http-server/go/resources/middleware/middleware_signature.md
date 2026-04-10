# Firmas del Módulo 6

```go
type Middleware func(Handler) Handler

type loggingResponseWriter struct {
    ResponseWriter // embed: hereda Write y SetHeader
    statusCode int
}

func (w *loggingResponseWriter) WriteHeader(code int)

func LoggerMiddleware(next Handler) Handler
func CORSMiddleware(next Handler)   Handler
func ApplyMiddlewares(base Handler, middlewares ...Middleware) Handler
```

Headers CORS que debes setear:

```
Access-Control-Allow-Origin:  *
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Headers: Content-Type
```

Paquetes:

```go
import (
    "log"
    "time"
)

time.Now() time.Time
time.Since(t) time.Duration
log.Printf(format string, args ...any)
```
