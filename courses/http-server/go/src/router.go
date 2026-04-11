package main

// Handler is anything that can serve an HTTP request/response cycle.
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

// HandlerFunc adapts a plain function to the Handler interface.
type HandlerFunc func(w ResponseWriter, r *Request)

// ServeHTTP makes HandlerFunc implement Handler.
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

// routeKey identifies a registered route by (method, path).
type routeKey struct {
	method string
	path   string
}

// Router is a method+path multiplexer.
// Fallback is consulted only when no exact route matches.
type Router struct {
	routes   map[routeKey]Handler
	Fallback Handler
}

// NewRouter returns an empty router ready to accept Handle calls.
// This constructor is pre-implemented — you only need to implement Handle
// and ServeHTTP below.
func NewRouter() *Router {
	return &Router{routes: make(map[routeKey]Handler)}
}

// Handle registers a handler for the given method and path.
func (router *Router) Handle(method, path string, handler HandlerFunc) {
	// TODO: implement
}

// ServeHTTP looks up the handler for (r.Method, r.Path) and dispatches.
// Rules:
//   - Exact match → call the handler.
//   - Path exists under a different method → 405 Method Not Allowed.
//   - No match and Fallback != nil → delegate to Fallback.
//   - Otherwise → 404 Not Found.
func (router *Router) ServeHTTP(w ResponseWriter, r *Request) {
	// TODO: implement
	WriteError(w, 404, "Not Found")
}
