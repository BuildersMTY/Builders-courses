package main

import "net"

// ResponseWriter abstracts writing an HTTP response back to the client socket.
// The student implements a concrete type behind this interface.
type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
	SetHeader(key, value string)
}

// response is the concrete ResponseWriter backed by a TCP connection.
// You must implement SetHeader, WriteHeader and Write on this type.
type response struct {
	conn        net.Conn
	headers     map[string]string
	statusCode  int
	wroteHeader bool
}

// NewResponseWriter builds a ResponseWriter for the given connection.
func NewResponseWriter(conn net.Conn) ResponseWriter {
	return &response{
		conn:    conn,
		headers: make(map[string]string),
	}
}

// SetHeader stores a header that will be flushed by WriteHeader.
// Calling SetHeader after WriteHeader has no effect.
func (r *response) SetHeader(key, value string) {
	// TODO: implement
}

// WriteHeader emits the status line followed by all stored headers and
// the terminating CRLF. Subsequent calls must be ignored (guarded by
// r.wroteHeader) to keep the HTTP framing valid.
func (r *response) WriteHeader(statusCode int) {
	// TODO: implement
}

// Write writes the body to the socket. If WriteHeader has not been called
// yet, assume 200 OK and auto-set Content-Length to len(b) before emitting
// the header.
func (r *response) Write(b []byte) (int, error) {
	// TODO: implement
	return len(b), nil
}

// StatusText returns the reason-phrase for common HTTP status codes.
func StatusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 400:
		return "Bad Request"
	case 403:
		return "Forbidden"
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

// WriteError is a helper to send an error response in plain text.
func WriteError(w ResponseWriter, statusCode int, message string) {
	w.SetHeader("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
