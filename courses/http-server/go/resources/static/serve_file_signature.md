# Signature: Servir un archivo

## Método a implementar (primera parte)

```go
func (h *StaticHandler) ServeHTTP(w ResponseWriter, r *Request)
```

### Flujo básico

1. Unir el path del request con el directorio base: `filepath.Join(h.baseDir, r.Path)`
2. Verificar que el archivo existe con `os.Stat(fullPath)` — si no, responder 404
3. Abrir el archivo con `os.Open(fullPath)` — si falla, responder 500
4. Defer `file.Close()`
5. Detectar el MIME type: `getMimeType(strings.ToLower(filepath.Ext(fullPath)))`
6. Setear headers: `Content-Type` y `Content-Length` (de `info.Size()`)
7. `WriteHeader(200)`
8. Copiar el archivo al socket: `io.Copy(w, file)`

### Paquetes necesarios
- `path/filepath`, `os`, `io`, `strings`, `fmt`
