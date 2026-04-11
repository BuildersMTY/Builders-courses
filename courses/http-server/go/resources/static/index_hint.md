# Hint: Fallback a index.html

Después del `os.Stat` exitoso, agrega:

```go
if info.IsDir() {
    fullPath = filepath.Join(fullPath, "index.html")
    info, err = os.Stat(fullPath)
    if err != nil {
        WriteError(w, 404, "File Not Found")
        return
    }
}
```

Cuando el usuario pide `/` (un directorio), servimos automáticamente `index.html` de ese directorio.
