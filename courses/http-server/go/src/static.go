package main

// StaticHandler serves files from a base directory on disk.
type StaticHandler struct {
	baseDir string
}

// NewStaticHandler returns a handler that serves files under baseDir.
func NewStaticHandler(baseDir string) *StaticHandler {
	return &StaticHandler{baseDir: baseDir}
}

// ServeHTTP resolves r.Path under h.baseDir and streams the file back.
// You must:
//  1. Clean the path to defeat traversal (filepath.Clean).
//  2. Absolutize both baseDir and the resolved file path and verify
//     the result still starts with the absolute baseDir. If not → 403.
//  3. If the target is a directory, fall back to its index.html.
//  4. Set Content-Type from the extension (see getMimeType).
//  5. Set Content-Length from the file size.
//  6. WriteHeader(200) and stream the file with io.Copy.
//
// On stat errors, reply 404. On open errors, reply 500.
func (h *StaticHandler) ServeHTTP(w ResponseWriter, r *Request) {
	// TODO: implement
	WriteError(w, 404, "Not Found")
}

// getMimeType returns a Content-Type string for the given lowercase
// file extension (including the leading dot). Unknown extensions
// fall back to application/octet-stream.
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
