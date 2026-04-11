# Signature: 405, 404 y Fallback

## Mejora a ServeHTTP

Después del match exacto, antes del 404 final, agrega dos pasos:

### Detección de 405
Recorre todas las rutas registradas. Si alguna tiene el mismo `path` pero diferente `method`, responde `WriteError(w, 405, "Method Not Allowed")` y retorna.

### Fallback
Si no hay match por path y `router.Fallback != nil`, delega con `router.Fallback.ServeHTTP(w, r)` y retorna.

### Orden completo de ServeHTTP
1. Match exacto → dispatch
2. Mismo path, diferente method → 405
3. Fallback != nil → delegar
4. Nada → 404
