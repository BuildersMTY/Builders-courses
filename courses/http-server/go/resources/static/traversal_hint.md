# Hint: Protección path traversal

Agrega esto ANTES de `os.Stat`:

```go
cleanPath := filepath.Clean(r.Path)
fullPath := filepath.Join(h.baseDir, cleanPath)

absBase, _ := filepath.Abs(h.baseDir)
absPath, _ := filepath.Abs(fullPath)
if !strings.HasPrefix(absPath, absBase) {
    WriteError(w, 403, "Forbidden")
    return
}
```

`filepath.Clean` normaliza `..` y `//`. Luego comparamos las rutas absolutas para verificar que el resultado sigue dentro del directorio permitido.
