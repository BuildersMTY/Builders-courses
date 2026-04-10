# Hint: StaticHandler.ServeHTTP paso a paso

```go
cleanPath := filepath.Clean(r.Path)
fullPath  := filepath.Join(h.baseDir, cleanPath)

absBase, _ := filepath.Abs(h.baseDir)
absPath, _ := filepath.Abs(fullPath)
if !strings.HasPrefix(absPath, absBase) {
    WriteError(w, 403, "Forbidden")
    return
}

info, err := os.Stat(fullPath)
if err != nil {
    WriteError(w, 404, "Not Found")
    return
}

if info.IsDir() {
    fullPath = filepath.Join(fullPath, "index.html")
    info, err = os.Stat(fullPath)
    if err != nil {
        WriteError(w, 404, "Not Found")
        return
    }
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
```

El orden de SetHeader → WriteHeader → io.Copy es crítico: todos los headers
deben estar fijados antes del primer byte al socket.
