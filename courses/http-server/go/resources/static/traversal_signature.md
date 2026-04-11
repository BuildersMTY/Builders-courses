# Signature: Protección path traversal

## Mejora a ServeHTTP

ANTES de abrir el archivo, agregar validación de seguridad:

1. Limpiar el path: `cleanPath := filepath.Clean(r.Path)`
2. Unir con base: `fullPath := filepath.Join(h.baseDir, cleanPath)`
3. Obtener rutas absolutas: `absBase, _ := filepath.Abs(h.baseDir)` y `absPath, _ := filepath.Abs(fullPath)`
4. Verificar contención: `if !strings.HasPrefix(absPath, absBase)` → responder `WriteError(w, 403, "Forbidden")`

Esto previene que un atacante acceda a archivos fuera del directorio usando `../../`.
