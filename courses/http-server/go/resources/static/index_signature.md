# Signature: Fallback a index.html

## Mejora a ServeHTTP

Después del `os.Stat` exitoso, verificar si es un directorio:

1. Si `info.IsDir()` es true:
   - Reapuntar: `fullPath = filepath.Join(fullPath, "index.html")`
   - Hacer `os.Stat` de nuevo sobre el nuevo path
   - Si index.html no existe, responder 404
2. Continuar con el flujo normal de servir el archivo
