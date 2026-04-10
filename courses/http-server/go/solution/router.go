package main

// Handler interface permite construir middlewares y enrutadores.
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

// HandlerFunc es una función que implementa Handler.
type HandlerFunc func(w ResponseWriter, r *Request)

// ServeHTTP permite que las funciones ordinarias sirvan como handlers HTTP.
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

// routeKey se utiliza para identificar una ruta única.
type routeKey struct {
	method string
	path   string
}

// Router actúa como multiplexor para mapear peticiones a sus handlers.
type Router struct {
	routes   map[routeKey]Handler
	Fallback Handler // Se usa resolver rutas dinámicas o servir archivos estáticos cuando no hay match exacto
}

// NewRouter crea un nuevo enrutador vacío.
func NewRouter() *Router {
	return &Router{
		routes: make(map[routeKey]Handler),
	}
}

// Handle registra un handler específico para el método y path dados.
func (router *Router) Handle(method, path string, handler HandlerFunc) {
	key := routeKey{method: method, path: path}
	router.routes[key] = handler
}

// ServeHTTP examina el request, y busca en el mapa de rutas para despachar el handler apropiado.
// Resuelve casos base como 404 (Not Found) o 405 (Method Not Allowed).
func (router *Router) ServeHTTP(w ResponseWriter, r *Request) {
	key := routeKey{method: r.Method, path: r.Path}

	// Match exacto
	if handler, ok := router.routes[key]; ok {
		handler.ServeHTTP(w, r)
		return
	}

	// Verificar si la ruta existe bajo otro método (generando 405)
	for k := range router.routes {
		if k.path == r.Path {
			WriteError(w, 405, "Method Not Allowed")
			return
		}
	}

	// Si hay fallback, delegamoste la petición (Ej. para archivos estáticos, Módulo 5)
	if router.Fallback != nil {
		router.Fallback.ServeHTTP(w, r)
		return
	}

	// Fallback final
	WriteError(w, 404, "404 Not Found")
}
