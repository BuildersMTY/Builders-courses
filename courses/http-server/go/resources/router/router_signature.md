# Firmas del Módulo 4

```go
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

type HandlerFunc func(w ResponseWriter, r *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

type routeKey struct {
    method string
    path   string
}

type Router struct {
    routes   map[routeKey]Handler
    Fallback Handler
}

func NewRouter() *Router
func (r *Router) Handle(method, path string, handler HandlerFunc)
func (r *Router) ServeHTTP(w ResponseWriter, req *Request)
```
