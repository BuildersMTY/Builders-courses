# Hint: Servir un archivo básico

```go
func (h *StaticHandler) ServeHTTP(w ResponseWriter, r *Request) {
    fullPath := filepath.Join(h.baseDir, r.Path)

    info, err := os.Stat(fullPath)
    if err != nil {
        WriteError(w, 404, "File Not Found")
        return
    }

    file, err := os.Open(fullPath)
    if err != nil {
        WriteError(w, 500, "Internal Server Error")
        return
    }
    defer file.Close()

    ext := strings.ToLower(filepath.Ext(fullPath))
    w.SetHeader("Content-Type", getMimeType(ext))
    w.SetHeader("Content-Length", fmt.Sprintf("%d", info.Size()))
    w.WriteHeader(200)
    io.Copy(w, file)
}
```
