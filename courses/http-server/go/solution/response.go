package main

import (
	"fmt"
	"net"
)

// ResponseWriter abstrae la escritura de la respuesta HTTP en el socket.
type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
	SetHeader(key, value string)
}

type response struct {
	conn        net.Conn
	headers     map[string]string
	statusCode  int
	wroteHeader bool
}

// NewResponseWriter crea una nueva instancia de ResponseWriter.
func NewResponseWriter(conn net.Conn) ResponseWriter {
	return &response{
		conn:    conn,
		headers: make(map[string]string),
	}
}

func (r *response) SetHeader(key, value string) {
	r.headers[key] = value
}

func (r *response) WriteHeader(statusCode int) {
	if r.wroteHeader {
		return
	}
	r.statusCode = statusCode
	r.wroteHeader = true

	// Escribir Status Line
	reason := StatusText(statusCode)
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, reason)
	r.conn.Write([]byte(statusLine))

	// Escribir Headers
	for k, v := range r.headers {
		headerLine := fmt.Sprintf("%s: %s\r\n", k, v)
		r.conn.Write([]byte(headerLine))
	}

	// Fin de headers (línea en blanco)
	r.conn.Write([]byte("\r\n"))
}

func (r *response) Write(b []byte) (int, error) {
	if !r.wroteHeader {
		// Si no se llamó a WriteHeader explícitamente, asumimos 200 OK
		// Y calculamos el Content-Length automáticamente (para Body de una sola escritura)
		r.SetHeader("Content-Length", fmt.Sprintf("%d", len(b)))
		r.WriteHeader(200)
	}
	return r.conn.Write(b)
}

// StatusText devuelve el Reason Phrase estándar para los códigos comunes.
func StatusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown Status"
	}
}

// WriteError es un helper para devolver errores comunes fácilmente en texto plano.
func WriteError(w ResponseWriter, statusCode int, message string) {
	w.SetHeader("Content-Type", "text/plain")
	w.SetHeader("Content-Length", fmt.Sprintf("%d", len(message)))
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
