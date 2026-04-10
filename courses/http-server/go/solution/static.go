package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// StaticHandler sirve directorio de archivos estáticos.
type StaticHandler struct {
	baseDir string
}

// NewStaticHandler inicializa el handler estático con la ruta base permitida.
func NewStaticHandler(baseDir string) *StaticHandler {
	return &StaticHandler{baseDir: baseDir}
}

// ServeHTTP busca un archivo en disco basado en el path del request HTTP.
func (h *StaticHandler) ServeHTTP(w ResponseWriter, r *Request) {
	// 1. Limpieza inicial para prevenir path traversal.
	// filepath.Clean normaliza ".", "..", y slashes redundantes.
	cleanPath := filepath.Clean(r.Path)

	// 2. Unimos la ruta limipia a nuestro directorio base local
	fullPath := filepath.Join(h.baseDir, cleanPath)

	// 3. Robustez contra Traversal: Asegurarse de que fullPath absoluto
	// realmente esté dentro del directorio base resuelto de forma absoluta.
	absBase, _ := filepath.Abs(h.baseDir)
	absPath, _ := filepath.Abs(fullPath)
	if !strings.HasPrefix(absPath, absBase) {
		WriteError(w, 403, "Forbidden - Intentos de traversal bloqueados")
		return
	}

	// 4. Verificamos si es directorio y manejamos index.html por defecto
	info, err := os.Stat(fullPath)
	if err != nil {
		WriteError(w, 404, "File Not Found")
		return
	}

	if info.IsDir() {
		fullPath = filepath.Join(fullPath, "index.html")
		info, err = os.Stat(fullPath)
		if err != nil {
			WriteError(w, 404, "File Not Found")
			return
		}
	}

	// 5. Lectura del archivo desde disco
	file, err := os.Open(fullPath)
	if err != nil {
		WriteError(w, 500, "Internal Server Error al abrir archivo")
		return
	}
	defer file.Close()

	// 6. Obtener mimetype en base a la extensión
	ext := strings.ToLower(filepath.Ext(fullPath))
	mimeType := getMimeType(ext)

	// 7. Enviar la respuesta HTTP
	w.SetHeader("Content-Type", mimeType)
	w.SetHeader("Content-Length", fmt.Sprintf("%d", info.Size()))
	w.WriteHeader(200)

	// io.Copy transmite en stream directo la data del open File al socket de red (ResponseWriter embebido sobre conn).
	io.Copy(w, file)
}

// getMimeType retorna el Content-Type para las extensiones soportadas
func getMimeType(ext string) string {
	switch ext {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	default:
		return "application/octet-stream"
	}
}
