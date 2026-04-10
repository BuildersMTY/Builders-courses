# Path Traversal y Static File Serving

Servir archivos desde disco es trivial — hasta que un atacante envía
`GET /../../../../etc/passwd`. Sin validar, tu servidor obedece.

## La amenaza

`filepath.Join("./static/", "../../../etc/passwd")` resuelve a
`"../../etc/passwd"`. Luego `os.Open` lee fuera de tu directorio público.

## La defensa (en dos capas)

**Capa 1 — Normalización con `filepath.Clean`.** Colapsa `..`, `.` y slashes
redundantes. Pero esto **no es suficiente**: `Clean("/static/../../etc")`
todavía escapa.

**Capa 2 — Directory jailing con rutas absolutas.**

```go
absBase, _ := filepath.Abs(h.baseDir)
absPath, _ := filepath.Abs(fullPath)
if !strings.HasPrefix(absPath, absBase) {
    // intento de traversal → 403 Forbidden
}
```

Solo autorizas la request si la ruta absoluta resuelta _empieza_ con la
ruta absoluta del directorio base. Cualquier intento de escapar es rechazado
con `403 Forbidden` antes de tocar el disco.

## Mime types

`Content-Type` se deriva de la extensión del archivo. Sin él, Chrome se
rehúsa a ejecutar JS/CSS aunque los sirvas. Mapea `.html`, `.css`, `.js`,
`.json`, `.png`, `.jpg`/`.jpeg` — cualquier otra cosa → `application/octet-stream`.

## Directorios → index.html

Cuando la request es `GET /`, `os.Stat` del directorio retorna
`info.IsDir() == true`. En ese caso, reapuntas `fullPath` a
`filepath.Join(fullPath, "index.html")` y sigues normalmente.
