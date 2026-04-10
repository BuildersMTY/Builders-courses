# Hint: Router.ServeHTTP

1. Construye `key := routeKey{method: r.Method, path: r.Path}`.
2. Si `handler, ok := router.routes[key]; ok`, llama a `handler.ServeHTTP(w, r)`
   y retorna.
3. Si no, recorre `for k := range router.routes` y si encuentras uno con
   `k.path == r.Path`, llama a `WriteError(w, 405, "Method Not Allowed")` y
   retorna. **Este chequeo va antes del 404 para reportar el error correcto.**
4. Si `router.Fallback != nil`, llama a `router.Fallback.ServeHTTP(w, r)` y
   retorna.
5. Último recurso: `WriteError(w, 404, "Not Found")`.

`Handle` es simplemente `router.routes[routeKey{method, path}] = handler`.

`NewRouter` inicializa el mapa: `return &Router{routes: make(map[routeKey]Handler)}`.
